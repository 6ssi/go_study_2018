package main

<<<<<<< HEAD
import "fmt"
import "os"
import "os/exec"

func main() {
	out, err :=exec.Command(os.Args).Output()
	if err != nil{
		fmt.Printf("Oops -> %s", err)
		os.Exit(1)
		/* Godoc says code "0" indicates success. */
	}
	fmt.Printf("%s", out)
}
=======
/* ex02 */

import (
	"fmt"
	"os"
)

func main() {
	countargs(os.Args)
	printargs(os.Args)
}

func countargs(inputargs []string) {
	fmt.Printf("You have "+"%d"+" arguments.\n", len(inputargs[1:]))
	/* inputargs[1:]
	can also be
		inputargs[1 : len(inputargs)] */
}

func printargs(inputargs []string) {
	fmt.Printf("Here are your arguments:\n")
	for i := 1; i < len(inputargs); i++ {
		fmt.Printf("%s\n", inputargs[i])
	}
}
>>>>>>> ead6b4ccdc31f18b99788f6f0160fc22bff96e97
