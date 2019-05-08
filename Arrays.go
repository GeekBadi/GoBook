package main

import ("fmt")

func main() {
	var a[5] int; //initialization of array
	b := []int{1,2,3,4,5} //another way of intialization

	fmt.Println(a) //[0 0 0 0 0 ]
	a[3] = 5 //modifcation of array
	fmt.Println(a) //[0 0 0 5 0]
	fmt.Println(b) //[1,2,3,4,5]
	fmt.Println(b)  //[1 2 3 5]
}