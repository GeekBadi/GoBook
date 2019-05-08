package main

import ("fmt")

//go channels 
func main() {
	c := make(chan string);
	go write(c)
	msg := <- c
	fmt.Println(msg)
}

func write(c chan string) {
	c <- "hello world"
}
