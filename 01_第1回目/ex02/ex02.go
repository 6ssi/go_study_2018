package main

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