package model

import "time"

// сервер для сохранения мероприятий (дата, продолжительность, название, имя пользователя)
// доменный объект Event {data Time, title, username string}
// валидация пришедших данных с /create_event и /update_event (не пустые поля, время в правльном формате)

type Event struct {
	Data     time.Time
	Title    string
	Username string
}

// CreateEvent обработчик паттерна /create_event для создания событий в календаре

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		throwError(w, ErrorWrongMethod)
		return
	}

	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		throwError(w, ErrorCanNotParseDate)
		return
	}

	id := r.FormValue("id")
	title := r.FormValue("title")

	event := Event{
		ID:    id,
		Title: title,
		Date:  date,
	}

	events.Store(id, event)

	v, err := json.Marshal(APIResult{APIEventID{ID: id}})
	if err != nil {
		throwError(w, ErrorInternalError)
		return
	}
	fmt.Fprint(w, string(v))
}

// UpdateEvent обработчик паттерна /update_event для изменения/обновления событий в календаре
func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodPost) {
		ev := EventModel{}
		if ok := json.NewDecoder(r.Body).Decode(&ev); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else if ok := s.events.UpdateEvent(ev.ToEvent()); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else {
			jsonResponse(false, w, http.StatusOK, "updated")
		}
	}
}

// DeleteEvent удаляет событие по ID который в теле запроса
func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodPost) {
		ev := EventModel{}
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil || ev.ID == "" {
			jsonResponse(true, w, http.StatusServiceUnavailable, "not parameter ID or isn't correct")
			return
		}
		if ok := s.events.DeleteEvent(ev.ID); !ok {
			jsonResponse(true, w, http.StatusServiceUnavailable, "not found")
			return
		}
		jsonResponse(false, w, http.StatusOK, "deleted")
	}
}

/////// get ///////

// EventsForDay Получаем событие за день
func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodGet, "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else {
			userID := r.URL.Query().Get("user_id")
			evs := s.events.GetEvents(userID, date, date.AddDate(0, 0, 1))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}

// EventsForWeek Получаем событие за неделю
func (s *Server) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodGet, "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusServiceUnavailable, ok.Error())
		} else {
			userID := r.URL.Query().Get("user_id")
			evs := s.events.GetEvents(userID, date, date.AddDate(0, 0, 7))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}

// EventsForMonth Получаем событие за месяц
func (s *Server) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if ValidateQuery(w, r, http.MethodGet, "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			jsonResponse(true, w, http.StatusBadRequest, ok.Error())
		} else {
			userID := r.URL.Query().Get("user_id")
			evs := s.events.GetEvents(userID, date, date.AddDate(0, 1, 0))
			jsonResponse(false, w, http.StatusOK, evs)
		}
	}
}
