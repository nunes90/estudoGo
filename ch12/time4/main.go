package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	onlyAfter, err := time.Parse(time.RFC3339, "2020-11-01T22:08:42+00:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now, onlyAfter)
	fmt.Println(now.After(onlyAfter))

	// .After()` retorna `true` se `now` é **posterior** a `onlyAfter`. Como a data parseada é 2020, sempre retornará `true`.
	if now.After(onlyAfter) {
		fmt.Println("Executin actions!")
	} else {
		fmt.Println("Now is not the time yet!!!")
	}
}
