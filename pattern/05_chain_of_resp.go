package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Паттерн Chain Of Responsibility относится к поведенческим паттернам уровня объекта.

Паттерн Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к объекту-получателю
запроса, при этом давая шанс обработать этот запрос нескольким объектам. Получатели связываются в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.
*/

//Базовый абстрактный класс Handler, описывающий интерфейс обработчиков в цепочки
type Sender interface {
	SendMsg(string) string
}

//Класс, реализующий конкретный обработчик
type ConcreteSender struct {
}

func (s *ConcreteSender) SendMsg(message string) string {
	return ("send: " + message)
}

//Класс, реализующий конкретный обработчик
type ConcreteSenderSecond struct {
}

func (s *ConcreteSenderSecond) SendMsg(message string) string {
	return ("send second: " + message)
}
