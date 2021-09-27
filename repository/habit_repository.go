package repository

func (r Repo) AddNewHabit(userId string, name string, description string, condition string, repetition int) (int, error) {

	habitId := 0

	sqlStatemente := "INSERT INTO habit (name, description, condition, repetition, client_id) VALUES($2, $3, $4, $5, $1) RETURNING habit_id"

	err := r.db.QueryRow(sqlStatemente, userId, name, description, condition, repetition).Scan(&habitId)

	if err != nil {
		return 0, err
	}

	return habitId, nil
}
