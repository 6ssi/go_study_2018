package main

import "fmt"

func main() {
	slice := [4]int[2,4,10,1]
	sum := 0
	for i:=0;i<len(slice);i++ {
		sum += slice[i]
	}
	fmt.Println("%d\n", sum)
}
