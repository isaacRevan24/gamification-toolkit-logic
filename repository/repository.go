package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
	_ "github.com/lib/pq"
)

var (
	Logs utility.LoggingInterface
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "gamify"
)

type Repo struct {
	db *sql.DB
}

func GetConnection() (*Repo, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		return nil, err
	}

	return &Repo{db: db}, nil
}
