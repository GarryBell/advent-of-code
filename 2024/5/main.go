package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	first  int
	second int
}

func swapVars(a_index int, b_index int, arr []int) []int {
	newArr := slices.Clone(arr)
	swapVar := newArr[a_index]
	newArr[a_index] = newArr[b_index]
	newArr[b_index] = swapVar
	return newArr
}
func parseRules(rule string) Rule {
	rules := strings.Split(rule, "|")
	out := Rule{}
	out.first, _ = strconv.Atoi(rules[0])
	out.second, _ = strconv.Atoi(rules[1])
	return out
}

func find(what int, where []int) (idx int) {
	for i, v := range where {
		if v == what {
			return i
		}
	}
	return -1
}

func checkStartOfLine(current int, start []int, rules []Rule) int {

	for _, rule := range rules[:] {
		if rule.first == current {
			found := find(rule.second, start)
			if found != -1 {
				return found
			}
		}
	}
	return -1
}

func checkLine(line string, rules []Rule) int {
	lineArray := strings.Split(line, ",")
	var lineInts = []int{}

	for _, i := range lineArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		lineInts = append(lineInts, j)
	}

	flag := false

	for index, current := range lineInts[:] { // Iterating twice? good thing I'm not golfing

		start := lineInts[:index]
		res := checkStartOfLine(current, start, rules)
		if res != -1 {
			fmt.Println("swapping ", lineArray[index], " and ", lineArray[res])
			fmt.Println("before: ", lineInts)
			lineInts = slices.Clone(swapVars(index, res, lineInts))
			fmt.Println("after: ", lineInts)
			flag = true
			index = 0
		}
	}
	if flag {
		fmt.Println("Flagged", lineInts, lineInts[len(lineInts)/2])
		return lineInts[len(lineInts)/2]
	}

	return 0
}

func main() {
	filePath := "input_small.txt"
	readFile, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	rulesArray := []Rule{}
	for fileScanner.Scan() {
		var line = strings.Split(fileScanner.Text(), "   ")[0]

		if line == "" {
			break
		}
		rulesArray = append(rulesArray, parseRules(line))
	}
	count := 0
	pages := []string{}
	for fileScanner.Scan() {
		var line = strings.Split(fileScanner.Text(), "   ")[0]
		pages = append(pages, line)
	}

	for i := 0; i < len(pages); i++ {
		count += checkLine(pages[i], rulesArray)
	}
	fmt.Println(count)

}
