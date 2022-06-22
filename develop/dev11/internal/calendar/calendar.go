package calendar

import (
	"github.com/cvenkman/wbL2/develop/dev11/internal/model"
)

type Calendar struct {
	// данные инкапсулированы
	data map[string]*model.Event
}

func NewCalendar() *Calendar {
	return &Calendar{data: make(map[string]*model.Event)}
}

func (c *Calendar) Add(e *model.Event) {
	c.data[e.ID] = e
}

func (c *Calendar) Get(id string) *model.Event {
	return c.data[id]
}

func (c *Calendar) Delete(id string) {
	delete(c.data, id)
}
