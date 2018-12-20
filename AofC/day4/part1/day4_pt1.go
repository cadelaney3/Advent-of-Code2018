package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	time          time.Time
	line          string
	guard         *Guard
	minutesAsleep map[int]int
}

type timeSlice []*Record

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].time.Before(p[j].time)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Guard struct {
	id              int
	totalTimeAsleep int
	minuteAsleep    map[int]*Minute
}

type Minute struct {
	minute int
	count  int
}

func max(numbers interface{}) (int, int) {
	var maxNumber int
	var maxKey int
	switch v := numbers.(type) {
	case map[int]*Minute:
		for key, val := range v {
			maxNumber = val.count
			maxKey = key
			break
		}
		for key, val := range v {
			if val.count > maxNumber {
				maxNumber = val.count
				maxKey = key
			}
		}
	case map[int]*Guard:
		for key, val := range v {
			maxNumber = val.totalTimeAsleep
			maxKey = key
			break
		}
		for key, val := range v {
			if val.totalTimeAsleep > maxNumber {
				maxNumber = val.totalTimeAsleep
				maxKey = key
			}
		}
	default:
		maxNumber = 0
		maxKey = 0
	}
	return maxKey, maxNumber
}

func main() {

	file, err := ioutil.ReadFile("day4.txt")
	if err != nil {
		panic(err)
	}

	str := string(file)
	lines := strings.Split(str, "\n")

	records := map[time.Time]*Record{}

	const form = "2006-01-02 15:04:05"

	for _, v := range lines[:len(lines)-1] {
		t, _ := time.Parse(form, v[1:17]+":00")
		newRec := Record{time: t, line: v}
		records[t] = &newRec
	}

	recSlice := make(timeSlice, 0, len(records))
	for _, val := range records {
		recSlice = append(recSlice, val)
	}
	sort.Sort(recSlice)

	var currGuard int
	var fallAsleepTime time.Time
	var awakeTime time.Time
	guards := map[int]*Guard{}

	for _, value := range recSlice {

		// regex to get only the nums from each line in day3.txt
		re := regexp.MustCompile("[0-9]+")

		// array of nums extracted from regex
		// still in string form, will need to convert to int
		nums := re.FindAllString(value.line, -1)

		if len(nums) == 6 {
			currGuard, _ = strconv.Atoi(nums[5])
		}

		if _, ok := guards[currGuard]; !ok {
			newGuard := Guard{id: currGuard, minuteAsleep: make(map[int]*Minute)}
			guards[currGuard] = &newGuard
		}

		if string(value.line[19]) == "f" {
			fallAsleepTime = value.time
		}
		if string(value.line[19]) == "w" {

			fallAsleepM := int(fallAsleepTime.Minute())
			awakeTime = value.time
			diff := awakeTime.Sub(fallAsleepTime)
			timeAsleep := int(diff.Minutes())

			lastMinAsleep := fallAsleepM + int(diff.Minutes())

			guards[currGuard].totalTimeAsleep += timeAsleep

			for i := fallAsleepM; i <= lastMinAsleep; i++ {
				if _, ok := guards[currGuard].minuteAsleep[i]; ok {
					guards[currGuard].minuteAsleep[i].count++
				} else {
					newMin := Minute{minute: i, count: 1}
					guards[currGuard].minuteAsleep[i] = &newMin
				}
			}

		}
	}

	guardAsleepMost, mostMinutesAsleep := max(guards)
	fmt.Println("guard asleep the most and time asleep: ", guardAsleepMost, ": ", mostMinutesAsleep)

	minuteAsleepMost, countAsleep := max(guards[guardAsleepMost].minuteAsleep)

	fmt.Println("Minute asleep most and count: ", minuteAsleepMost, ": ", countAsleep)
	fmt.Println("guardID * minAsleepMost = ", guardAsleepMost*minuteAsleepMost)
	fmt.Println("correct aswer was 76357 for minute 29,",
		"but multiple minutes have a count of 14 so answer varies when run")
}
