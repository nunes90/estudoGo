// activity12.01 - Formatting a date according to user requirements
package main

import (
	"fmt"
	"strconv"
	"time"
)

// takes the current date
// output: 15:32:30 2023/10/17

func main() {

	currentTime := time.Now()

	//-------------------------------------------------------------------------
	// day := strconv.Itoa(currentTime.Day())
	day := fmt.Sprintf("%02d", currentTime.Day())
	// month := strconv.Itoa(int(currentTime.Month()))
	month := fmt.Sprintf("%02d", int(currentTime.Month()))
	year := strconv.Itoa(currentTime.Year())
	hour := strconv.Itoa(currentTime.Hour())
	minute := strconv.Itoa(currentTime.Minute())
	second := strconv.Itoa(currentTime.Second())

	fmt.Println(hour + ":" + minute + ":" + second + " " + year + "/" + month + "/" + day)

	//-------------------------------------------------------------------------
	day2 := currentTime.Day()
	month2 := currentTime.Month()
	year2 := currentTime.Year()
	hour2 := currentTime.Hour()
	minute2 := currentTime.Minute()
	second2 := currentTime.Second()

	fmt.Printf("%02d:%02d:%02d %02d/%02d/%02d\n", hour2, minute2, second2, year2, month2, day2)

	//-------------------------------------------------------------------------
	fmt.Println(currentTime.Format("15:04:05 2006/01/02"))

}
