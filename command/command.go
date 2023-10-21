package main

import "fmt"

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type OpenDoor struct {
	car Car
}

func (od *OpenDoor) execute() {
	od.car.open()
}

type CloseDoor struct {
	car Car
}

func (cd *CloseDoor) execute() {
	cd.car.close()
}

type Alarm struct {
	car Car
}

func (a *Alarm) execute() {
	a.car.alarm()
}

type Car interface {
	open()
	close()
	alarm()
}

type ConcreteCar struct {
	isOpen bool
}

func (c *ConcreteCar) open() {
	fmt.Println("Car is open")
}

func (c *ConcreteCar) close() {
	fmt.Println("Car is closed")
}

func (c *ConcreteCar) alarm() {
	fmt.Println("ALARM!!!")
}

func main(){
	merc := &ConcreteCar{}

	closeDoor := &Button{command: &CloseDoor{car:merc}} 
	openDoor :=&Button{command: &OpenDoor{car:merc}}
	alarm :=&Button{command: &Alarm{car:merc}}

	closeDoor.press()
	openDoor.press()
	alarm.press()
}