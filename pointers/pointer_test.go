package pointers_test

import (
	"testing"
    "fmt"

    "github.com/leocbs/go-scratchpad/pointers"
)

func TestSumWithoutPointer(t *testing.T) {
    value := pointers.SumValue(100)
    fmt.Printf("value " , value)
    if value == 1000{
        t.Fatal("Pointer don't work")
    }
}