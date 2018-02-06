package main

import (
	"fmt"
	"time"
)

// Stricts that cha is sender
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "Ping"
	}
}

// Bi-directional channel
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "Pong"
	}
}

func printer(c chan string) {
	for {
		msg:= <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "str 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "str 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
				case msg1 := <- c1:
					fmt.Println("first print", 	msg1)
				case msg2 := <- c2:
					fmt.Println(msg2)

				// This will work enourmus times before any of first will happen if uncommented
				// default:
				// 	fmt.Println("nothing here")
			}
		}
	}()
	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	var input string
	fmt.Scanln(&input)
}