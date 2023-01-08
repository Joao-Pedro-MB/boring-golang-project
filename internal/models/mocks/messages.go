package mocks

import (
	"time"

	"github.com/Joao-Pedro-MB/boring-golang-project/internal/models"
)

var mockMessage = &models.Message{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type MessageModel struct{}

func (m *MessageModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *MessageModel) Get(id int) (*models.Message, error) {
	switch id {
	case 1:
		return mockMessage, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *MessageModel) Latest() ([]*models.Message, error) {
	return []*models.Message{mockMessage}, nil
}
