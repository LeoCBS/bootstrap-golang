# pointers

The code have one very simple application that runs a calculation to demonstrate diferrent in use pointer or value.



below calculator created by pointer:

```
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

```

 it meens when we call method Sum and GetMemory, we reference same memory pointer and get same value from attribute Memory:

 pointer_test.go

 ```
 func TestSumPointer(t *testing.T) {
	calculator := pointers.NewCalculatorPointer(100)
    calculator.Sum(100)
    value := calculator.GetMemory()
    fmt.Printf("value " , value)
    if value != 200{
        t.Fatal("calculator don't sum value")
    }
}
 ```

 test ouput:

 ```
 ?   	github.com/leocbs/go-scratchpad	[no test files]
=== RUN   TestSumPointer
value %!(EXTRA int=200)--- PASS: TestSumPointer (0.00s)

 ```

 Now, when we created calculator and return without pointer (&), go copy value to another pointer and don't use same pointer:

 ```
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
 ```

 and when we get memory value, it's continue 100:

 pointer_test.go:

 ```
 func TestSumPerValue(t *testing.T) {
	calculator := pointers.NewCalculatorValue(100)
    calculator.SumNoPointer(100)
    value := calculator.GetMemoryNoPointer()
    fmt.Printf("value " , value)
    if value != 100{
        t.Fatal("calculator sum value")
    }
}
 ```

 test ouput:

 ```
 === RUN   TestSumPerValue
value %!(EXTRA int=100)--- PASS: TestSumPerValue (0.00s)
 ```

 To run this test, just clone and `make run`

 Now that we already know pointer behavior, read this [this](http://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values) question in stack over flow to learn when user pointer or value

