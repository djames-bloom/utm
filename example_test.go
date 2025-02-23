package utm_test

import (
	"log"

	"t25.tokyo/utm"
)

func Example() {
	options := []utm.Option{
		utm.Uint("MyUintOption", 12, 0),
		utm.String("MyStringOption", "Hello", ""),
		utm.Int("MyIntOption", 77, 0),
		utm.StringArray{
			Param:  "arrayParam",
			Values: []string{"one", "two", "three"},
		},
	}

	result := utm.ProcessOptions(options)
	log.Println(result)
	// Output:
	// MyUintOption 12 MyStringOption Hello MyIntOption 77 arrayParam one arrayParam two arrayParam three
}
