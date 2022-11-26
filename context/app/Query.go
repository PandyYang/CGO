package app

import "context"

func Query(ctx context.Context, str <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for n := range str {
			out <- n
		}
		close(out)
	}()
	return out
}
