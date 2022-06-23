package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Паттерн State относится к поведенческим паттернам уровня объекта.

Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния и
является объектно-ориентированной реализацией конечного автомата.
Поведение объекта изменяется настолько, что создается впечатление, будто изменился класс объекта.
*/

// Абстрактный класс State, определяющий интерфейс различных состояний
type State interface {
	Alert() string
}

// MobileAlert implements an alert depending on its state.
type MobileAlert struct {
	state State
}

// Alert returns a alert string
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState changes state
func (a *MobileAlert) SetState(state State) {
	a.state = state
}

// NewMobileAlert is the MobileAlert constructor.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

/****  Классы, которые реализуют одно из поведений, ассоциированное с определенным состоянием  *****/

// MobileAlertVibration implements vibration alert
type MobileAlertVibration struct {
}

// Alert returns a alert string
func (a *MobileAlertVibration) Alert() string {
	return "..."
}

// MobileAlertSong implements beep alert
type MobileAlertSong struct {
}

// Alert returns a alert string
func (a *MobileAlertSong) Alert() string {
	return "***"
}
