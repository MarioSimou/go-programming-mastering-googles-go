package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id int
	first_name string
	last_name string
	is_alive bool
}

func main(){
	db, e := sql.Open("sqlite3", "users.db")
	if e != nil {
		log.Fatalf("Error: %s", e)
	}
	defer db.Close()

	if e := db.Ping(); e != nil {
		log.Fatalln("Unable to ping the database")
	}

	rows, e := db.Query("SELECT * FROM users")
	if e != nil {
		log.Fatalf("Error: %s", e)
	}
	defer rows.Close()
	var users []user
	for rows.Next() {
		var u user
		if e := rows.Scan(&u.id,&u.first_name,&u.last_name,&u.is_alive); e != nil {
			log.Fatalf("Error: %s", e)
		}
		users = append(users,u)
	}

	for _, user := range users {
		fmt.Printf("User: %v\n", user)
	}

	
	
}