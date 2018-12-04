package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type goengineer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Ability struct {
		Programming string `json:"programming"`
		Operation   string `json:"argus"`
	}
}

func main() {
	file, err := os.Open("./kon.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := json.Unmarshal(bytes, &persons); err != nil {
			log.Fatal(err)
		}
	}
}
