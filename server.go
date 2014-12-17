package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "dbname=timeapp sslmode=disable")
	PanicIf(err)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run("localhost:3000")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, email FROM users")
	PanicIf(err)
	defer rows.Close()

	var id int
	var email string
	for rows.Next() {
		err := rows.Scan(&id, &email)
		PanicIf(err)
		fmt.Fprintf(w, "id: %d email:%s", id, email)
	}

}
