// Напишите программу, которая делает следующее: одна горутина по порядку
// отсылает числа от 1 до 100 в канал, вторая горутина их принимает в
// правильном порядке и печатает на экран (в консоль).

package main

import (
	"fmt"
)

func sender(c chan<- int, done chan<- struct{}) {
	defer close(c)
	defer close(done)
	for i := 1; i <= 100; i++ {
		c <- i
	}
}

func receiver(c <-chan int) {
	for val := range c {
		fmt.Print(val, " ")
	}
}

func main() {
	intChan := make(chan int)
	done := make(chan struct{})

	go sender(intChan, done)
	go receiver(intChan)

	<-done
}
