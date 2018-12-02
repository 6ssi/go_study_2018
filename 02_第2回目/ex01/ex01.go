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
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
