package main

//字符串读取

import (
	"io"
	"strings"
	"fmt"
)

/**
读取信息
num:读取的字节数量
 */
func ReadFrom(reader io.Reader,num int)([]byte,error){
	//创建一个缓冲切片
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n>0{
		return p[:n],nil
	}

	return p,err
}

//字符串读取器
func sampleReadFromString(){
	data,_ :=ReadFrom(strings.NewReader("hello golang"),12)

	fmt.Println(data)
}


func main(){
	sampleReadFromString()
}