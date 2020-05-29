package main

import (
	"context"
	"fmt"
	"time"
)

type a struct {
	Info string
}

func main() {
	asyncCallCMD(context.Background())
}

func asyncCallCMD(ctx context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	ctx = context.WithValue(ctx, "id", &a{Info: "1"})
	defer cancel()
	cancel()
	go callCMD(ctx)
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Value("id").(*a).Info)
		fmt.Println("call successfully!!!")
		return
	case <-time.After(time.Duration(time.Second * 5)):
		fmt.Println("timeout!!!")
		return
	}
}

func callCMD(ctx context.Context) {
	fmt.Println("h:" + ctx.Value("id").(*a).Info)
	ctx.Value("id").(*a).Info = "2"
}
