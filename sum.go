package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args
	numA, err := strconv.Atoi(input[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	numB, err := strconv.Atoi(input[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	result := numA + numB
	fmt.Printf("%d + %d = %d\n", numA, numB, result)

}
