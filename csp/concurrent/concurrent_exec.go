package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	Workers    = 5
	SplitRange = 10
)

func main() {
	testDemo3()
}

func testDemo1() {
	start, end := 1, 1000000
	before := time.Now()
	fmt.Printf("cal sum from %d to end %d, the res is %d, ", start, end, calSum(start, end))
	fmt.Println("cost time:", time.Since(before))
}

func testDemo2() {
	start, end := 1, 100
	before := time.Now()
	res := 0
	for cur := start; cur < cur+SplitRange && cur <= end; cur += SplitRange {
		tEnd := cur + SplitRange - 1
		if tEnd > end {
			tEnd = end
		}
		res += calSum(cur, tEnd)
	}
	fmt.Printf("cal sum from %d to end %d, the res is %d, ", start, end, res)
	fmt.Println("cost time:", time.Since(before))
}

func testDemo3() {
	start, end := 1, 100
	before := time.Now()
	res := 0
	exectuors := make([]func() error, 0)
	lock := sync.Mutex{}
	for cur := start; cur < cur+SplitRange && cur <= end; cur += SplitRange {
		tEnd := cur + SplitRange - 1
		if tEnd > end {
			tEnd = end
		}
		cur := cur // 闭包函数 必须重新赋值

		exectuors = append(exectuors, func() error {
			t := calSum(cur, tEnd)
			lock.Lock()
			defer lock.Unlock()
			res += t
			return nil
		})

	}
	ConcurrentExecFunc(Workers, exectuors)
	fmt.Printf("cal sum from %d to end %d, the res is %d, ", start, end, res)
	fmt.Println("cost time:", time.Since(before))

}

// 单个任务：累加[start, end]
func calSum(start, end int) int {
	res := 0
	for cur := start; cur <= end; cur++ {
		res += cur
	}
	return res
}

func ConcurrentExecFunc(workers int, execFuncs []func() error) []error {
	errs := make([]error, len(execFuncs))
	if len(execFuncs) == 0 {
		return errs
	}
	ch := make(chan struct{}, workers)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(len(execFuncs))
	for i := range execFuncs {
		ch <- struct{}{}
		go func(i int) {
			defer func() {
				if r := recover(); r != nil {
					err := fmt.Errorf("recover concurrent exec panic, err is %+v", r)
					errs[i] = err
				}
				<-ch
				wg.Done()
			}()
			function := execFuncs[i]
			errs[i] = function()
		}(i)
	}
	wg.Wait()
	return errs
}
