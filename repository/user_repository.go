package repository

import (
	"log"
)

func (r Repo) SignUp(userId string) {

	con, _ := getConnection()

	_, err := con.db.Exec(`insert into client(client_id) values($1)`, userId)

	if err != nil {
		log.Panic(err)
	}

}
