package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	get()
}

func get() {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	count := 1
	doc.Find("span").Each(func(_ int, s *goquery.Selection) {
		span, exist := s.Attr("class")
		if exist == true && span == "d-inline-block float-sm-right" && count <= 10 {
			stars := s.Find("span class").Text()
			fmt.Printf("%d位のStarの数 : %s\n", count, stars)
			count++
		}
	})
}
