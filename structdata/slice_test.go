package structdata

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestSlice(t *testing.T) {
	var array [10]int // output : 0,0,0,0,0,0,0,0,0,0
	for i, i2 := range array {
		fmt.Println("array i ", i, " value ", i2)
	}
	var slice = array[5:6]                         // copy from index 5（包含） and end with 6(不包含) ,切片的值改了，原来的值也会改变
	fmt.Println("lenth of slice: ", len(slice))    // output 1
	fmt.Println("capacity of slice: ", cap(slice)) // output 5
	slice[0] = 1
	fmt.Println("array[5]=", array[5]) // output
	array[5] = 100
	fmt.Println("slice[0] =", slice[0])

	fmt.Println(&slice[0] == &array[5]) //output : true
}

func TestSliceAppend(t *testing.T) {
	slice1 := make([]int, 4)
	fmt.Println(&slice1)
	slice1 = append(slice1, 1, 2, 3)
	slice2 := slice1[1:2]
	fmt.Println(cap(slice1))
	fmt.Println(&slice1)
	fmt.Println(cap(slice2))
	fmt.Println(slice2)

	t1 := time.Now()
	chan1 := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 3)
		chan1 <- 1
		fmt.Println("go routine chan1")
	}()
	<-chan1
	t2 := time.Since(t1)
	fmt.Println(t2)

}

func TestStringSlice(t *testing.T) {
	// 如果是直接字符串截取可能出现乱码，用rune类型 ,一个中文字符3个字节
	str := "测试123"
	s := []rune(str)
	// the index is start from 0
	idx1 := strings.IndexRune(str, '2') // output 7 ,按照字节的index
	fmt.Println("idx1 = ", idx1)
	idx2 := strings.IndexRune(str, '试') // output 3 , 按照字节的index
	fmt.Println("idx2 = ", idx2)
	idx3 := strings.Index(str, "12") // output 6 , 按照字节的index
	fmt.Println("idx3 = ", idx3)
	s1 := s[2:4] // from index to end (不包含) output : 12
	fmt.Println(string(s1))
	s2 := s[1:] // output : 试123
	fmt.Println(string(s2))

	// 测试this is a test
	// strings.NewReader("this is a test")
}
