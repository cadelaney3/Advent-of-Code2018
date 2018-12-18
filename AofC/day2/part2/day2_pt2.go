package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countLetters(boxID string) (bool, bool) {
	var containsDubs bool
	var containsTrips bool
	for i := 0; i < len(boxID); i++ {
		if strings.Count(boxID, string(boxID[i])) >= 3 {
			containsTrips = true
		}
		if strings.Count(boxID, string(boxID[i])) == 2 {
			containsDubs = true
		}
	}
	return containsDubs, containsTrips
}

func findCommonLetters(ids []string) []string {
	var common []string
	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(ids); j++ {
			if i != j {
				diff := 0
				common = common[:0]
				for k := 0; k < len(ids[j]); k++ {
					if string(ids[i][k]) != string(ids[j][k]) {
						diff++
					} else {
						common = append(common, string(ids[i][k]))
					}
					if diff > 1 {
						break
					}
				}
				if diff == 1 {
					return common
				}
			}
		}
	}
	return common
}

func main() {
	file, err := ioutil.ReadFile("day2.txt")
	if err != nil {
		panic(err)
	}

	// make file into '\n' separated string
	str := string(file)

	// convert string into slice, splitting at '\n'
	lines := strings.Split(str, "\n")

	common := findCommonLetters(lines)
	fmt.Println(common)
	/*
		numDubs := 0
		numTrips := 0

		for i := 0; i < len(lines); i++ {
			dub, trip := countLetters(lines[i])
			if dub == true {
				numDubs++
			}
			if trip == true {
				numTrips++
			}
		}

		checkSum := numDubs * numTrips

		fmt.Println(checkSum)
	*/

	//testStr := "abcdef"
	//testStr2 := "ajcdef"
}
