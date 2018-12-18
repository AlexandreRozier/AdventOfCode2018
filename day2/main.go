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

	// Part 1
	twiceCount := 0      // Custom map to implement a list
	threeTimesCount := 0 // Custom map to implement a list
	var nbOfIDs = 0
	file, err := os.Open("data.csv")
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nbOfIDs++
		frequencies := make(map[string]int) // Custom map to implement a list

		// Build frequencies table for each character of the line
		for _, letter := range strings.Split(scanner.Text(), "") {
			frequencies[letter] ++
		}
		burned_2s := false
		burned_3s := false
		// Update counters based on each individual frequency
		for _, value := range frequencies {
			if value == 2 && !burned_2s {
				twiceCount++
				burned_2s = true
			}
			if value == 3 && !burned_3s {
				threeTimesCount++
				burned_3s = true
			}
		}
	}
	fmt.Printf("Checksum is: %d \n", twiceCount*threeTimesCount)
	file.Close()

	// Part 2
	file, err = os.Open("data.csv")
	check(err)
	scanner = bufio.NewScanner(file)
	listOfIDs := make([]string, nbOfIDs)
	i := 0
	for scanner.Scan() {

		listOfIDs[i] = scanner.Text()
		i++
	}
	for i, elt := range listOfIDs {
		for j := i; j < len(listOfIDs); j++ {
			elt2 := listOfIDs[j]
			// Get number of different letters
			differentLettersCount := 0
			for k, letter := range elt {
				if uint8(letter) != elt2[uint8(k)] {
					differentLettersCount++
				}

			}
			if differentLettersCount == 1 {
				fmt.Printf(elt + "\n")
				fmt.Printf(elt2 + "\n")
			}

		}
	}

}
