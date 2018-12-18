package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//"BoardPos is a struct for board position so that count value can be updated when used in a map"
type BoardPos struct {
	id    int
	count int
}

func main() {

	file, err := ioutil.ReadFile("day3.txt")
	if err != nil {
		panic(err)
	}

	str := string(file)

	lines := strings.Split(str, "\n")

	// make a map to store each position of fabric that shows up in a claim
	// use BoardPos struct as value to keep track of number of times a
	// fabric position appears
	claim := map[int]*BoardPos{}

	for i := 0; i < len(lines)-1; i++ {

		// regex to get only the nums from each line in day3.txt
		re := regexp.MustCompile("[0-9]+")

		// array of nums extracted from regex
		// still in string form, will need to convert to int
		nums := re.FindAllString(lines[i], -1)

		// nums[0] is fabric claim id, which we don't need
		// nums[1] is inches from left edge
		fromLeft, _ := strconv.Atoi(nums[1])
		// nums[2] is inches from top edge
		fromTop, _ := strconv.Atoi(nums[2])
		// nums[3] is width dimension of fabric claim
		width, _ := strconv.Atoi(nums[3])
		// nums[4] is height dimension of fabric claim
		height, _ := strconv.Atoi(nums[4])

		// calculate indexes of each fabric position in a claim assuming
		// fabric is 1000x1000 inches
		// use top left as (0,0) position, increasing left to right and top to bottom
		for j := fromTop; j < fromTop+height; j++ {
			for k := fromLeft + 1; k <= fromLeft+width; k++ {
				// calculate the fabric position
				fabricPos := j*1000 + k
				// if fabric position already appeared and is stored in claim map,
				// update the count
				// if not already appeared, add position to claim map
				if key, ok := claim[fabricPos]; ok {
					key.count++
				} else {
					newC := BoardPos{id: fabricPos, count: 1}
					claim[fabricPos] = &newC
				}
			}
		}
	}

	// total the number of positions with an appearance count >= 2
	total := 0
	for _, v := range claim {
		if v.count >= 2 {
			total++
		}
	}

	fmt.Println(total)

}
