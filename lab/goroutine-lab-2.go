package lab

import (
	"fmt"
	"runtime"
	"sync"
)

func GoRoutineLab2() {
	x := 10
	y := 10
	c := make(chan int)

	var wg sync.WaitGroup

	go func() {
		for i := 0; i < x; i++ {
			wg.Add(1)
			go func(m int) {
				for j := 0; j < y; j++ {
					c <- (10 * m) + j
				}
				wg.Done()
			}(i)
			fmt.Println("ROUTINES", runtime.NumGoroutine())

		}
		wg.Wait()
		close(c)
	}()

	for number := range c {
		fmt.Println(number)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}
