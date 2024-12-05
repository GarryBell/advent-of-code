package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	first  int
	second int
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

func checkStartOfLine(current int, start []int, rules []Rule) bool {

	for _, rule := range rules {
		if rule.first == current {
			if find(rule.second, start) != -1 {
				return false
			}
		}
	}
	return true
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

	flag := true

	for index, current := range lineInts { // Iterating twice? good thing I'm not golfing

		start := lineInts[:index]
		if !checkStartOfLine(current, start, rules) {
			flag = false
		}
	}
	if flag {
		return lineInts[len(lineInts)/2]
	}

	return 0
}

func main() {
	filePath := "input.txt"
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
