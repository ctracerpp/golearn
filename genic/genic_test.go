package genic

import (
	"fmt"
	"testing"
)

func AddSum[T int | float64](params ...T) (sum T) {
	// 一些操作
	return sum
}

// 1. 泛型的类型限制，在函数上直接申明该函数支持的多个类型, 操作符重载
func AddElem[T int | string](params []T) (sum T) {
	for _, elem := range params {
		sum += elem
	}
	return
}

func TestGenic(t *testing.T) {
	intSum := AddElem([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	t.Logf("测试1.1: 类型=%T，val=%+v", intSum, intSum)

	// 1.2 传入支持的string
	strSum := AddElem([]string{"静", "以", "修", "身", "，", "俭", "以", "养", "德"})
	strSum1 := AddElem([]string{"格物", "致知", "诚意", "正心", "，", "修身", "治国", "平天下"})
	fmt.Println(strSum1)
	t.Logf("测试1.2: 类型=%T，val=%+v", strSum, strSum)

}

func TestPanic(t *testing.T) {

}
