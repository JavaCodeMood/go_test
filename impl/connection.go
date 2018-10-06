package impl

import (
	"github.com/gorilla/websocket"
	"sync"
	"errors"
)

//定义WebSocket的连接
type Connection struct{
	wsConn *websocket.Conn   //长连接
	inChan chan []byte   //接收读到的消息
	outChan chan[]byte    //需要发送的消息
	closeChan chan byte   //出错

	mutex sync.Mutex   //加锁
	isClosed bool   //关闭的开关
}

//封装WebSocket长连接
func InitConnection(wsConn *websocket.Conn)(conn *Connection,err error){
	conn = &Connection{
		wsConn:wsConn,  //封装长连接
		inChan:make(chan []byte,1000),   //1000个消息容量
        outChan:make(chan []byte,1000),
        closeChan:make(chan byte, 1),
	}
	//启动读协程
	go conn.readLoop()   //建立协程

	//启动写协程
	go conn.writeLoop()

	return
}

//读取消息API
func(conn *Connection) ReadMessage() (data []byte, err error){
	select {
	    case data = <- conn.inChan:
	  	case <- conn.closeChan:
	  		err = errors.New("连接已被关闭")
	}
	return
}

//发送消息API
func(conn *Connection) WriteMessage(data []byte)(err error){
	select {
	    case conn.outChan <- data:
	  	case <- conn.closeChan:
	  		err = errors.New("连接已被关闭")
	}
	return
}

//关闭连接API
func(conn *Connection) Close(){
	//线程安全，可重入的Close
	conn.wsConn.Close()

	//关闭chan,一个chan只能关闭一次,存在多次关闭的问题
	//要保证这行代码只能执行一次
	conn.mutex.Lock()   //锁住整个连接
	//判断是否已经被关闭
	if !conn.isClosed{   //如果没有被关闭
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()   //释放锁

}


//内部是实现
// 读取消息
func (conn *Connection) readLoop(){
	//定义变量
	var(
		data []byte
		err error
	)
	for{
		if _,data,err = conn.wsConn.ReadMessage(); err != nil{
			goto ERR   //如果发生错误
		}

		//如果容量满了，就会发生阻塞，等待inChan有空闲的位置
		select {
		    case conn.inChan <- data:
		  	case <-conn.closeChan:
		  		//closeChan关闭的时候
		  		goto ERR

		}
		conn.inChan <- data
	}

	ERR:
		conn.Close()
}

//发送消息
func (conn *Connection) writeLoop(){
	var(
		data []byte
		err error
	)
	for{
		select {
		   case data = <- conn.outChan: //取出需要发送的消息
		   case <- conn.closeChan:
		   	  goto ERR

		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err!= nil{
			goto ERR
		}
	}
	ERR:
		conn.Close()
}
