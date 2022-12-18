package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type MessageModel struct {
	DB *sql.DB
}

func (m *MessageModel) Insert(title string, content string, expires int) (int, error) {

	stmt := `INSERT INTO messages (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *MessageModel) Get(id int) (*Message, error) {
	return nil, nil
}

func (m *MessageModel) Latest() ([]*Message, error) {
	return nil, nil
}
