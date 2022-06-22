package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cvenkman/wbL2/develop/dev11/internal/model"
)

const timeLayot = "2006-01-02"
const errorUnmarshal = "{\"error:\" \"can't unmashal\"}"

// curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-09", "title": "fff"}' http://localhost:8080/create_event
func (s *Server) Create(w http.ResponseWriter, r *http.Request) []byte {
	if r.Method != "POST" {
		log.Println("wrong method, must be post")
		return nil
	}

	// читаем что пришло в body
	data := make([]byte, 1024)
	n, err := r.Body.Read(data)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// переводим json в структуру Event
	event, err := model.Unmarshal(data[:n])
	if err != nil {
		fmt.Println(err)
		return []byte(errorUnmarshal)
	}
	log.Println(r.Method, event.ID, event.Date, event.Title)

	// добавляем пришедшее событие в календарь
	s.calendar.Add(event)
	return data[:n]
}

// http://localhost:8080/events_for_day?user_id=2&date=2019-09-09
func (s *Server) GetEventsForDay(w http.ResponseWriter, r *http.Request) []byte {
	if r.Method != "GET" {
		log.Println("wrong method, must be post")
		return nil
	}
	// берем данные из query
	user_id := r.FormValue("user_id")
	day, err := time.Parse(timeLayot, r.FormValue("date"))
	if err != nil {
		log.Fatal(err)
	}
	// берем все данные за день у пользователя с user_id
	dayEvents := s.calendar.GetEventsForDay(user_id, day)

	// переводим массив Event в json
	res, err := json.Marshal(dayEvents)
	if err != nil {
		log.Println(err)
		return nil
	}
	w.Write(res)
	return res
}
