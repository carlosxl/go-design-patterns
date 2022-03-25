package main

import "fmt"

type flyBehavior interface {
	fly()
}

type flyWithWings struct {}
func (f *flyWithWings) fly() {
	fmt.Println("i'm really flying")
}

type flyNoWay struct {}
func (f *flyNoWay) fly() {
	fmt.Println("i'm can't fly")
}

type Duck struct{
	flyBehavior flyBehavior
}

func NewDuck(flyBehavior flyBehavior) *Duck {
	return &Duck{flyBehavior: flyBehavior}
}

func (d *Duck) quack() {
	fmt.Println("quack")
}

func (d *Duck) swim() {
	fmt.Println("swim")
}

func (d *Duck) display() {
	fmt.Println("draw a duck")
}

func (d *Duck) fly() {
	d.flyBehavior.fly()
}

type MallardDuck struct{ *Duck }
func NewMallardDuck(flyBehavior flyBehavior) *MallardDuck {
	return &MallardDuck{Duck: NewDuck(flyBehavior)}
}

type RubberDuck struct{ *Duck }
func NewRubberDuck(flyBehavior flyBehavior) *RubberDuck {
	return &RubberDuck{Duck: NewDuck(flyBehavior)}
}

// ... all different kinds of ducks

func main() {
	d1 := NewMallardDuck(new(flyWithWings))
	d1.quack()
	d1.swim()
	d1.display()
	d1.fly()

	d2 := NewRubberDuck(new(flyNoWay))
	d2.quack()
	d2.swim()
	d2.display()
	d2.fly()
}
