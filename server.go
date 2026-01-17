package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Println("goroutine 1 = ", i)
		}
	}()

	wg.Add(1)
	go func() {
		duration := time.Second
		time.Sleep(duration * 3)
		defer wg.Done()

		for i := 11; i <= 20; i++ {
			fmt.Println("goroutine 2 = ", i)
		}
	}()

	wg.Wait()

	port := "3000"
	msg := "Server starting on port "
	fmt.Println(msg + port)
	http.ListenAndServe(":3000", nil)
}
