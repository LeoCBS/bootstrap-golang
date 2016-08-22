package main

import "fmt"

func numberValue(value string) {
	fmt.Println("value:", value)
}

func numberPointer(value *string) {
	fmt.Println("pointer:", value)
}

func main() {
	numberValue("go-scratchpad")
	testString := "go-scratchpad"
	numberPointer(&testString)
}
