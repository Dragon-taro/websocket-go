package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Message struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func GetAll() (messages []Message, err error) {
	dbx, err := connect()
	if err != nil {
		return nil, err
	}

	if err := dbx.Select(&messages, "select * from messages"); err != nil {
		return nil, err
	}
	return messages, nil
}

func connect() (*sqlx.DB, error) {
	return sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			"root",
			"password",
			"localhost",
			"3306",
			"websocket_go",
		),
	)
}
