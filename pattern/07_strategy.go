package pattern

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Плюсы паттерна «стратегия»:
* Горячая замена алгоритмов на лету.
* Изолирует код и данные алгоритмов от остальных классов.
* Уход от наследования к делегированию.
* Реализует принцип открытости/закрытости.
* Можно протестировать каждый алгоритм в отдельности.

Минусы паттерна «стратегия»:
* Усложняет программу за счёт дополнительных классов.
* Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

func MakeStrategy() {
	rand.Seed(time.Now().UnixNano())
	arr1, arr2 := make([]int, 10), make([]int, 10)
	for i := range arr1 {
		arr1[i] = rand.Intn(9)
		arr2[i] = rand.Intn(9)
	}

	fmt.Println("arr1 before sort: ", arr1)
	ctx := new(Context)
	ctx.Algorithm(&BubbleSort{})
	ctx.Sort(arr1)
	fmt.Println("arr1 bubble sort: ", arr1)

	fmt.Println("arr2 before sort:   ", arr2)
	ctx.Algorithm(&InsertionSort{})
	ctx.Sort(arr2)
	fmt.Println("arr2 insertion sort:", arr2)
}

type StrategySort interface {
	Sort([]int)
}

type BubbleSort struct {
}

type InsertionSort struct {
}

type Context struct {
	strategy StrategySort
}

func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}

func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}
