package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func calculateTotalDistance(list1, list2 []int) int {
	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		totalDistance += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return totalDistance
}

func simScore(list1, list2 []int) int {
	score := 0
	for i := 0; i < len(list1); i++ {
		test := func(s int) bool { return s == list1[i] }
		split := filter(list2, test)
		score += int(list1[i] * len(split))
	}
	return score
}

func main() {
	filePath := "input.txt"
	readFile, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var leftArray []int
	var rightArray []int

	for fileScanner.Scan() {
		var line = strings.Split(fileScanner.Text(), "   ")
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		leftArray = append(leftArray, l)
		rightArray = append(rightArray, r)
	}
	slices.Sort(leftArray)
	slices.Sort(rightArray)
	readFile.Close()
	var diff = calculateTotalDistance(leftArray, rightArray)
	fmt.Println(diff)
	fmt.Println(simScore(leftArray, rightArray))
}
