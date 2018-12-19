package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type conflictsAndIDs struct {
	conflictsNB int
	idList      []int
}

func main() {

	c := [1500][1500]conflictsAndIDs{}

	file, _ := os.Open("data.csv")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//#1285 @ 803,711: 16x24

		str := strings.Split(scanner.Text(), " ")
		id, _ := strconv.Atoi(strings.Trim(str[0], "#"))
		pos := strings.Split(str[2], ",")
		pos[1] = strings.Replace(pos[1], ":", "", -1)
		dim := strings.Split(str[3], "x")
		j, _ := strconv.Atoi(pos[0])
		i, _ := strconv.Atoi(pos[1])
		nDim, _ := strconv.Atoi(dim[0])
		mDim, _ := strconv.Atoi(dim[1])

		for m := i; m < i+mDim; m++ {
			for n := j; n < j+nDim; n++ {
				c[m][n] = conflictsAndIDs{c[m][n].conflictsNB + 1, append(c[m][n].idList, id)}
			}
		}

	}

	conflictsNB := 0
	// Get conflicts
	for _, row := range c {
		for _, elt := range row {
			if elt.conflictsNB > 1 {
				conflictsNB++
			}
		}
	}
	// Part 1
	fmt.Println(conflictsNB)
	file.Close()


	// Part 2
	file, _ = os.Open("data.csv")
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		id, _ := strconv.Atoi(strings.Trim(str[0], "#"))
		pos := strings.Split(str[2], ",")
		pos[1] = strings.Replace(pos[1], ":", "", -1)
		dim := strings.Split(str[3], "x")
		j, _ := strconv.Atoi(pos[0])
		i, _ := strconv.Atoi(pos[1])
		nDim, _ := strconv.Atoi(dim[0])
		mDim, _ := strconv.Atoi(dim[1])

		conflicts := false
		for m := i; m < i+mDim; m++ {
			for n := j; n < j+nDim; n++ {
				if c[m][n].conflictsNB > 1 {
					conflicts = true
				}
			}
		}
		if !conflicts {
			println(id)
		}

	}


}
