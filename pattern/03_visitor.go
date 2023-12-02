package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

Плюсы паттерна «посетитель»:
* Упрощает добавление операций, работающих со сложными структурами объектов.
* Объединяет родственные операции в одном классе.
* Посетитель может накапливать состояние при обходе структуры элементов.

Минусы паттерна «посетитель»:
* Может усложнить код при большом количестве операций
* Может привести к нарушению инкапсуляции элементов.
* Требует больше времени на разработку и тестирование
* Паттерн не оправдан, если иерархия элементов часто меняется.
*/

func MakeVisitor() {
	visitor := NewTestVisitor()
	visitorShow := VisitorShow{}
	visitor.Visit(visitorShow)
}

type VisitorI interface {
	VisitCountry(country *Country)
	VisitContinent(continent *Continent)
}

type Visitor struct {
	country   Country
	continent Continent
}

type Country struct {
	name                string
	latitude, longitude float64
}

type Continent struct {
	name string
}

type VisitorShow struct{}

func (VisitorShow) VisitCountry(country *Country) {
	fmt.Printf("Visitor in the country: %+v\n", *country)
}

func (VisitorShow) VisitContinent(continent *Continent) {
	fmt.Printf("Visitor on continent: %+v\n", *continent)
}

func (s *Visitor) Visit(v VisitorI) {
	v.VisitCountry(&s.country)
	v.VisitContinent(&s.continent)
}

func NewTestVisitor() *Visitor {
	v := new(Visitor)
	v.country.name = "Cuba"
	v.country.latitude = 21.521757
	v.country.longitude = -77.781167
	v.continent.name = "Africa"
	return v
}
