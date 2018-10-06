package main

//字符串读取

import (
	"io"
	"strings"
	"fmt"
	"os"
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
	data,_ :=ReadFrom(strings.NewReader("from string"),12)

	fmt.Println(data)
}

//从命令行读取字符串
func sampleReadStdin(){
	fmt.Println("请输入一个字符串")

	//调用命令行，读取11个字符
	data,_:= ReadFrom(os.Stdin,11)

	fmt.Println(data)
}

//从文件读取字符
func sampleReadFile(){
	//打开文件
	file,_ := os.Open("go_io.go")
	//关闭文件
	defer file.Close()

	data,_ := ReadFrom(file, 10)
	fmt.Println(data)
}

func main(){
	//sampleReadFromString()
	//sampleReadStdin()
	sampleReadFile()
}