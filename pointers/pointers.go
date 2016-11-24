package pointers

import "fmt"

type Calculator struct {
	Memory int
}

func (calculator *Calculator) Sum(value int) {
	fmt.Println("calculator pointer: %s", &calculator.Memory)
	calculator.Memory = calculator.Memory + 100
}

func (calculator *Calculator) GetMemory() int {
	fmt.Println("calculator pointer: %s", &calculator.Memory)
	return calculator.Memory
}

func NewCalculatorPointer(initValue int) *Calculator {
	return &Calculator{
		Memory: initValue,
	}
}

func NewCalculatorValue(initValue int) Calculator {
	return Calculator{
		Memory: initValue,
	}
}

func (calculator Calculator) SumNoPointer(value int) {
	fmt.Println("calculator pointer: %s", &calculator.Memory)
	calculator.Memory = calculator.Memory + 100
}

func (calculator Calculator) GetMemoryNoPointer() int {
	fmt.Println("calculator pointer: %s", &calculator.Memory)
	return calculator.Memory
}
