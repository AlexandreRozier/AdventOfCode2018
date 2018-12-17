package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	twiceCount := 0      // Custom map to implement a list
	threeTimesCount := 0 // Custom map to implement a list

	file, err := os.Open("data.csv")
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		frequencies := make(map[string]int) // Custom map to implement a list

		// Build frequencies table for each character of the line
		for _, letter := range strings.Split(scanner.Text(), "") {
			frequencies[letter] ++
		}
		burned_2s := false
		burned_3s := false
		// Update counters based on each individual frequency
		for _, value := range frequencies {
			if value == 2 && !burned_2s{
				twiceCount++
				burned_2s = true
			}
			if value == 3 && !burned_3s{
				threeTimesCount++
				burned_3s = true
			}
		}
	}
	file.Close()
	fmt.Printf("Checksum is: %d \n", twiceCount*threeTimesCount)
}
