package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var knownFrequencies map[int]interface{}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i := 0
	total := 0
	knownFrequencies = make(map[int]interface{})
	for i < 1000 {
		file.Seek(0, 0)
		total = scanFile(file, total)
		i++
	}

}

func scanFile(file io.Reader, total int) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		total += i

		if _, ok := knownFrequencies[total]; ok {
			fmt.Println(total)
			os.Exit(1)
		}

		knownFrequencies[total] = total

	}

	return total
}
