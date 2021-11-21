package main

/*
go run task1/main.go 2>~/work/go2/Golang2_Lesson6/task1/trace.txt &&
go tool trace task1/trace.txt

*/

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var (
	generalResource int
	countWorkers    int        = 10
	mutex           sync.Mutex = sync.Mutex{}
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	wg := sync.WaitGroup{}
	wg.Add(countWorkers)
	for i := 0; i < countWorkers/2; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()

			generalResource++
		}()
	}
	// Чтобы получше посмотреть трейс-лог
	for i := 0; i < countWorkers/2; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()

			generalResource++
		}()
	}

	wg.Wait()
	fmt.Printf("Все потоки завершены. Финальное значение общего ресурса равно %d\n", generalResource)
}
