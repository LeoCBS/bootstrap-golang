package pointers

import "fmt"

func SumValue(value int) int {
	return value + 100
}

func numberPointer(value *string) {
	fmt.Println("pointer:", value)
}

func loadvalueFromPointer(value *string) {
	fmt.Println("value from point: %s", value)
}
