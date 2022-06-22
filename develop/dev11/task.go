package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cvenkman/wbL2/develop/dev11/internal/config"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

// сервер для сохранения мероприятий (дата, продолжительность, название, имя пользователя)
// доменный объект Event {data Time, title, username string}
// валидация пришедших данных с /create_event и /update_event (не пустые поля, время в правльном формате)

func main() {
	var configPath string
	flag.StringVar(&configPath, "conf", "configs/config.json", "path to config file")
	conf := config.ReadConfig(configPath)
	go startServer(conf)
	fmt.Printf("Server started on http://%s:%s/\n", conf.Host, conf.Port)
	fmt.Println(conf)

	// wait Ctrl+C signal
	// buffer of size 1 because channel used for notification of just one signal
	signalChan := make(chan os.Signal, 1)
	// Ctrl+C and kill
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	fmt.Printf("\nClose connection...\n")

	// Попытка корректного завершения
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// done := make(chan struct{})
	// go func() {
	// 	for range signalChan {
	// 		done <- struct{}{}
	// 	}
	// }()
	// <-done
}

func startServer(conf config.Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет, мир!")
	})
	http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Host, conf.Port), nil)
}
