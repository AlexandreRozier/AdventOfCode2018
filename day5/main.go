package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	pouet, _ := ioutil.ReadFile("day5/data")
	dat := string(pouet)
	shouldReact := true
	for shouldReact {
		for i := 0; i< len(dat)-1;i++{
			elt1 := string(dat[i])
			elt2 := string(dat[i+1])

			if (strings.ToUpper(elt1)==elt2 || strings.ToUpper(elt2)==elt1) && elt1!=elt2  {
				// react
				dat = dat[:i]+dat[i+2:]
				break
			}
			if i == len(dat)-2 {
				shouldReact = false
			}

		}
		// No more possible reactions



	}
	println(len(dat))

}
