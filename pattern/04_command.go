package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Плюсы паттерна «комманда»:
* Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
* Позволяет реализовать простую отмену и повтор операций.
* Позволяет реализовать отложенный запуск операций.
* Позволяет собирать сложные команды из простых.
* Реализует принцип открытости/закрытости.

Минусы паттерна «комманда»:
* Усложняет код программы из-за введения множества дополнительных классов.
* Команды могут привести к избыточности кода, если их использовать слишком часто.
* Создание и тестирование команд может быть трудоемким процессом.
*/

func MakeCommand() {
	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})

	result := invoker.Execute()
	fmt.Println(result)
}

type Command interface {
	Execute() string
}

type Receiver struct {
}

type ToggleOnCommand struct {
	receiver *Receiver
}

type Invoker struct {
	commands []Command
}

func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Receiver
}

func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}
