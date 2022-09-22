package structdata

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
	fmt.Println("等待2个任务完成")
}
