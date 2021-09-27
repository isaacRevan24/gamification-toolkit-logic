package repository

func (r Repo) SignUpRepository(userId string) error {
	const functionName string = "SignUpRepository"
	Logs.LogDebug("Start " + functionName)

	_, err := r.db.Exec(`insert into client(client_id) values($1)`, userId)

	if err != nil {
		Logs.LogError(err)
		return err
	}

	Logs.LogDebug("End " + functionName)

	return nil

}
