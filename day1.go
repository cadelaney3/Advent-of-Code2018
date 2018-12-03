package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("day1.txt")
	if err != nil {
		panic(err)
	}

	// make file into '\n' separated string
	str := string(file)

	// convert string into slice, splitting at '\n'
	lines := strings.Split(str, "\n")

	// len(lines)-1 because last index is blank
	total := 0
	for i := 0; i < len(lines)-1; i++ {
		if lines[i][:1] == "+" {
			value, err := strconv.Atoi(lines[i][1:])
			if err != nil {
				panic(err)
			}
			total += value
		}
		if lines[i][:1] == "-" {
			value, err := strconv.Atoi(lines[i][1:])
			if err != nil {
				panic(err)
			}
			total -= value
		}
	}
	fmt.Println("total = ", total)
}
