package entities

import "github.com/google/uuid"

type Event struct {
	ID     string `json:"id"`
	Bucket string `json:"bucket"`
}

func NewEvent() *Event {
	event := Event{
		ID: uuid.New().String(),
	}

	return &event
}
