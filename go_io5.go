package main

import (
	"encoding/binary"
	"os"
	"fmt"
)

//go语言读取bmp文件

//bmp文件信息
type BitmapInfoHeader struct {
	Size uint32  //文件大小
	Width int32 //文件宽度
	Height int32 //文件高度
	Places uint16  //文件面数
	BitCount uint16 //每一个像素所占字节数
	Compression uint32 //压缩方式
	SizeImage uint32  //图片数据大小
	XperlsPerMeter int32 //水平分辨率
	YperlsPerMeter int32  //垂直分辨率
	ClusUsed uint32 //颜色数
	ClrImportant uint32
}

func main(){
	//实例化
	infoHeader := new(BitmapInfoHeader)
	//读取bmp文件
	file,_ := os.Open("timg.bmp")
	defer file.Close()

	binary.Read(file,binary.LittleEndian, infoHeader)

	fmt.Println(infoHeader)
}
