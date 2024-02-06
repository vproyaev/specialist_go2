package db

import (
	"database/sql"
	"time"
)

func (c *Connector) GetDueTasks(year int, month *int, day *int) ([]Task, error) {
	var tasks []Task

	dueDate := time.Date(
		year,
		time.Month(1),
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	if month != nil {
		dueDate = dueDate.AddDate(0, *month, 0)
	}
	if day != nil {
		dueDate = dueDate.AddDate(0, 0, *day)
	}

	rows, err := c.Db.Query("SELECT * FROM task WHERE due >= $1", dueDate)
	if err != nil {
		return tasks, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.TagId, &task.Due)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
