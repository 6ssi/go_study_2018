package main

import (
	"database/sql"
	"fmt"

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
	fmt.Println("--------------------------------------")
	for rows.Next() {
		err := rows.Scan(rowvals...)
		if err != nil {
			fmt.Printf("Unable to Scan values from Select Query")
		}
		for i, col := range clmvals {
			fmt.Println(clm[i], ":", string(col))
		}
		fmt.Println("--------------------------------------")
	}

}
