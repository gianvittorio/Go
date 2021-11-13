package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
		log.Fatal("Could not connect, error ", err.Error())
	}
	defer db.Close()

	cw := GetCrewByPositions(db, []string{"'Mechanic'", "'Technician'"})
	fmt.Println(cw)

	fmt.Println(GetCrewMemberById(db, 4))

	AddCrewMember(db, crewMember{name: "Steve Bee", secClearance: 4, position: "Biologist"})

	// fmt.Println(GetCrewMemberByPosition(db, "Chemist"))
}

func GetCrewMemberById(db *sql.DB, id int) (cm crewMember) {
	row := db.QueryRow("select * from personnel where id = $1", id)

	err := row.Scan(&cm.id, &cm.name, &cm.secClearance, &cm.position)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func AddCrewMember(db *sql.DB, cm crewMember) int64 {

	res, err := db.Exec("insert into personnel (name, security_clearance, position) values($1, $2, $3)", cm.name, cm.secClearance, cm.position)
	if err != nil {
		log.Fatal(err)
	}
	ra, _ := res.RowsAffected()
	id, _ := res.LastInsertId()

	log.Println("Rows Affected", ra, "Last inserted id", id)

	return id
}

func GetCrewByPositions(db *sql.DB, positions []string) Crew {
	Qs := fmt.Sprintf("select id, name, security_clearance, position from personnel as p where position in (%s)", strings.Join(positions, ","))

	rows, err := db.Query(Qs)
	if err != nil {
		log.Fatal("Could not get data from the personnel table ", err)
	}
	defer rows.Close()

	retVal := Crew{}
	cols, _ := rows.Columns()
	fmt.Println("Columns detected: ", cols)

	for rows.Next() {
		member := crewMember{}
		err = rows.Scan(&member.id, &member.name, &member.secClearance, &member.position)
		if err != nil {
			log.Fatal("Error scanning row ", err)
		}
		retVal = append(retVal, member)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return retVal
}
