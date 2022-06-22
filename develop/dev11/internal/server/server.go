package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
	http.HandleFunc("/create_event", Logger(s.Create))
	http.HandleFunc("/events_for_day", Logger(s.GetEventsForDay))
	http.ListenAndServe(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port), nil)
}

// middleWare функция для логирования запросов
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s, %s, %s\n", r.Method, r.URL, time.Since(start))
	}
}
