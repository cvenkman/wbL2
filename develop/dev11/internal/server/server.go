package server

import (
	"fmt"
	"net/http"

	"github.com/cvenkman/wbL2/develop/dev11/internal/calendar"
	"github.com/cvenkman/wbL2/develop/dev11/internal/config"
)

type Server struct {
	config   *config.Config
	calendar *calendar.Calendar
}

// NewServer создает новый сервер
func NewServer(c *config.Config) *Server {
	return &Server{config: c, calendar: calendar.NewCalendar()}
}

func (s *Server) Start() {
	http.HandleFunc("/create_event", s.Create)
	http.HandleFunc("/events_for_day", s.GetEventsForDay)
	http.ListenAndServe(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port), nil)
}
