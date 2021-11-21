package main

/*
go run task2/main.go 2>~/work/go2/Golang2_Lesson6/task2/trace.txt &&
go tool trace task2/trace.txt

*/

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, startTime time.Time, i int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					cancel()
					return
				default:
					if time.Since(startTime) >= time.Second {
						startTime = time.Now()
						//fmt.Printf("Goroutine nmumber %d %v\n", i, time.Since(startTime))

						runtime.Gosched()
					}
				}
			}
		}(ctx, time.Now(), i)
	}

	wg.Wait()
	fmt.Printf("Все потоки завершены.\n")
}
