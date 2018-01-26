package main

import (
	"fmt"
	"time"


	fl "../umelink/file"



)




func main() {


	start := time.Now().UnixNano()
	fl.FileLoader()
	end := time.Now().UnixNano()
	fmt.Printf("cost:%d ms\n", ((end-start)/1000000))
	time.Sleep(1*time.Second)

}
