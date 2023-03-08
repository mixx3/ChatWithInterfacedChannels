package main

import "fmt"

func process() int {
	return 666
}

func futures() {
	//futures
	future := make(chan int, 1)
	go func() { //async
		future <- process()
	}()
	result := <-future // await
	fmt.Println(result)
}

func scatter() chan int {
	c := make(chan int, 10)
	for i := 0; i < cap(c); i++ {
		go func() {
			val := process()
			c <- val
		}()
	}
	return c
}

func gather() int {
	c := scatter()
	var total int
	for i := 0; i < cap(c); i++ {
		res := <-c
		total += res
	}
	return total
}
