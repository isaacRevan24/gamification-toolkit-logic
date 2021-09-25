package repository

func (r Repo) SignUp(userId string) error {

	con, _ := GetConnection()

	_, err := con.db.Exec(`insert into client(client_id) values($1)`, userId)

	if err != nil {
		return err
	}

	return nil

}
