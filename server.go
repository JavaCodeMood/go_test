package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var(
	//握手的过程中允许跨域
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin:func(r *http.Request) bool{
			return true   //总是允许跨域
		},
	}
)
//回调函数，传给一个请求和应答
func wsHandler(w http.ResponseWriter,r *http.Request){
	var(
		conn *websocket.Conn   //长连接
		err error
		data []byte
	)
	//Upgrader:websocket
	if conn, err = upgrader.Upgrade(w,r,nil); err != nil{
		return
	}

	//websocket.Conn
	//for循环做数据的收发
	for {
		//数据类型：Text, Binary, json
		if _,data,err = conn.ReadMessage();err!= nil{
			goto ERR    //如果出错就关闭连接
		}
		if err = conn.WriteMessage(websocket.TextMessage,data);err!= nil{
			goto ERR
		}
	}

	ERR:
		conn.Close()    //关闭连接

}

func main(){
	//配置路由  http://localhost:7777/ws
	http.HandleFunc("/ws",wsHandler)

	//启动服务端
	http.ListenAndServe("0.0.0.0:7777", nil)
}

