package model

import (
	"encoding/json"
	"errors"
	"time"
)

// сервер для сохранения мероприятий (дата, продолжительность, название, имя пользователя)
// доменный объект Event {data Time, title, username string}
// валидация пришедших данных с /create_event и /update_event (не пустые поля, время в правльном формате)

type Event struct {
	ID      string `json:"user_id"`
	Date    time.Time
	DateStr string `json:"date"`
	Title   string `json:"title"`
}

const timeLayot = "2006-01-02"

func NewEvent(id string, date time.Time, title string) *Event {
	return &Event{ID: id, Date: date, Title: title}
}

func Unmarshal(data []byte) (*Event, error) {
	event := Event{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}
	if event.DateStr == "" || event.ID == "" || event.Title == "" {
		return nil, errors.New("empty ID/Date/Title")
	}
	event.Date, err = time.Parse(timeLayot, event.DateStr)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
