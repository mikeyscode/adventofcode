package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		total += i
	}

	fmt.Println(total)
}
