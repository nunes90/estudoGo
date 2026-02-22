package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type cat struct{}

func main() {
	c := cat{}
	// fmt.Println(c.Speak())
	chatter(c)
}

// `cat` matches the `Speak()` method of the `Speaker{}` interface,
// so a `cat` is a `Speaker{}`
func (c cat) Speak() string {
	return "Purr Meow"
}

func chatter(s Speaker) {
	fmt.Println(s.Speak())
}
