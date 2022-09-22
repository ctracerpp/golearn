package structdata

import (
	"fmt"
	"testing"
	"time"
)

//https://www.kancloud.cn/one_go/golang/2572156
// ChanRandomSelect 会随机选择一个准备好的信道进行执行
func TestChanRandomSelect(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		//time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 1

		//time.Sleep(5 * time.Second)
	}()

	select {

	case <-chan1:

		fmt.Println("chan1 ready.")

	case <-chan2:

		fmt.Println("chan2 ready.")

	default:

		fmt.Println("default")

	}
	fmt.Println("main exit.")
}

func TestChanBlock(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)
	writeFlag := false
	go func() {
		for {
			if writeFlag {
				chan1 <- 1
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			if writeFlag {
				chan2 <- 1
			}
			time.Sleep(time.Second)

		}
	}()
	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}
	fmt.Println("main exit.")
}

// CloseChanReadVal 向关闭的chan读取数据不会阻塞
func TestCloseChanReadNoBlock(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		close(chan1)
	}()
	go func() {
		close(chan2)
	}()
	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}
	fmt.Println("main exit.")
}
