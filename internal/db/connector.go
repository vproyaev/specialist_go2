package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Connector struct {
	Db *sql.DB
}

func (c *Connector) Create(uri string) {
	db, err := sql.Open("postgres", uri)
	c.Db = db
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Connector) Close() error {
	return c.Db.Close()
}

// TODO: Add transactions!
//
//func (c *Connector) executeTransaction(fn func(tx *sql.Tx) (*sql.Rows, error)) *sql.Rows {
//	tx, err := c.Db.Begin()
//	if err != nil {
//		return nil
//	}
//	defer func() {
//		if p := recover(); p != nil {
//			tx.Rollback()
//			log.Fatal(p)
//		} else if err != nil {
//			tx.Rollback()
//		} else {
//			err = tx.Commit()
//		}
//	}()
//	var rows *sql.Rows
//	rows, err = fn(tx)
//	return rows
//}
//
//func (c *Connector) Execute(query string, args ...interface{}) (*sql.Rows, error) {
//	return c.executeTransaction(func(tx *sql.Tx) (*sql.Rows, error) {
//		rows, err := tx.Query(query, args...)
//		if err != nil {
//			return nil, err
//		}
//		return rows, nil
//	}), nil
//}
