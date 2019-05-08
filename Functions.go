package main

import (
				"fmt"
				"errors"
				"math"
			)


//if else conditions in go
func main() {
	sum := sum(2,3)
	fmt.Println(sum) //5
	result,err := sqrt(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result) //2.23606797749979
}

func sum(x int,y int) int {
	return x + y
}

func sqrt(x float64) (float64,error) {
	if x < 0 {
		return 0, errors.New("undefined")
	}
	return math.Sqrt(x),nil
}