package file

import (
	"fmt"
	 "os"
	"time"
	"bufio"
	"strconv"
	"io"
)



const BUFF_SIZE = 128
const FILE_NAME = "aaa.ebl"

func FileLoader()  {
	//wg:=&sync.WaitGroup{}
	current_dir, _ := os.Getwd()
	fmt.Println(current_dir)
	inputFile, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Fprintf(os.Stderr, "OTA File open err:%s\n", err)
	}
	defer inputFile.Close()
	fileInfo, err := inputFile.Stat()
	fmt.Printf("the size of %s is :%d Byte\r\n", FILE_NAME, int32(fileInfo.Size()))
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	buffer1 := make([]byte, BUFF_SIZE)
	r := bufio.NewReader(inputFile)
	id:=0


	for {

		num,err:=io.ReadAtLeast(r,buffer1,BUFF_SIZE)
		if num!=BUFF_SIZE{
			if err == io.EOF {
				fmt.Println("write over")
				break
			}
		}

		CreatPart(buffer1,id)
		id++

	}

	//wg.Wait()
}


func CreatPart(buf []byte, sn int) {
	fo, err := os.Create("ota_image_" + strconv.Itoa(sn))
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	_, err = fo.Write(buf)
}


