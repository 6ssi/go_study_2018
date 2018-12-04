package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type goengineer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Ability struct {
		Programming string `json:"programming"`
		Operation   string `json:"operation"`
	}
}

func main() {
	file, err := os.Open("./go_study_member.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	engdecoder := json.NewDecoder(file)
	for scanner.Scan() {
		goeng, err := engdecoder.Token()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T%v", goeng, goeng)
	}
}
