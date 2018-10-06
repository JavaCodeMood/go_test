package main

//通道(channel)和goroutine搭配，实现用通信代替内存共享的CSP模型

//消费者
func consumer(data chan int, done chan bool){
	for x:= range data{   //接收数据直到通道被关闭
		println("recv: ",x)
	}
	done <- true   //通知main，消费结束
}

//生产者
func producer(data chan int){
	for i:=0;i < 5;i++ {   //发送数据
		data <- i
	}
	close(data)   //产生结束，关闭通道
}

func main(){
	done := make(chan bool)   //用于接收消费结束的信号
	data := make(chan int)   //数据管道

	go consumer(data, done)  //启动消费者
	go producer(data)   //启动生产者

	<- data    //阻塞，直到消费者返回结束信息
}


