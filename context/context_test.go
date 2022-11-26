package context

import (
	"CGO/context/app"
	"context"
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	str := "123"
	ctx, _ := context.WithCancel(context.Background())
	out := make(chan string)
	out <- str
	res := app.Query(ctx, out)
	fmt.Println(<-res)
}

// 处理多组数据 入管道
func gen(str ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, n := range str {
			out <- n
		}
		close(out)
	}()
	return out
}
