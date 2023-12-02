package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Плюсы паттерна «состояние»:
* Объект может изменять свое поведение динамически в зависимости от состояния.
* Избавляет от множества больших условных операторов машины состояний.
* Концентрирует в одном месте код, связанный с определённым состоянием
* Легко тестировать, так как можно изолировать состояние и проверить его поведение.

Минусы паттерна «состояние»:
* Добавление новых состояний может усложнить систему.
* Может неоправданно усложнить код, если состояний мало и они редко меняются.
* Если состояние не переключается корректно, система может работать неправильно
*/

func MakeState() {
	laptop := NewLaptopStatus()
	result := laptop.State()
	laptop.SetStatus(&LaptopSleep{})
	result += laptop.State()
	laptop.SetStatus(&LaptopWork{})
	result += laptop.State()
	fmt.Println(result)
}

type LaptopLauncherI interface {
	Status() string
}

type LaptopStatus struct {
	state LaptopLauncherI
}

type LaptopWork struct {
}

type LaptopSleep struct {
}

func (a *LaptopStatus) State() string {
	return a.state.Status()
}

func (a *LaptopStatus) SetStatus(state LaptopLauncherI) {
	a.state = state
}

func NewLaptopStatus() *LaptopStatus {
	return &LaptopStatus{state: &LaptopWork{}}
}

func (a *LaptopWork) Status() string {
	return "Work"
}

func (a *LaptopSleep) Status() string {
	return "Sleep"
}
