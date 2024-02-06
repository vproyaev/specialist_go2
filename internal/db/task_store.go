package db

import (
	"database/sql"
	"time"

	"specialist/internal/models"
)

func (c *Connector) GetTasks() ([]Task, error) {
	var tasks []Task

	rows, err := c.Db.Query("SELECT * FROM task")
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

func (c *Connector) GetTask(taskId string) (Task, error) {
	var task Task

	rows, err := c.Db.Query(
		"SELECT * FROM task WHERE id = $1 limit 1", taskId,
	)
	if err != nil {
		return task, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.TagId, &task.Due)
		if err != nil {
			return task, err
		}
		return task, nil
	}

	return task, nil
}

func (c *Connector) CreateTask(data models.InputTaskData) error {
	dueDate := time.Date(
		data.DueYear,
		time.Month(1),
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	if data.DueMonth != nil {
		dueDate = dueDate.AddDate(0, *data.DueMonth, 0)
	}
	if data.DueDay != nil {
		dueDate = dueDate.AddDate(0, 0, *data.DueDay)
	}

	_, err := c.Db.Query(
		"INSERT INTO task (tag_id, due) VALUES ($1, $2)", data.TagId, dueDate,
	)
	return err
}
func (c *Connector) DeleteTask(taskID string) error {
	_, err := c.Db.Query(
		"DELETE FROM task WHERE id = $1", taskID,
	)
	return err
}
func (c *Connector) DeleteTasks() error {
	_, err := c.Db.Query(
		"DELETE FROM task WHERE true",
	)
	return err
}
