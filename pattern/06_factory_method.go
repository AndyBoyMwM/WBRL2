package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Плюсы паттерна «фабричный метод»:
* Позволяет скрыть детали создания объектов от клиентов, предоставляя им общий интерфейс создания объектов.
* Позволяет использовать одну и ту же фабричную функцию для создания разных объектов, определяя фабричные методы в подклассах.
* Выделяет код производства продуктов в одно место, упрощая поддержку кода.
* Упрощает добавление новых продуктов в программу.

Минусы паттерна «фабричный метод»:
* Может добавить дополнительную сложность в код, особенно если фабрика имеет сложную иерархию подклассов.
* Может привести к созданию большого количества фабричных классов, если каждый подкласс создает свой собственный объект.
*/

func MakeFactoryMethod() {
	factory := NewCreator()
	products := []Product{
		factory.CreateProduct(A),
		factory.CreateProduct(B),
		factory.CreateProduct(C),
	}
	for _, product := range products {
		fmt.Println(product)
	}
}

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

type action string

type Creator interface {
	CreateProduct(action action) Product
}

type Product interface {
	Use() string
}

type ConcreteCreator struct{}

type ConcreteProductA struct {
	action string
}

type ConcreteProductB struct {
	action string
}

type ConcreteProductC struct {
	action string
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	case C:
		product = &ConcreteProductC{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}
	return product
}

func (p *ConcreteProductA) Use() string {
	return p.action
}

func (p *ConcreteProductB) Use() string {
	return p.action
}

func (p *ConcreteProductC) Use() string {
	return p.action
}
