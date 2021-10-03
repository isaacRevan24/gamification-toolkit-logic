//go:generate go run github.com/golang/mock/mockgen -source connection.go -destination mock/connection_mock.go -package mock
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

type UserRepository interface {
	SignUpRepository(userId string) error
}

type HabitRepository interface {
	AddNewHabitRepository(userId string, name string, description string, condition string, repetition int) (int, error)
}

type repo struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	repo, _ := getConnection()
	return repo
}

func NewHabitRepository() HabitRepository {
	repo, _ := getConnection()
	return repo
}

func getConnection() (*repo, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		return nil, err
	}

	return &repo{db: db}, nil
}
