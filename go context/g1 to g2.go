package main

import (
	"context"
	"fmt"
	"sync"

	"time"
)

var wg2 sync.WaitGroup

func worker3(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知,这个是表示ctx过期或者是取消是，产生的信号。
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker3(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg2.Wait()
	fmt.Println("over")
}
