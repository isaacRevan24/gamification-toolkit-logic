package repository

func (r *repo) AddNewHabitRepository(userId string, name string, description string, condition string, repetition int) (int, error) {
	const functionName string = "AddNewHabitRepository"
	Logs.LogDebug("Start " + functionName)

	habitId := 0

	sqlStatemente := "INSERT INTO habit (name, description, condition, repetition, client_id) VALUES($2, $3, $4, $5, $1) RETURNING habit_id"

	err := r.db.QueryRow(sqlStatemente, userId, name, description, condition, repetition).Scan(&habitId)

	if err != nil {
		Logs.LogError(err)
		Logs.LogDebug("End " + functionName)
		return 0, err
	}

	Logs.LogDebug("End " + functionName)

	return habitId, nil
}
