package parallel

import (
	"fmt"
	"sync"
)

func CustomReturn(x int, ch chan int, wg *sync.WaitGroup) {
	fmt.Println("CustomReturn:", x)
	ch <- x

	wg.Done()

}
