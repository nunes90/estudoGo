// When a function accepts an interface as an input parameter, any concrete type that implements that interface can be passed as an argument. Now, you have achieved polymorphism by being able to pass various concrete types to a method or function that has an interface type as an input parameter.

package main

import "fmt"

type Speaker interface {
	Speak() string
}

func saySomething(say ...Speaker) {
	for _, s := range say {
		fmt.Println(s.Speak())
	}
}

type cat struct{}

func (c cat) Speak() string {
	return "Purr Meow"
}

type dog struct{}

func (d dog) Speak() string {
	return "Woof Woof"
}

type person struct {
	name string
}

func (p person) Speak() string {
	return fmt.Sprintf("Hello, my name is %s", p.name)
}

func main() {
	c := cat{}
	d := dog{}
	p := person{name: "Heather"}
	saySomething(c, d, p)
}
