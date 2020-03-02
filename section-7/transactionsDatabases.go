package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var e error

type user struct {
	id int
	first_name string
	last_name string
	is_alive bool
}

func (u user) String() string {
	return fmt.Sprintf("(%d) %s %s - %t", u.id, u.first_name, u.last_name, u.is_alive)
}

type users []user

func (u users) String() string {
	var s []string
	for _, user := range u {
		s = append(s,user.String())
	}
	return strings.Join(s,"\n")
} 


func main(){
	var mu users

	if db, e = sql.Open("sqlite3", "users.db"); e != nil {
		logError(e)
	}
	if e = db.Ping(); e != nil {
		logError(e)
	}
	if mu, e = readUsers(); e != nil {
		logError(e)
	}
	fmt.Printf("Initial set of users:\n%s\n", mu)

	var trans, e = db.Begin()
	logError(e)

	if _, e = trans.Exec("UPDATE users SET is_alive=?", false); e != nil {
		logError(e)		
	}

	if mu, e = readUsers(); e != nil {
		logError(e)
	}
	fmt.Printf("Before commit users:\n%s\n", mu)

	if e = trans.Commit(); e != nil {
		logError(e)
	}

	if mu, e = readUsers(); e != nil {
		logError(e)
	}
	fmt.Printf("After commit users:\n%s\n", mu)
}

func logError(e error) {
	if e != nil {
		log.Fatalf("Error: %v\n", e)
	}
}

func readUsers() (users, error){
	rows, e := db.Query("SELECT * FROM users")
	if e != nil {
		return nil, e
	}

	var mu users
	for rows.Next() {
		var u user
		if e := rows.Scan(&u.id,&u.first_name,&u.last_name,&u.is_alive); e !=  nil {
			return nil, e
		}
		mu = append(mu,u)
	}

	return mu,nil
}