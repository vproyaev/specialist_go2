package db

import (
	"database/sql"

	"specialist/internal/models"
)

func (c *Connector) GetTags() ([]Tag, error) {
	var tags []Tag

	rows, err := c.Db.Query("SELECT * FROM tag")
	if err != nil {
		return tags, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return tags, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (c *Connector) GetTag(tagId string) (Tag, error) {
	var tag Tag

	rows, err := c.Db.Query(
		"SELECT * FROM tag WHERE id = $1 limit 1", tagId,
	)
	if err != nil {
		return tag, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return tag, err
		}
		return tag, nil
	}

	return tag, nil
}

func (c *Connector) CreateTag(data models.InputTagData) error {
	_, err := c.Db.Query(
		"INSERT INTO tag (name) VALUES ($1)", data.Name,
	)
	return err
}
