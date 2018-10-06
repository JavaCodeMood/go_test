package main

//字符串读取

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func main(){
	strReader := strings.NewReader("hello world")

	bufReader := bufio.NewReader(strReader)

	//1.使用Peek()方法读取数据，只读，不会改变缓冲的大小
	data,_ := bufReader.Peek(10)

	fmt.Println(data)
	//bufReader.Buffered()  返回字节个数， 抓取缓冲数据
	fmt.Println(string(data), bufReader.Buffered())

	//2.使用ReadString()读取数据,读取到空格为止
	str,_ := bufReader.ReadString(' ')
	fmt.Println(str, bufReader.Buffered())

	str1,_ := bufReader.ReadString(4)
	fmt.Println(str1)

	//写出到输出设置中
	w := bufio.NewWriter(os.Stdout)   //输出
	fmt.Fprint(w, "hello")  //写入到bufio中
	fmt.Fprint(w, " world!")
	w.Flush()   //真正写入到Stdout中

}
