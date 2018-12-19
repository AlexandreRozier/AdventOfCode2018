package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func react(dat string) string {
	shouldReact := true

	for shouldReact {
		for i := 0; i < len(dat)-1; i++ {
			elt1 := string(dat[i])
			elt2 := string(dat[i+1])

			if (strings.ToUpper(elt1) == elt2 || strings.ToUpper(elt2) == elt1) && elt1 != elt2 {
				// react
				dat = dat[:i] + dat[i+2:]
				break
			}
			if i == len(dat)-2 {
				shouldReact = false
			}

		}
		// No more possible reactions

	}
	return dat
}

func main() {
	data, _ := ioutil.ReadFile("day5/data")
	// Part 1
	println(len(react(string(data))))

	// Part 2
	var lenghts [26]int

	for i := 0; i < 26; i++ {
		letter := string('a' + i)
		trimmedData := strings.Replace(string(data), letter, "", -1)
		trimmedData = strings.Replace(trimmedData, strings.ToUpper(letter), "", -1)
		lenghts[i] = len(react(trimmedData))
		println("Processed "+letter)
	}

	m := len(data)
	for _, e := range lenghts {
		if e < m {
			m = e
		}
	}
	println(m)

}
