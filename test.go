package main

import (
	pylist "Pyutils/pyutils/Pylist"
	"Pyutils/pyutils/genericList"
	"fmt"
)

func main() {
	fmt.Println("Your test is here: ")
	var x pylist.List
	// = PyList.List{}
	// fmt.Println(x)
	for _, value := range []int{5, 6, 7, 8, 9} {
		x.Append(value)
	}
	fmt.Println(x)
	pylist.TestList()
	fmt.Println("End of test list Interface: ")
	fmt.Println("Start of test list Integer: ")
	genericList.TestInt()
	fmt.Println("Start of test list Float: ")
	genericList.TestFloat()
	fmt.Println("Start of test list String: ")
	genericList.TestString()
	fmt.Println("Start of test Dict: ")
	pylist.TestDict()
	// blockchain.PrettyPrint(balance)
}
