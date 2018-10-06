package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"time"
)

var(
	//握手的过程中允许跨域
	Upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin:func(r *http.Request) bool{
			return true   //总是允许跨域
		},
	}
)
//回调函数，传给一个请求和应答
func wsHandler(w http.ResponseWriter,r *http.Request){
	var(
		wsConn *websocket.Conn   //长连接
		err error
		data []byte
		conn *impl.Connection
	)
	//Upgrader:websocket
	if wsConn, err = upgrader.Upgrade(w,r,nil); err != nil{
		return
	}

	if conn.err = impl.InitConnection(wsConn); err!= nil{
		goto ERR
	}

	//实现一个独立的协程，每个一秒发送一条消息
	go func(){
		var(
			err error
		)
		if err = conn.WriteMeeage([]byte("heartbeat")); err!=nil{
           return
		}
		time.Sleep(1 * time.Second)
	}()

	//处理消息
	for{
		if data,err = conn.ReadMessage(); err != nil{
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil{
			goto ERR
		}
	}

	ERR:
		//关闭连接的操作
        conn.Close()

}

func main(){
	//配置路由  http://localhost:7777/ws
	http.HandleFunc("/ws",wsHandler)

	//启动服务端
	http.ListenAndServe("0.0.0.0:7777", nil)
}

