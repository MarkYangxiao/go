package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go handler(ctx)
	time.Sleep(10 * time.Second) // 假设运行10s后，前端用户取消请求
	cancelRequest(cancel)
}

func cancelRequest(cancelFunc context.CancelFunc) {
	cancelFunc()
}

// 每隔2秒返回计算结果
func handler(ctx context.Context) {
	for {
		fmt.Println("正在处理请求...")
		fmt.Println("发送计算结果...")
		select {
		case <-ctx.Done():
			// 请求被被取消，返回c
			fmt.Println("请求被取消：原因:", ctx.Err().Error())
			return
		case <-time.After(2 * time.Second):
			fmt.Printf("block 2秒, %s\n", time.Now())

		}
	}
}
