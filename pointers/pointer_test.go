package main_test

import (
	"main"
	"testing"
)

func TestSumWithoutPointer(t *testing.T) {
	main.sumValue(100)
}
