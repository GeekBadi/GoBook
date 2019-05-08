package main

import (
			"fmt"
			"time"
			)

//go channels
func main() {
	c1 := make(chan string);	
	c2 := make(chan string);	
	go func() {
		for {
			c1 <- "every 500seconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "every 2 seconds";
			time.Sleep(time.Second * 2)
		}
		
	}()
	
	for {
		select {
			case msg1 := <- c1: fmt.Println(msg1)
			case msg2 := <- c2: fmt.Println(msg2)
		}
	}
}
