package structdata

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	mu := sync.Mutex{}

	defer func() {
		fmt.Println(" recover .... ")
		if e := recover(); e != nil {
			fmt.Println(fmt.Sprintf("err=%v", e))
		}
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("1 - > 我做完了")
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("2 - > 我做完了")
	}()

	//mu.Unlock() // 这个不能被defer recover
	panic("测试 recover panic")

	fmt.Println("结束.")

}

func TestFoo1(t *testing.T) {
	goExit()
}

func goExit() {
	runtime.Goexit()
	fmt.Println("do not print!!!")
}

// array , map , slice, interface{} , channel , mutex,
