package pointers

type Calculator struct{
	Memory int
}

func (calculator *Calculator) Sum(value int) {
	calculator.Memory = calculator.Memory + 100
}

func (calculator *Calculator) GetMemory() int {
	return calculator.Memory
}

func NewCalculatorPointer(initValue int) *Calculator{
	return &Calculator{
		Memory:initValue,
	}
}

func NewCalculatorValue(initValue int) Calculator{
	return Calculator{
		Memory:initValue,
	}
}

func (calculator Calculator) SumNoPointer(value int) {
	calculator.Memory = calculator.Memory + 100
}

func (calculator Calculator) GetMemoryNoPointer() int {
	return calculator.Memory
}
