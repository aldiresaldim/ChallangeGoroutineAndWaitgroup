package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(8)

	go func() {
		data := []interface{}{"bisa1", "bisa2", "bisa3"}
		for i := 0; i < 4; i++ {
			fmt.Println(data, i+1)
			time.Sleep(time.Millisecond * 500)
			wg.Done()
		}
	}()

	go func() {
		data := []interface{}{"coba1", "coba2", "coba3"}
		for i := 0; i < 4; i++ {
			fmt.Println(data, i+1)
			time.Sleep(time.Millisecond * 500)
			wg.Done()
		}
	}()

	wg.Wait()
}
