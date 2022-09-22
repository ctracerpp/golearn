package structdata

import (
	"fmt"
	"testing"
)

func deferFuncParameter() {
	var aInt = 1
	defer fmt.Println(aInt)

	aInt = 2
	return
}

func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferFuncParameter1() {
	var aArray = [3]int{1, 2, 3}

	defer printArray(&aArray)

	aArray[0] = 10
	return
}

func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}

func TestDeferVal(t *testing.T) {
	deferFuncParameter()
}

func TestFuncParam(t *testing.T) {
	deferFuncParameter1()
}

func TestFuncReturn(t *testing.T) {
	deferFuncReturn()
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func TestDeferA(t *testing.T) {
	a()
}

func deferFuncReturn2() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}

func TestDeferReturn2(t *testing.T) {
	println(deferFuncReturn2())
}

func foo() int {
	var i int

	defer func() {
		i++
	}()

	return 1
}

func TestDeferFoo(t *testing.T) {
	fmt.Println(foo())
}
