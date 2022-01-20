package main

import (
	"fmt"
	"sync"
)

var x = 1
var sum = 0

func increment(wg *sync.WaitGroup, sum chan int, i int) {
	var temp int
	for ; i < 1000; i = i + 5 {
		temp += i * i
	}
	sum <- temp
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	sum := make(chan int)
	sum1 := make(chan int)
	sum2 := make(chan int)
	sum3 := make(chan int)
	sum4 := make(chan int)

	w.Add(5)
	go increment(&w, sum, 1)
	go increment(&w, sum1, 2)
	go increment(&w, sum2, 3)
	go increment(&w, sum3, 4)
	go increment(&w, sum4, 5)

	total := <-sum + <-sum1 + <-sum2 + <-sum3 + <-sum4
	w.Wait()
	fmt.Println("The sum is : ", total)

}
