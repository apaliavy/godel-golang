package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 0

	var mu sync.RWMutex

	for i := 0; i < 20; i++ {
		go func() {
			mu.Lock()
			x++
			mu.Unlock()
		}()

		go func() {
			mu.RLock()
			y := x
			mu.RUnlock()

			if y%2 == 0 {
				fmt.Println(y)
			}
		}()
	}
}
