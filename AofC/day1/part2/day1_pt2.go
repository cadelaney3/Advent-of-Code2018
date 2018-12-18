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

	frequencies := []int{0}
	noRepeat := true

	// len(lines)-1 because last index is blank
	total := 0
	for noRepeat == true {
		for i := 0; i < len(lines)-1; i++ {
			//var value int
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
			for j := 0; j < len(frequencies); j++ {
				if frequencies[j] == total {
					fmt.Println("repeat = ", total)
					noRepeat = false
					break
				}
			}

			frequencies = append(frequencies, total)

			if noRepeat == false {
				break
			}
		}
	}
}
