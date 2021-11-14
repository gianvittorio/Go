package hydrablayer

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type mySqlDataStore struct {
	*sql.DB
}

func NewMySQLDataStore(conn string) (*mySqlDataStore, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &mySqlDataStore{DB:db}, nil
}

func (msql *mySqlDataStore) AddMember(cm *CrewMember) error {
	_, err := msql.Exec("insert into personnel (name, security_clearance, position) values($1, $2, $3)", cm.Name, cm.SecClearance, cm.Position)
	return err
}

func (msql *mySqlDataStore) FindMember(id int) (CrewMember, error) {
	row := msql.QueryRow("select * from personnel where id = $1", id)
	cm := CrewMember{}
	err := row.Scan(&cm.ID, &cm.Name, &cm.SecClearance, &cm.Position)
	return cm, err
}

func (msql *mySqlDataStore) AllMembers() (crew, error) {
	rows, err := msql.Query("select * from personnel;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := crew{}
	for rows.Next() {
		member := CrewMember{}
		err := rows.Scan(&member.ID, &member.Name, &member.SecClearance, &member.Position)
		if err == nil {
			members = append(members, member)
		}
	}

	err = rows.Err()
	return members, err
}
