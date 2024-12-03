package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func generate(mul string) int {
	body := mul[4 : len(mul)-1]
	first, _ := strconv.Atoi(strings.Split(body, ",")[0])
	second, _ := strconv.Atoi(strings.Split(body, ",")[1])
	return first * second
}

func main() {
	filePath := "input.txt"
	readFile, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	count := 0
	count2 := 0
	var doFlag = true

	for fileScanner.Scan() {
		var line = fileScanner.Text()
		pattern := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
		splits := pattern.FindAllString(line, -1)
		for i := 0; i < len(splits); i++ {
			if splits[i] == "do()" {
				doFlag = true
				continue
			}
			if splits[i] == "don't()" {
				doFlag = false
				continue
			}
			res := generate(splits[i])
			if doFlag {
				count2 += res
			}
			count += res
		}

	}
	fmt.Println(count)
	fmt.Println(count2)
}
