package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type user struct {
	id int
	lastName string
	firstName string
	isAlive bool
}

func index(w http.ResponseWriter, r *http.Request){
	var lastName = r.FormValue("lastName")
	if lastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Invalid lastname\n")		
		return
	}
	
	var row = db.QueryRow("SELECT * FROM users WHERE last_name=?", lastName)
	var u user
	if e := row.Scan(&u.id,&u.firstName,&u.lastName,&u.isAlive); e != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w,e.Error())		
		return		
	}
	fmt.Println(u)
	fmt.Fprintln(w,u.firstName)
}

func main(){
	var e error
	var port = fmt.Sprintf(":%d", 3000)

	db, e = sql.Open("sqlite3", "users.db")
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
	if e := db.Ping(); e != nil {
		log.Fatalf("Error: %v", e)
	}
	
	http.HandleFunc("/api", index)
	
	log.Fatalln(http.ListenAndServe(port, nil))
}