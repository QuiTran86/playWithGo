package main

import (
	"fmt"
	"reflect"
)

const (
	Pi               = 3.14
	avogardo float32 = 6.022e23
)

func main() {
	fmt.Println("What is the value of Pi? Pi is: ", Pi)
	fmt.Println("The type of Pi is: ", reflect.TypeOf(Pi))
	fmt.Println("What is the value of Avogardo? Avogardo is: ", avogardo)
	fmt.Println("The type of Avogardo is: ", reflect.TypeOf(avogardo))
}
