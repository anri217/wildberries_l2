package pattern

import "fmt"

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
	позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
	а также поддерживать отмену операций.


	ПРИМЕНИМОСТЬ

	Когда вы хотите параметризовать объекты выполняемым действием.
	Команда превращает операции в объекты. А объекты можно передавать, хранить и взаимозаменять внутри других объектов.
	Скажем, вы разрабатываете библиотеку графического меню и хотите, чтобы пользователи могли использовать меню в разных приложениях,
	не меняя каждый раз код ваших классов. Применив паттерн, пользователям не придётся изменять классы меню, вместо этого они будут
	конфигурировать объекты меню различными командами.


	ПРЕИМУЩЕСТВА

	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	Позволяет реализовать простую отмену и повтор операций.
	Позволяет реализовать отложенный запуск операций.
	Позволяет собирать сложные команды из простых.
	Реализует принцип открытости/закрытости.


	НЕДОСТАТКИ

	Усложняет код программы из-за введения множества дополнительных классов.
*/

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type Device interface {
	on()
	off()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
