// ex7.03 - analyzing empty interface{} data
// The any keyword as of Go 1.18 is basically an alias to interface{}.
package main

import "fmt"

type record struct {
	key       string // name of the map key
	valueType string // type of data stored
	data      any    // data itself
}

type person struct {
	lastName  string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

func main() {
	m := make(map[string]any)
	a := animal{name: "oreo", category: "cat"}
	p := person{lastName: "Doe", isMarried: false, age: 19}

	m["person"] = p
	m["animal"] = a
	m["age"] = 54
	m["isMarried"] = true
	m["lastName"] = "Smith"

	rs := []record{}
	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}

	for _, v := range rs {
		fmt.Println("Key: ", v.key)
		fmt.Println("Data: ", v.data)
		fmt.Println("Type: ", v.valueType)
		fmt.Println()
	}

}

// any - interface{}
func newRecord(key string, i any) record {
	r := record{}
	r.key = key
	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
	case bool:
		r.valueType = "bool"
		r.data = v
	case string:
		r.valueType = "string"
		r.data = v
	case person:
		r.valueType = "person"
		r.data = v
	default:
		r.valueType = "unknown"
		r.data = v
	}
	return r
}
