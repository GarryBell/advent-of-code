package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkNext(arr [10][]string, i int, j int, previousChar string, directionX int, directionY int) bool {
	if i < 0 || j < 0 || i >= 10 || j >= 10 {
		return false
	}
	fmt.Println("Checking", j, i, directionX, directionY)
	fmt.Println("which is", string(arr[j][0][i]), "prev: ", previousChar)

	if string(arr[j][0][i]) == "S" && previousChar == "A" {
		fmt.Println("Actually finished!")
		return true
	} else if previousChar == "X" && string(arr[j][0][i]) == "M" {
		return checkNext(arr, i+directionX, j+directionY, "M", directionX, directionY)
	} else if previousChar == "M" && string(arr[j][0][i]) == "A" {
		return checkNext(arr, i+directionX, j+directionY, "A", directionX, directionY)
	}
	return false
}

func startCheck(arr [10][]string, x int, y int) bool {
	fmt.Println("starting at x: ", x, "and  y: ", y)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if x+i < 0 || y+j < 0 || x+i > 10 || y+j > 10 {
				continue
			}
			if string(arr[y+j][0][x+i]) == "M" {
				res := checkNext(arr, x+i, y+j, "X", i, j)
				if res {
					return true
				}
			}
		}
	}
	return false

}

func main() {
	filePath := "input_small.txt"
	readFile, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	arr := [10][]string{}
	y := 0

	for fileScanner.Scan() {
		var line = strings.Split(fileScanner.Text(), "   ")
		fmt.Println(line)
		arr[y] = line
		y++
	}
	count := 0
	fmt.Println(arr)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(arr[j][0][i]) == "X" {
				fmt.Println("Caught 1")
				res := startCheck(arr, i, j)
				if res {
					count += 1
				}
			}
		}
		fmt.Println("count: ", count)
	}
	fmt.Println(count)

}
