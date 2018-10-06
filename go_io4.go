package main

import (
	"os"
	"encoding/binary"
	"fmt"
)

//go语言读取bmp文件

func main(){
	//读取bmp文件
	file,_ := os.Open("timg.bmp")
	defer file.Close()
	var headA,headB byte
	//对二进制进行封装,读取文件头
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	//文件大小
	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	//保留的字节
	var reservedA, reservedB byte
	binary.Read(file,binary.LittleEndian,&reservedA)
	binary.Read(file,binary.LittleEndian,&reservedB)

	//偏移数
	var offbits uint32
	binary.Read(file,binary.LittleEndian,&offbits)

	fmt.Println("headA = ", headA, "headB = ", headB)
	fmt.Println("size =" , size)
	fmt.Println("reservedA = ",reservedA,"reservedB = ", reservedB)
	fmt.Println("offbits = ", offbits)

}
