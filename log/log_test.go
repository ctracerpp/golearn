package log

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestName(t *testing.T) {
	g.Log("test1").Infof("%s 测试", "log ")
}
