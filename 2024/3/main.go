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
	// fmt.Println(mul)
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

	for fileScanner.Scan() {
		fmt.Println("----------------------")
		var line = fileScanner.Text()
		pattern := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		splits := pattern.FindAllString(line, -1)
		for i := 0; i < len(splits); i++ {
			res := generate(splits[i])
			fmt.Print(res, ' ')
			count += res
		}
		fmt.Println(count)
	}

	fmt.Println(count)

}
