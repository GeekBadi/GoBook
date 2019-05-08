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
			time.Sleep(time.Millisecond * 2)
		}
		
	}()
	
	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
