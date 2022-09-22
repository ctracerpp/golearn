package structdata

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
	context 应用场景
   1) 值传递
   2) 取消子链路任务
   3) 超时控制
  @see: https://zhuanlan.zhihu.com/p/420127690
*/

func HandelRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}

func TestReadRealMem(t *testing.T) {
	// 读取实时内存占用
	var m runtime.MemStats
	_ = make([]byte, 1024*1024*1024)

	_ = make([]byte, 1024*1024*512)
	_ = make([]byte, 1024*1024*1)
	runtime.ReadMemStats(&m)

	fmt.Println("Memory:", m.Alloc/1024/1024, "MB")
}

// CancelContext 取消控制
func TestWithCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()
	//Just for test whether sub goroutines exit or not 59. time.Sleep(5 * time.Second)
	time.Sleep(5 * time.Second)

}

// TimeOutContext 超时控制
func TestTimeOutContext(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go HandelRequest(ctx)

	time.Sleep(10 * time.Second)
}

func HandelRequestWithParam(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

// ValueContext 传值控制
func TestValueContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go HandelRequestWithParam(ctx)

	time.Sleep(10 * time.Second)
}

// 生成策略模式
type Strategy interface {
	DoSomething()
}

func TestDeadLine(t *testing.T) {
	deadline, _ := context.WithDeadline(context.TODO(), time.Now().Add(3*time.Second))
	go func(c context.Context) {
		time.Sleep(10 * time.Second)
		// cancelFunc()
		fmt.Println("cancelFunc")

	}(deadline)
	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				fmt.Println("Done")
				return
			default:
				fmt.Println("Running")
				time.Sleep(1 * time.Second)
			}

		}
	}(deadline)
	time.Sleep(15 * time.Second)

}

func TestCancelSubGoRoute(t *testing.T) {
	cancel, cancelFunc := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		go func(subCtx context.Context) {
			for {
				select {
				case <-subCtx.Done():
					fmt.Println("Done1")
					return
				default:
					fmt.Println("Running1")
					time.Sleep(1 * time.Second)
				}

			}
		}(ctx)

		go func(subCtx context.Context) {
			for {
				select {
				case <-subCtx.Done():
					fmt.Println("Done2")
					return
				default:
					fmt.Println("Running2")
					time.Sleep(1 * time.Second)
				}

			}
		}(ctx)
	}(cancel)

	go func(ctx context.Context) {
		go func(subCtx context.Context) {
			for {
				select {
				case <-subCtx.Done():
					fmt.Println("Done1-2")
					return
				default:
					fmt.Println("Running1-2")
					time.Sleep(1 * time.Second)
				}

			}
		}(ctx)

		go func(subCtx context.Context) {
			for {
				select {
				case <-subCtx.Done():
					fmt.Println("Done2-2")
					return
				default:
					fmt.Println("Running2-2")
					time.Sleep(1 * time.Second)
				}

			}
		}(ctx)
	}(cancel)

	time.Sleep(3 * time.Second)

	cancelFunc()

	time.Sleep(3 * time.Second)

}

func TestContextPassValue(t *testing.T) {
	value := context.WithValue(context.Background(), "trace-id", "test")
	go func(ctx context.Context) {
		println("sub go ", value.Value("trace-id").(string))
		go func(subCtx context.Context) {
			println("sub1 go ", value.Value("trace-id").(string))
		}(ctx)
		go func(subCtx context.Context) {
			println("sub2 go ", value.Value("trace-id").(string))
		}(ctx)
	}(value)

	time.Sleep(time.Second * 1)
}

func TestContextTimeOut(t *testing.T) {
	timeout, _ := context.WithTimeout(context.Background(), time.Second*1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				println("Done")
				return
			default:
				println("Running")
				time.Sleep(time.Second * 1)
			}
		}
	}(timeout)
	// go route will exit after 3 seconds
	time.Sleep(time.Second * 5)
}
