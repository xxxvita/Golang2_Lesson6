package main

// go run -race task3/main.go

import "fmt"

var generalResource int

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			generalResource++
		}()
	}

	fmt.Printf("После 1000 попыток увеличения значения: %d\n", generalResource)
}
