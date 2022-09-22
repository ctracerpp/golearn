package structdata

import (
	"fmt"
	"strconv"
	"testing"
)

// http://t.zoukankan.com/fanbi-p-10928965.html

func TestIntToStr(t *testing.T) {
	var i int = 1
	s := strconv.Itoa(i)
	fmt.Println(s)

}
