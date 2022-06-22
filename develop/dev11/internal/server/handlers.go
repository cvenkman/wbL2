package server

import (
	"fmt"
	"github.com/cvenkman/wbL2/develop/dev11/internal/model"
	"log"
	"net/http"
	"time"
)

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	// create_event?user_id=3&date=2019-09-09
	// returns the first value for the named component of the query
	if r.Method != "POST" {
		log.Println("wrong method, must be post")
		return
	}

	id := r.FormValue("user_id")

	// TODO validate data
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		log.Fatal(err)
	}

	// add event to calendar
	event := model.NewEvent(id, date)
	s.calendar.Add(event)
	fmt.Println("Method:", r.Method, id, s.calendar.Get(id))
}

func (s *Server) GetDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("wrong method, must be post")
		return
	}
	event := s.calendar.Get("2")

	for _, key

	fmt.Println(event.Date.Day())
}
