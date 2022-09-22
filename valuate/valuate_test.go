package valuate

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strings"
	"testing"
)

/**
@see https://github.com/Knetic/govaluate
*/
func TestValuate1(t *testing.T) {
	fmt.Println("组织表达式...")
	expression, err := govaluate.NewEvaluableExpression("foo > 2 ? '真a':'假a'")
	if err != nil {
		panic(err)
	}
	fmt.Println("组织参数...")
	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = 4

	result, err := expression.Evaluate(parameters)
	if err != nil {
		panic(err)
	}
	fmt.Println("计算结果...")
	fmt.Println(result)
	var lowwer = result.(string)
	// 转换为大写
	fmt.Println(strings.ToUpper(lowwer))
	// result is now set to "false", the bool value.
}

// go内存标记，三色法
