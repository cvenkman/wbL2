package calendar

import (
	"time"

	"github.com/cvenkman/wbL2/develop/dev11/internal/model"
)

type Calendar struct {
	// данные инкапсулированы
	data map[string][]*model.Event
}

func NewCalendar() *Calendar {
	return &Calendar{data: make(map[string][]*model.Event)}
}

func (c *Calendar) Add(e *model.Event) {
	c.data[e.ID] = append(c.data[e.ID], e)
}

func (c *Calendar) Update(e *model.Event) {
	events := c.data[e.ID]
	for _, event := range events {
		if event.ID == e.ID && event.Date == e.Date {
			event.Title = e.Title
		}
	}
}

// Удаляет запись из мапы если все данные совпадаеют с e
func (c *Calendar) Delete(e *model.Event) {
	// все записи c данным ID
	events := c.data[e.ID]
	// удалить все запсии где данные совпадают с e
	for i, event := range events {
		// разыменовываем чтобы сравнить значения, а не ссылки
		if *event == *e {
			events = append(events[:i], events[i+1:]...)
		}
	}
	// записываем данные обратну в мапу, чтобы там сохранились данные
	c.data[e.ID] = events
}

// ищет все записи за данный день и возвращает массив Event с ними
func (c *Calendar) GetEventsForDay(user_id string, day time.Time) []*model.Event {
	eventsForDay := make([](*model.Event), 0)

	for key, date := range c.data {
		if key == user_id {
			// проходимся по массиву event у этого user_id
			for _, event := range date {
				if event.Date == day {
					eventsForDay = append(eventsForDay, event)
				}
			}
		}
	}
	return eventsForDay
}

func (c *Calendar) GetEventsForWeek(user_id string, day time.Time) {

}
