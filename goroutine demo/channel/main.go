package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	i int
}

var cter Counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go func(i int) {

			v := Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
		}(i)
	}
	time.Sleep(time.Second * 5)
}
