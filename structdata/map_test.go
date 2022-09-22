package structdata

import (
	"fmt"
	"testing"
)

func TestMapEach(t *testing.T) {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3

	for k, v := range m {
		fmt.Println("k", k, "v ", v)
	}
	// output: 遍历是无序的
	// k 3 v  3
	// k 1 v  1
	// k 2 v  2
}

func TestAppend(t *testing.T) {
	m := make(map[int]int)
	// m[1] = 1
	if _, ok := m[1]; !ok {
		fmt.Println("m[1] = 1")
		m[1] = 1
	}
}

// TestRemove1 第一种方法删除，重建
func TestRemove1(t *testing.T) {
	a := make(map[string]int, 10000000)

	a["a"] = 1
	a["b"] = 2

	// clear all 122ms
	a = make(map[string]int)
}

func TestRemove2(t *testing.T) {
	// 通过Go的内部函数mapclear方法删除。这个函数并没有显示的调用方法，当你使用for循环遍历删除所有元素时，Go的编译器会优化成Go内部函数mapclear
	m := make(map[byte]int, 10000000)

	m[1] = 1
	m[2] = 2
	// 144ms
	for k := range m {
		delete(m, k)
	}
}

type refObj struct {
	name string
}

func TestMapHoleRefObj(t *testing.T) {
	m := make(map[string]*refObj)
	obj1 := refObj{name: "1"}
	obj2 := obj1
	obj2.name = "2"
	m["1"] = &obj1
	m["2"] = &obj2
	// output address
	// s 1 obj 0xc0000060a0
	// s 2 obj 0xc0000060a0
	// output name
	// s 1 obj 1
	// s 2 obj 2
	for s, obj := range m {
		fmt.Println("s", s, "obj", obj.name)
	}
}

func cloneTags(tags map[string]string) map[string]string {
	cloneTags := make(map[string]string)
	for k, v := range tags {
		cloneTags[k] = v
	}
	return cloneTags
}

// 值拷贝
func TestMapCopy(t *testing.T) {
	m := make(map[string]string)
	m["1"] = "1"
	m["2"] = "2"
	rm := cloneTags(m)
	for s, s2 := range rm {
		fmt.Println("key", s, "value", s2)
	}
	m["2"] = "modify2"
	// output
	// key 1 value 1
	// key 2 value 2
	for s, s2 := range rm {
		fmt.Println("key", s, "value", s2)
	}
}

func TestMapDeep(t *testing.T) {

}
