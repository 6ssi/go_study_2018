package main

import (
	"fmt"
	"log"
	"os"
)

/* 第2回目のex01 */

func main() {
	file, err := os.Open("./kon.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	//byteタイプを100個分格納するスライスを作る、という意味らしい
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", data[:count])
}
