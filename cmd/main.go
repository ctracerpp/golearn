package main

import (
	"errors"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/glog"
)

func main() {
	path := "./glog"
	glog.SetPath(path)
	glog.SetStdoutPrint(true)

	glog.Infof("{%s} 测试", "main")
	glog.Infof("{%s} 测试2", " main ")
	if err := testWrapError(); err != nil {
		glog.Error(gerror.Stack(err))
	}
}

func testWrapError() error {
	return gerror.Wrap(errors.New("测试error"), "test")
}
