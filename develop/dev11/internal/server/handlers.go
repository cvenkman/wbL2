package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cvenkman/wbL2/develop/dev11/internal/model"
)

const timeLayot = "2006-01-02"

// const errorUnmarshal = "{\"error:\" \"can't unmashal\"}"

// создает новую запись
// curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-09", "title": "fff"}' http://localhost:8080/create_event
func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("wrong method, must be post")
		return
	}

	event, err := unmarshalBody(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// добавляем пришедшее событие в календарь
	s.calendar.Add(event)
	// return data[:n]
}

// curl -i -X POST -H 'Content-Type: application/json' -d '{"user_id": "1", "date": "2019-09-09", "title": "fff"}' http://localhost:8080/update_event
func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("wrong method, must be post")
		return
	}

	event, err := unmarshalBody(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.calendar.Update(event)
}

// выводит все записи за день
// http://localhost:8080/events_for_day?user_id=2&date=2019-09-09
func (s *Server) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("wrong method, must be get")
		return
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
		return
	}
	w.Write(res)
}

// удаляет все записи по определенному user_id и date
func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("wrong method, must be post")
		return
	}
	event, err := unmarshalBody(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.calendar.Delete(event)
}

// читает body post запроса и переводит его в структуру Event
func unmarshalBody(Body io.ReadCloser) (*model.Event, error) {
	// читаем что пришло в body
	data := make([]byte, 1024)
	n, err := Body.Read(data)
	// переводим json в структуру Event
	event, err := model.Unmarshal(data[:n])
	if err != nil {
		return nil, err
	}
	return event, nil
	// for log
	// log.Println(r.Method, event.ID, event.Date, event.Title)
}
