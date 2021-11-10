package main

import (
	p "dojo-golang/parallel"
	"sync"
)

func main() {
}

func testMain() int {
	return 2

}

// fazer um loop com 100 chamadas para hello world e verificar se todos os numeros serão chamados
// paralelizar pela quantidade de processadores da máquina

func callHundred(ch chan int) {

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go p.CustomReturn(i, ch, &wg)
	}
	wg.Wait()
	close(ch)

}
