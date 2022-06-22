package model

import "time"

// сервер для сохранения мероприятий (дата, продолжительность, название, имя пользователя)
// доменный объект Event {data Time, title, username string}
// валидация пришедших данных с /create_event и /update_event (не пустые поля, время в правльном формате)

type Event struct {
	ID   string    `json:"id"`
	Date time.Time `json:"Date"`
}

func NewEvent(id string, date time.Time) *Event {
	return &Event{ID: id, Date: date}
}
