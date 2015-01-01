package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/fcasserfelt/timeapp/data"
	"github.com/fcasserfelt/timeapp/domain"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

var userRepo domain.UserRepository

func init() {
	var err error
	var db *sql.DB
	db, err = sql.Open("postgres", "dbname=timeapp sslmode=disable")
	PanicIf(err)

	userRepo = data.NewDbUserRepo(db)
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

	user := userRepo.FindByEmail("fredrik@bitjoy.se")

	fmt.Fprintf(w, "id: %d email:%s", user.Id, user.Email)

	/*
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
	*/

}
