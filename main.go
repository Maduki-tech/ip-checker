package main

import (
	"strings"
)

func sliceInput(input string) []string{
	sliced := strings.Split(input, ".")
	return sliced
}

func main(){
	sliceInput("192.168.0.0")
}
