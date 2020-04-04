package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int) //Creating a channel

	a := []int{8, 7, 6, 5, 3, 0, 9, -1} // Array of integer

	go printCount(c) // Calling the function in goroutine

	for _, v := range a {
		c <- v // Pass the integer value into channel
	}

	time.Sleep(time.Millisecond * 1) // main pauses before ending
	fmt.Println("End of main")

}
