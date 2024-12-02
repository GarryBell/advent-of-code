package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func checkRecord(recordArr []string) bool {
	first, _ := strconv.Atoi(recordArr[0])
	second, _ := strconv.Atoi(recordArr[1])
	isAscending := first < second
	for i := 0; i < (len(recordArr) - 2); i++ {
		current, _ := strconv.Atoi(recordArr[i])
		next, _ := strconv.Atoi(recordArr[i+1])
		diff := next - current
		if !isAscending {
			diff = -diff
		}
		if 1 > diff || diff > 3 {
			return false
		}

	}
	return true
}

func main() {
	filePath := "input.txt"
	readFile, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	count := 0
	countNormal := 0
	for fileScanner.Scan() {
		var recordArr = strings.Split(fileScanner.Text(), " ")
		var current = false

		if checkRecord(recordArr) {
			current = true
			countNormal += 1
		}

		for i := 0; i < len(recordArr); i++ {
			var arr = slices.Clone(recordArr)
			RemoveIndex(arr, i)
			if checkRecord(arr) {
				current = true
			}
		}
		if current {
			count += 1
		}
	}
	fmt.Println(countNormal)
	fmt.Println(count)

}
