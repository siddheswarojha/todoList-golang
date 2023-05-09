package main

import "strings"

func input(inputData string, in []string) {

	var data = strings.Trim(inputData, " ")
	in = append(in, data)
	dashboard(in)

}
