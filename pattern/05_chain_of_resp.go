package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно
	по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать
	запрос дальше по цепи.


	ПРИМЕНИМОСТЬ

	Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,
	какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	С помощью Цепочки обязанностей вы можете связать потенциальных обработчиков в одну цепь и
	при получении запроса поочерёдно спрашивать каждого из них, не хочет ли он обработать запрос.

	Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	Цепочка обязанностей позволяет запускать обработчиков последовательно один за другим в том порядке,
	в котором они находятся в цепочке.


	ПРЕИМУЩЕСТВА

	Уменьшает зависимость между клиентом и обработчиками.
	Реализует принцип единственной обязанности.
	Реализует принцип открытости/закрытости.


	НЕДОСТАТКИ

	Запрос может остаться никем не обработанным.
*/

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

type Department interface {
	execute(*Patient)
	setNext(Department)
}

type DepartmentBase struct {
	nextDepartment Department
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
