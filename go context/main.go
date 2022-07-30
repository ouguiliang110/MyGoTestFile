package main

import (
	"context"
	"fmt"
	"time"
)

func g2(ctx context.Context) {
	fmt.Println(ctx.Value("g2"))
	fmt.Println("我是g2")
	fmt.Println(ctx.Value("begin"))
}

func g1(ctx context.Context) {
	fmt.Println(ctx.Value("begin"))
	fmt.Println("g1:你是猪")
	go g2(context.WithValue(ctx, "g2", "g2:你也是猪")) //接着上文值
}

func main() {
	ctx := context.WithValue(context.Background(), "begin", "从台词看到一句话")
	go g1(ctx)
	time.Sleep(time.Second)
}
