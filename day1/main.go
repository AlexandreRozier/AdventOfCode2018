package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func check(err error){
	if err !=nil { panic(err)}
}

func main() {

	result := 0
	frequencies := make(map[int]int) // Custom map to implement a list
	continueSearch := true

	for i:=0;i<1000 && continueSearch;i++ {

		file, err := os.Open("data")
		check(err)

		scanner := bufio.NewScanner(file)

		for scanner.Scan(){

			stringy := scanner.Text()
			toadd, err := strconv.Atoi(stringy)

			check(err)

			// Check if frequency has already been reached
			result += toadd // new freq
			_, exists := frequencies[result]

			// If frequency has already been seen
			if exists {
				fmt.Printf("First frequency already reached: %d \n", result)
				continueSearch = false
				break
			} else {
				frequencies[result] = 1
			}

		}
		file.Close()
		if i==0 {fmt.Printf("Total sum: %d \n", result)}
	}

}
