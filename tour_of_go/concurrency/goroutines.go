package main

import (
	"fmt"
	"slices"
	"time"
)

func channels_range() {
	// Closing is only necessary when the receiver must be told
	// there are no more values coming, such as to terminate a range loop.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

func channels_buffered() {
	c := make(chan int, 2)
	c2 := make(chan int)

	c <- 1
	c <- 2
	c2 <- 1
	c2 <- 2

	// > would error
	// fetch too much
	<-c
	<-c
	<-c

	// not fetch enough or too much
	<-c2

	// > would not error
	// can leave remaining
	c <- 1
	c <- 2
	<-c

	// must fetch enough
	c2 <- 1
	c2 <- 2
	<-c2
	<-c2
}



func channels_basic() {
	get_sum := func(prefix string, arr []int, channel chan int) {
		_sum := 0
		for _, item := range arr {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("%s,%v \n", prefix, item)

			_sum += item
		}

		channel <- _sum
	}

	arr := []int{7, 2, 8, -9, 4, 1, 1}
	sum := 0
	channel := make(chan int)
	chunkIdx := 0
	chunkSz := 3

	for chunkArr := range slices.Chunk(arr, chunkSz) {
		prefix := fmt.Sprintf("Chunk %v", chunkIdx)
		chunkIdx++

		go get_sum(prefix, chunkArr, channel)
	}

	//cannot use range channel loop
	chunkCnt := len(arr) / chunkSz
	if len(arr)%chunkSz != 0 {
		chunkCnt++
	}

	for i := 0; i < chunkCnt; i++ {
		fmt.Printf("> Chunk %v\n", i)
		sum += <-channel
	}

	fmt.Println("Sum:", sum)
}

func goroutine_basic() {
	make_loop := func(prefix string) {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("%s,%s\n", prefix, i)
		}
	}

	go make_loop("with goroutine")
	make_loop("no goroutine")
}

func main() {
	// goroutine_basic()

	// channels_basic()

	// channels_buffered()

	// channels_range()
	
	Mutex_basic()
}
