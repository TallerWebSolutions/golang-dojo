package main

import (
	"fmt"
	"testing"
)

func TestTestMain(t *testing.T) {
	r := testMain()
	expect := 2
	if r != expect {
		t.Error("Expected", expect, "got", r)
	}

}

func TestHundred(t *testing.T) {
	ch := make(chan int, 100)
	callHundred(ch)
	arrayResponse := make([]int, 100)
	for k := range ch {
		res := <-ch
		arrayResponse[k] = res
	}

	for i := 0; i < 100; i++ {
		isInArray := false
		for _, v := range arrayResponse {
			//fmt.Println("dentro do array Response: ", v)
			if v == i {
				fmt.Println("entrou no if")
				isInArray = true
				return
			}
		}

		if !isInArray {
			t.Error("Expected", i, "got", arrayResponse)
		}
	}
}
