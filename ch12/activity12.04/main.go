// activity12.04 - Calculating the future date and time
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Printf("The current time: %s\n", now.Format(time.ANSIC))

	// 6 hours, 6 minutes, 6 seconds = 21966
	sss := time.Duration(21966 * time.Second)

	future := now.Add(sss)

	fmt.Printf("6 hours, 6 minutes and 6 seconds from now the time will be: %s\n", future.Format(time.ANSIC))
}
