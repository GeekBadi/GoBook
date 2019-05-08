package main

import ("fmt")

func main() {
	vertices := make(map[string] int);
	vertices["one"] =1
	vertices["two"] = 2
	vertices["three"] = 3
	fmt.Println(vertices["one"]) //1
	fmt.Println(vertices) //map[one:1 two:2 three:3]
	delete(vertices,"three")
	fmt.Println(vertices) //map[one:1 two:2]
}