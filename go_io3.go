package main

import (
	"os"
	"fmt"
	"bufio"
)

//读取源码文件，并统计行数

func main(){
	//从命令行传入参数,如果参数的长度小于，就认为是非法的
	if len(os.Args)  < 2{
		return
	}

	//切片的第二个参数作为文件名
	filename := os.Args[1]

	//打开文件
	file,err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()   //关闭文件

	//读写器
	reader := bufio.NewReader(file)
	var line int   //行号
	//循环读取文件
	for{
      _,isPrefix,err := reader.ReadLine()  //读取文件的每一行
      if err != nil {
      	break
	  }

	  if !isPrefix {  //false时，统计行号
	  	line++
	  }
	}

	fmt.Println(line)
}
