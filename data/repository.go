package data

import (
	"database/sql"
	"fmt"
	"github.com/fcasserfelt/timeapp/domain"
	_ "github.com/lib/pq"
)

type DbRepo struct {
	db *sql.DB
}

type DbUserRepo DbRepo

func NewDbUserRepo(db *sql.DB) *DbUserRepo {
	dbUserRepo := new(DbUserRepo)
	dbUserRepo.db = db
	return dbUserRepo
}

func (repo *DbUserRepo) FindByEmail(email string) domain.User {

	row, err := repo.db.Query("SELECT id, email FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var id int
	row.Next()
	row.Scan(&id, &email)
	user := domain.User{Id: id, Email: email}
	return user

}
