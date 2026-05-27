package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	go func() {
		defer fmt.Print("done")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println(i, len(c), cap(c))
		}

		time.Sleep(1 * time.Second)
	}()

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println(num)
	}
	fmt.Println("123")
}
