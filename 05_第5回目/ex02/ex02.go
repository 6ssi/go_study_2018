package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	if err != nil {
		fmt.Println("DB Connection Error")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		fmt.Println("Unable to select from table")
	}

	clm, err := rows.Columns()
	if err != nil {
		fmt.Println("Unable to fetch colums from table")
	}

	clmvals := make([]sql.RawBytes, len(clm))
	rowvals := make([]interface{}, len(clmvals))

	for i := range clmvals {
		rowvals[i] = &clmvals[i]
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("--------------------------------------")
	for rows.Next() {
		for i, col := range clmvals {
			fmt.Printf(clm[i], ":", string(col))
		}
		fmt.Println("--------------------------------------")
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		fmt.Fprint(w, "hello web server!")
	} else {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}
}
