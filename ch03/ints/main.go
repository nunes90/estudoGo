/*
 This code creates a collection of int or int8 numbers. It then adds 10 million values to the collection. Once that’s done, it uses the runtime package to give us a reading of how much heap memory is being used. We can convert that reading to MB and then print it out:
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// var list []int
	var list []int8
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v Mib\n", m.TotalAlloc/1024/1024)
}
