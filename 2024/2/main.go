package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkRecord(record string) bool {
	var recordArr = strings.Split(record, " ")
	first, _ := strconv.Atoi(recordArr[0])
	second, _ := strconv.Atoi(recordArr[1])
	isAscending := first < second
	for i := 0; i < (len(recordArr) - 1); i++ {
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
	for fileScanner.Scan() {

		if checkRecord(fileScanner.Text()) {
			count += 1
		}
	}
	fmt.Println(count)

}