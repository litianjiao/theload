package main

import (
	//"io/ioutil"
	"fmt"
	"os"
	//"syscall"

	"bufio"
	//"strconv"
<<<<<<< HEAD
	"io"
	"strconv"
=======
>>>>>>> origin/master
)

const BUFF_SIZE = 128
const FILE_NAME = "aaa.ebl"

func main() {
	current_dir, _ := os.Getwd()
	fmt.Println(current_dir)

	inputFile, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Fprintf(os.Stderr, "OTA File open err:%s\n", err)
	}
	defer inputFile.Close()
	fileInfo, err := inputFile.Stat()
	fmt.Printf("the size of %s is :%d Byte\r\n", FILE_NAME, int32(fileInfo.Size()))
	buffer1 := make([]byte, BUFF_SIZE)
	for{
		r := bufio.NewReader(inputFile)
		if err!=nil&&err!=io.EOF{
			panic(err)
		}
		if nil == 0 {
			s1:=strconv.Itoa(index)
			fw,err:=os.Create("ota_image_"+s1)
		}
	}



	r := bufio.NewReader(inputFile)

	buffer1, err = r.Peek(BUFF_SIZE)
	fmt.Printf("%s\n", string(buffer1))
	s1 := "1"
	fw, err := os.Create("ota_image_" + s1)
	if err != nil {
		panic(err)

	}
	defer fw.Close()
	fw.Write(buffer1)
	//批量操作
	//
	//file, err := os.Open(FILE_NAME)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//var num int16
	//num = (int16(fileInfo.Size()) + BUFF_SIZE - 1) / BUFF_SIZE
	//flag := false
	//for flag == true {
	//	part, pos := cat(file, num)
	//	fo, err := os.Create("ota_image_" + strconv.Itoa(num))
	//	if err != nil {
	//		panic(err)
	//	}
	//	_, err = fo.Write(part)
	//}

}

//func cat(f *os.File, num int16) ([]byte, int64) {
//	var pkg []byte
//	var index int64
//	index = int64(num * BUFF_SIZE)
//	for i := 0; i < 128; i++ {
//		buf := make([]byte, BUFF_SIZE)
//		switch nr, err := f.Read(buf[index : index+int64(i)]); true {
//		case nr < 0:
//			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
//			os.Exit(1)
//
//		case nr == 0: //EOF
//
//			return pkg, 0
//		case nr > 0 :
//			pkg = append(pkg, buf...)
//		}
//		pos := index
//		return pkg, pos
//	}
//}

//func createPart(f *os.File, id int) {
//	var offset int64
//	offset = int64(id * BUFF_SIZE)
//	buf := make([]byte, BUFF_SIZE)
//
//	buf=append(buf,)
//	for i := 0; i < BUFF_SIZE; i++ {
//		//f.Seek(offset, os.SEEK_SET)
//		switch nr,err:=f.Read(buf[offset:offset+int64(i)]);true {
//
//		}
//	}
//
//	fo, err := os.Create("ota_image_" + strconv.Itoa(id))
//	if err != nil {
//		panic(err)
//	}
//	_, err = fo.Write(buf)
//}
