package main

import "fmt"

type Duck struct{}

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
	fmt.Println("I'm flying")
}

type DabblingDuck struct{ Duck }
type GoldenEyeDuck struct{ Duck }
type MallardDuck struct{ Duck }
type RedheadDuck struct{ Duck }
type RubberDuck struct{ Duck }

// ... all different kinds of ducks

func main() {
	d1 := MallardDuck{}
	d1.quack()
	d1.swim()
	d1.display()
	d1.fly()

	d2 := RubberDuck{}
	d2.quack()
	d2.swim()
	d2.display()
	d2.fly()
}
