package main

/* ex02 */

import (
	"fmt"
	"os"
)

func main() {
	inputargs := []string(os.Args)
	fmt.Printf("You have "+"%d"+" arguments.\n", len(inputargs[1:]))
	/* inputargs[1:]
	is same as
		inputargs[1 : len(inputargs)]
	*/
	fmt.Printf("Here are your arguments:\n")
	for i := 1; i < len(inputargs); i++ {
		fmt.Printf("%s\n", inputargs[i])
	}
}
