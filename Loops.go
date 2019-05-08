package main

import ("fmt")

//if else conditions in go
func main() {
	x := 16
	if x > 18 {
		fmt.Println("You are major")
	} else if x > 10 {
		fmt.Println("You are youth")
	} else {
		fmt.Println("You are a child")
	}
}