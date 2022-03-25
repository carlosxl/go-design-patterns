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

type MallardDuck struct{ Duck }

func (d *MallardDuck) display() {
	fmt.Println("draw a mallard duck")
}

type RubberDuck struct{ Duck }

func (d *RubberDuck) display() {
	fmt.Println("draw a rubber duck")
}

func main() {
	d1 := MallardDuck{}
	d1.quack()
	d1.swim()
	d1.display()

	d2 := RubberDuck{}
	d2.quack()
	d2.swim()
	d2.display()
}
