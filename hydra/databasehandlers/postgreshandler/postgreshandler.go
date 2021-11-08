package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type crewMember struct {
	id int
	name string
	secClearance int
	position string
}

type Crew []crewMember

func main() {
	connStr := "user=test password=test dbname=test sslmode=disable port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from personnel as p limit 10")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
}
