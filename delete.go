package main

import "strconv"

func delete(index string, in []string) {

	index1, _ := strconv.ParseInt(index, 0, 255)

	in = append(in[:index1], in[index1+1:]...)
	dashboard(in)
}
