package structdata

import (
	"fmt"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestRuneLength1(t *testing.T) {
	address := "this is hangzhou"
	fmt.Println("len(address):", len(address))
}

func TestRuneLength2(t *testing.T) {
	address := "this is hangzhou"
	fmt.Println("len(address):", len(address))
}

func TestRuneLength3(t *testing.T) {
	addressThree := "这是在杭州"
	fmt.Println("len(address):", utf8.RuneCountInString(addressThree))
}

func TestRuneLength4(t *testing.T) {
	addressThree := "这是在杭州"
	fmt.Println("len(address):", len([]rune(addressThree)))
}

func TestRune5(t *testing.T) {
	c := '\u0000'
	is := unicode.Is(unicode.Han, c) //可以判断字符是否是汉语
	if is {
		fmt.Println("是汉字")
	} else {
		fmt.Println("不是汉字")
	}
}

func TestForEach(t *testing.T) {
	r := "这次真的ai你"
	fmt.Println("len(r):", len(r)) //result is 17
	rArray := []rune(r)
	// i:0,i2:这
	//i:3,i2:次
	//i:6,i2:真
	//i:9,i2:的
	//i:12,i2:a
	//i:13,i2:i
	//i:14,i2:你
	for i, i2 := range r {
		fmt.Println(fmt.Sprintf("i:%d,i2:%c", i, i2))
	}

	// some result is :
	//这
	//次
	//真
	//的
	//a
	//i
	//你
	for _, v := range rArray {
		fmt.Println(fmt.Sprintf("%c", v))
	}

}
