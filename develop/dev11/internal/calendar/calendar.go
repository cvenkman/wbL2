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

func (c *Calendar) GetAllEvents(id string) []*model.Event {
	return c.data[id]
}

func (c *Calendar) Delete(id string) {
	delete(c.data, id)
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

// for _, v := range es.data {
// 	if v.UserID == userID {
// 		if (v.Date == start || v.Date.After(start)) && v.Date.Before(end) {
// 			ev = append(ev, v)
// 		}
// 	}
// }
