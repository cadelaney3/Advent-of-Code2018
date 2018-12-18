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

func main() {
	file, err := ioutil.ReadFile("day2.txt")
	if err != nil {
		panic(err)
	}

	// make file into '\n' separated string
	str := string(file)

	// convert string into slice, splitting at '\n'
	lines := strings.Split(str, "\n")

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
}
