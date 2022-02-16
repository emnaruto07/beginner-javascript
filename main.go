package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gumroaduser"
	password = "shazeb123"
	dbname   = "dvdrental"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	var (
		actor_id    int
		first_name  string
		last_name   string
		last_update time.Time
	)
	rows, err := db.Query("select actor_id, first_name,last_name,last_update from actor where actor_id = $1", 3)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&actor_id, &first_name, &last_name, &last_update)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(actor_id, first_name, last_name, last_update)
	}
	err = rows.Err()
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}
}
