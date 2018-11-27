package main

/* ex03 */

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//inputcmd := flag.String("cmd", "ls", "Type your command")
	//inputopt := flag.String("opt", "l", "Type your option")
	//flag.Parse()
	//execcmd := *inputcmd + " -" + *inputopt
	//fmt.Printf("%s\n", execcmd)
	comstr := [...]string{"ls", " -l"}
	cmd := exec.Command(comstr)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Oops -> %s", err)
		os.Exit(1)
	}
	fmt.Printf("Your command returned :\n%s\n", out)
}
