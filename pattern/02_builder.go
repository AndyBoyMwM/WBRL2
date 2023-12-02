package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Плюсы паттерна «строитель»:
* Позволяет создавать сложные объекты пошагово, добавляя детали одну за другой.
* Делает код более читаемым и понятным, так как каждая операция создания объекта выполняется отдельной функцией.
* Позволяет изменять порядок шагов создания объекта, если это необходимо.
* Может упростить параллельное создание объектов, так как каждый шаг создания объекта становится отдельной операцией.
* Может помочь в управлении зависимостями между объектами, так как создание объекта разбивается на несколько шагов.
* Изолирует сложный код сборки продукта от его основной бизнес-логики.
* Может использоваться, для создания гибких и расширяемых систем, без существующего изменения кода объектов.

Минусы паттерна «строитель»:
* Может привести к раздуванию кода, если каждая операция добавления детали становится слишком сложной.
* Может усложнить тестирование, так как для каждого шага создания объекта нужно написать отдельный тест.
* Может привести к созданию большого количества мелких объектов, что может негативно сказаться на производительности.
* Если шаги по созданию объекта сильно связаны между собой, то использование паттерна «строитель» может стать менее эффективным.
* Может усложнить понимание того, как создается объект, если шаги создания не описаны явно.
*/

func MakeBuilder() {
	compBuilder := NewComputerBuilder()
	comp := compBuilder.RAM(16).HDD(512).MOB("GIGABYTE Z790").CPU("Intel I5")
	fmt.Println(comp)
}

type ComputerBuilderI interface {
	RAM(val int) ComputerBuilderI
	HDD(val int) ComputerBuilderI
	MOB(val string) ComputerBuilderI
	CPU(val string) ComputerBuilderI
	Build() Computer
}

type Computer struct {
	RAM int
	HDD int
	MOB string
	CPU string
}

type ComputerBuilder struct {
	ram int
	hdd int
	mob string
	cpu string
}

func (c *ComputerBuilder) RAM(val int) ComputerBuilderI {
	c.ram = val
	return c
}

func (c *ComputerBuilder) HDD(val int) ComputerBuilderI {
	c.hdd = val
	return c
}

func (c *ComputerBuilder) MOB(val string) ComputerBuilderI {
	c.mob = val
	return c
}

func (c *ComputerBuilder) CPU(val string) ComputerBuilderI {
	c.cpu = val
	return c
}

func (c *ComputerBuilder) Build() Computer {
	return Computer{
		RAM: c.ram,
		HDD: c.hdd,
		MOB: c.mob,
		CPU: c.cpu,
	}
}

func NewComputerBuilder() ComputerBuilderI {
	return &ComputerBuilder{}
}
