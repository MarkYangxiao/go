package main

import (
	"context"
	"fmt"
)

const (
	UserName = "yangxiao"
	Sex      = "男"
)

func main() {
	parentCtx := context.Background() // 创建一个空的父context
	// 输出parent
	Process(parentCtx)
	ctx := context.WithValue(parentCtx, UserName, Sex) //
	Process(ctx)
}

func Process(ctx context.Context) {
	sex, ok := ctx.Value(UserName).(string)
	if !ok {
		fmt.Printf("process over, can not find value of key:%s in context \n", UserName)
	} else {
		fmt.Printf("process over, find value:%s of key:%s in context\n", sex, UserName)
	}
}
