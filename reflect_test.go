package golearn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect2(t *testing.T) {
	var x float64 = 3.4

	v := reflect.ValueOf(x) //v is reflext.Value

	var y float64 = v.Interface().(float64)
	fmt.Println("value:", y)
}

// fault : panic: reflect: reflect.Value.SetFloat using unaddressable value [recovered]
//	panic: reflect: reflect.Value.SetFloat using unaddressable value
func TestReflect3_1(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1) // Error: will panic. 值不能修改
}

func TestReflect3_2(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)
	fmt.Println("x :", v.Elem().Interface())
}
