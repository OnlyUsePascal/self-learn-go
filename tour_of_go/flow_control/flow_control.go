package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func while_loop() {
	// no semicolon
	sum := 0
	for sum < 10 {
		// for ; sum < 10; { // or this
		fmt.Println(sum)
		sum += 1
	}

	// forever loop
	for {
		fmt.Println(sum)
		sum += 1
		if sum >= 10 {
			break
		}
	}
}

func if_short(pow float64, lim float64) {
	if res := math.Pow(pow, 2); res > lim {
		fmt.Println("exceeded!")
	} else {
		fmt.Println("u r safe")
	}
}

func switch_case() {
	// swtich case value no need to be constant
	// can use other types than number
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch with no condition
	// can replace long if-else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func defer_ (useStack bool) {
	// its like finally in java
	if (!useStack){
		fmt.Println("> Defer basic")
		defer func () {
			fmt.Println("Print last regardless anything")
		}()
		
		fmt.Println("Hello world!")
	} else {
		fmt.Println("> Defer stack ")
		for i := 0; i < 5; i++ {
			defer fmt.Println(i)
		}
	}
}

func main() {
	while_loop()

	if_short(2, 10)
	if_short(4, 10)

	switch_case()
	
	defer_(false)
}
