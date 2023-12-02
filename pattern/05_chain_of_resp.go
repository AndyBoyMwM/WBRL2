package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Плюсы паттерна «цепочка вызовов»:
* Уменьшает зависимость между клиентом и обработчиками, каждый элемент в цепочке отвечает только за свою часть задачи.
  Это делает код более читаемым и облегчает его понимание.
* Добавление или удаление элементов из цепочки не требует изменений в других элементах.
* Реализует принцип единственной обязанности.
* Реализует принцип открытости/закрытости.

Минусы паттерна «цепочка вызовов»:
* Если цепочка содержит много элементов, это может усложнить управление ими.
* Передача запроса по всей цепочке может снизить производительность системы,
особенно если элементы в цепочке выполняют тяжелые операции.
* Запрос может остаться никем не обработанным.
*/

func MakeChainOfResponsibility() {
	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}
	result := handlers.SendRequest(2)
	fmt.Println(result)
}

type Handler interface {
	SendRequest(message int) string
}

type ConcreteHandlerC struct {
	next Handler
}

type ConcreteHandlerA struct {
	next Handler
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}
