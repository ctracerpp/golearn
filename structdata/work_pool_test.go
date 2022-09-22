package structdata

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// go run main.go -race
// 停止的时候，任务还未执行完成
/**
	简单线程池：
     提供功能，同时执行的任务数,任务执行开始时间，耗时，等待超时
     todo
     1.优先级
     2.动态添加任务
     3.动态删除任务
     4.动态调整池大小
     5.任务执行失败，重试
	 6.优雅的关闭chanel
*/

type jobInfo struct {
	id        interface{}
	startTime time.Time
	endTime   time.Time
	costTime  time.Duration
}

type WorkPool struct {
	total int
	// name of the work pool
	name string
	// 工作池大小
	size     int
	tasks    []interface{}
	jobs     chan interface{}
	results  chan interface{}
	task     func(interface{}) error
	statMaps map[int]map[interface{}]jobInfo
	stopCh   chan bool
	doneCh   chan bool
	mapLock  sync.Mutex
}

func (w *WorkPool) Init(name string, size int, task func(interface{}) error) {
	w.name = name
	w.size = size
	w.jobs = make(chan interface{}, size)
	w.results = make(chan interface{}, size)
	w.task = task
	w.tasks = make([]interface{}, 0)
	w.statMaps = make(map[int]map[interface{}]jobInfo)
	w.stopCh = make(chan bool, 1)
	w.doneCh = make(chan bool, 1)
	for i := 0; i < w.size; i++ {
		w.statMaps[i+1] = make(map[interface{}]jobInfo)
		go worker(w.statMaps, w.name, i+1, w.jobs, w.results, w.task, &w.mapLock)
	}
}

func (w *WorkPool) SubmitJob(job interface{}) {
	w.total++
	w.tasks = append(w.tasks, job)
}

// param timeoutMs 超时时间 ,0 表示不超时
func (w *WorkPool) WaitResult(timeout time.Duration) error {
	go func(jobs []interface{}) {
		for _, job := range jobs {
			w.jobs <- job
		}
		// close(w.jobs)
	}(w.tasks)
	var errList = make([]error, 0)
	go func(that *WorkPool, list []error) {
		for i := 0; i < that.total; i++ {
			select {
			case err := <-that.results:
				if err != nil {
					list = append(list, err.(error))
				}
			case <-w.stopCh:
				break
			}
		}
		w.doneCh <- true

	}(w, errList)
	t := timeout

	if timeout <= 0 {
		t = time.Duration(1<<63 - 1)
	}
	fmt.Println("timeout ", t.Seconds())
	select {
	case <-w.doneCh:
	case <-time.After(t):
		w.stopCh <- true
	}
	defer close(w.stopCh)
	//defer close(w.doneCh)
	//defer close(w.jobs)
	//defer close(w.results)
	if len(errList) > 0 {
		return errList[0]
	}
	return nil
}

// Stat 统计各工作线程执行的任务情况，耗时情况
func (w *WorkPool) Stat() {
	w.mapLock.Lock()
	defer w.mapLock.Unlock()
	for k, m := range w.statMaps {
		var count int64 = 0
		var avgTime = 0.0
		for _, v := range m {
			count++
			avgTime = avgTime + float64(v.costTime.Milliseconds())
		}
		if count > 0 {
			avgTime = avgTime / float64(count)
		}
		fmt.Println("worker ", k, " counter ", count, " avgTime ", avgTime)
	}

}

func (w *WorkPool) Stop() {
	//
}

// param statMaps 统计信息
func worker(stats map[int]map[interface{}]jobInfo, token string, id int, jobs <-chan interface{},
	results chan<- interface{}, pf func(interface{}) error, mux *sync.Mutex) {
	for j := range jobs {
		start := time.Now()
		fmt.Println("pool ", token, " worker ", id, " processing job ", j, " begin ")
		result := pf(j)
		time.Sleep(time.Second)
		end := time.Now()

		// SfLog().INFOformat("pool %s worker %d  processing job %+v end ", token, id, j)
		cost := time.Since(start)
		mux.Lock()
		stats[id][j] = jobInfo{
			id:        j,
			startTime: time.Now(),
			endTime:   end,
			costTime:  cost,
		}
		mux.Unlock()
		fmt.Println("pool ", token, " worker ", id, " processing job ", j, " done ")
		results <- result
	}
}

func TestWorkPool(t *testing.T) {
	var wp = WorkPool{}
	wp.Init("test", 3, func(i interface{}) error {
		return nil
	})
	for i := 0; i < 17; i++ {
		wp.SubmitJob(i)
	}
	_ = wp.WaitResult(3011 * time.Millisecond)
	// _ = wp.WaitResult(1001 * time.Millisecond)
	// wp.Stop()
	wp.Stat()

	// context.WithCancel(context.Background())
	// 查看内存大小、CPU占用率、goroutine数量
	// go tool pprof -alloc_space -inuse_space -inuse_objects -alloc_objects -sample_index=alloc_space -sample_index=inuse_space -sample_index=inuse_objects -sample_index=alloc_objects -top -cum -seconds 10 http://localhost:6060/debug/pprof/heap
}

func TestPoolSize(t *testing.T) {
	var wp = WorkPool{}
	wp.Init("test", 4, func(i interface{}) error {
		return nil
	})
	for i := 0; i < 55; i++ {
		wp.SubmitJob(i)
	}
	_ = wp.WaitResult(3011 * time.Millisecond)
	wp.Stat()

}
