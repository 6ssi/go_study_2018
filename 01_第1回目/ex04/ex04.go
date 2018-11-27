package main

/* ex03 */

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	inputcmd := flag.String("cmd", "ls", "Type your command")
	inputopt := flag.String("opt", "l", "Type your option")
	flag.Parse()
	strcmd := *inputcmd
	stropt := *inputopt
	cmd := exec.Command(strcmd, stropt)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Oops -> %s", err)
		os.Exit(1)
	}
	fmt.Printf("Your command returned :\n%s\n", out)
}
