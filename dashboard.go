package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dashboard(in []string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1.input, 2.read, 3.delete")
	readInput, _ := reader.ReadString('\n')

	readInput = strings.TrimSpace(strings.ReplaceAll(readInput, "\n", ""))

	fmt.Println(readInput)
	if readInput == "1" {
		fmt.Println("Enter a data: ")
		inputData, error := reader.ReadString('\n')
		if error != nil {
			fmt.Println("error")
		} else {
			input(inputData, in)
		}

	} else if readInput == "2" {
		read(in)
	} else {
		fmt.Println("delete called")
		index, _ := reader.ReadString('\n')
		delete(index, in)
	}
}
