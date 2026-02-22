// This code has a lot of redundant functions that perform similar actions.

// Continue on part3 - polymorphism3/main.go

package main

import "fmt"

type Speaker interface {
	Speak() string
}

// concrete type
type cat struct{}

// concrete type
type dog struct{}

// concrete type
type person struct {
	name string
}

func main() {
	c := cat{}
	d := dog{}
	p := person{name: "Heather"}
	catSpeak(c)
	dogSpeak(d)
	personSpeak(p)
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func (d dog) Speak() string {
	return "Woof Woof"
}

func (p person) Speak() string {
	return "Hi my name is " + p.name + "!"
}

func catSpeak(c cat) {
	fmt.Println(c.Speak())
}

func dogSpeak(d dog) {
	fmt.Println(d.Speak())
}

func personSpeak(p person) {
	fmt.Println(p.Speak())
}
