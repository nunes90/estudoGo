// activity9.01 - consuming multiple modules
package main

import (
	"fmt"
	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generated UUID: %s\n", id)

	rq := quote.Go()
	fmt.Println(rq)
}
