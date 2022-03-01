package lab

import (
	"fmt"
	"strings"
	"sync"
)

func GoRoutineLab() {

	const size = 5

	wg := new(sync.WaitGroup)
	ch := make(chan int, size)

	for i := 1; i <= size; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			ch <- x * x
		}(i)
	}
	//go func() {
	wg.Wait()
	close(ch)
	//}()

	for s := range ch {
		fmt.Println(s)
	}
	fmt.Println(strings.Repeat("#", 20))

}
