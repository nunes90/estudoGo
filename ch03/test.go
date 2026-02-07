package main

import "fmt"

func main() {

	username := "Sir_King_Über"
	fmt.Println(username)

	c := 0
	for k, v := range username {
		c++
		fmt.Println(username[k], v)
	}

	fmt.Println(c)
}
