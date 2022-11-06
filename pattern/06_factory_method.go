package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
	позволяя подклассам изменять тип создаваемых объектов.


	ПРИМЕНИМОСТЬ

	Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
	Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.
	Благодаря этому, код производства можно расширять, не трогая основной. Так, чтобы добавить поддержку нового продукта,
	вам нужно создать новый подкласс и определить в нём фабричный метод, возвращая оттуда экземпляр нового продукта.


	ПРЕИМУЩЕСТВА

	Избавляет класс от привязки к конкретным классам продуктов.
	Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	Упрощает добавление новых продуктов в программу.
	Реализует принцип открытости/закрытости.


	НЕДОСТАТКИ

	Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать
	свой подкласс создателя.
*/

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
