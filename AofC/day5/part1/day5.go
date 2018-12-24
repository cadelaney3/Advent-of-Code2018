package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func ASCII(r rune) rune {
	switch {
	case 97 <= r && r <= 122:
		return r - 32
	case 65 <= r && r <= 90:
		return r + 32
	default:
		return r
	}
}

func elimPairs(s string) string {
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			if math.Abs(float64(ASCII(rune(s[i]))-ASCII(rune(s[i+1])))) == float64(32) {
				s = strings.Replace(s, s[i:i+2], "", -1)
				i = -1
			}
		}
	}
	return s
}

func removePolymer(r rune, s string) string {
	s = strings.Replace(s, strings.ToLower(string(r)), "", -1)
	s = strings.Replace(s, strings.ToUpper(string(r)), "", -1)
	return s
}

func improvePolymer(s string, units string) string {
	best := s
	for _, v := range units {
		temp := elimPairs(removePolymer(v, s))
		if len(temp) < len(best) {
			best = temp
		}
	}
	return best
}

func main() {
	file, err := ioutil.ReadFile("day5.txt")
	if err != nil {
		panic(err)
	}

	str := string(file)
	lines := strings.Split(str, "\n")
	line := lines[0]

	fmt.Println("length of initial polymer: ", len(line))

	//testLine := "BaAbcdDefgHhijJ"
	//testLine2 := elimPairs(testLine)

	line = elimPairs(line)

	fmt.Println("part 1 answer: ", len(line))

	units := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("part 2 answer: ", len(improvePolymer(line, units)))
}
