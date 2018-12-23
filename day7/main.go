package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var prerequisitesByLetter = make(map[string][]string)

func checkPrerequisites(testLetter string, eventList []string) bool {
	// Check if the prerequisites are met
	allPrerequisitesMet := true
	for _, prerequisite := range prerequisitesByLetter[testLetter] {
		prerequisiteMet := false
		for _, elt := range eventList {
			if elt == prerequisite {
				prerequisiteMet = true
			}
		}
		if !prerequisiteMet {
			allPrerequisitesMet = false
		}
	}
	return allPrerequisitesMet
}

func main() {
	//Step X must be finished before step L can begin.
	file, err := os.Open("day7/data.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var lettersList []string

	var eventList []string
	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), " ")
		pred := coordinates[1]
		succ := coordinates[7]

		exists := false
		for _, elt := range lettersList {
			if elt == pred {
				exists = true
			}
		}
		if !exists {
			lettersList = append(lettersList, pred)
		}
		exists = false
		for _, elt := range lettersList {
			if elt == succ {
				exists = true
			}
		}
		if !exists {
			lettersList = append(lettersList, succ)
		}
		sort.Strings(lettersList)

		prerequisitesByLetter[succ] = append(prerequisitesByLetter[succ], pred)
	}
	sort.Strings(lettersList)

	// Part 1
	var toProcessSortedList = make([]string, len(lettersList))
	copy(toProcessSortedList, lettersList)
	eventList = nil
	for len(eventList) < 26 {
		for i := 0; i < len(lettersList); i++ {
			candidateLetter := lettersList[i]
			prerequisitesMet := checkPrerequisites(candidateLetter, eventList)
			if prerequisitesMet {
				eventList = append(eventList, candidateLetter)
				lettersList = append(lettersList[:i], lettersList[i+1:]...)
				break
			}
		}
	}
	fmt.Println(strings.Join(eventList, ""))

	// Part 2
	var timeRemainingPerTask = make(map[string]int)
	nbOfWorkers := 5
	eventList = nil
	timeElapsed := -1
	for len(eventList) < 26 {
		// Do a tick
		timeElapsed++

		workersBusy := 0
		for task, timeRemaining := range timeRemainingPerTask {
			workersBusy ++
			timeRemainingPerTask[task]--
			if timeRemaining <= 0 {
				delete(timeRemainingPerTask, task)
				eventList = append(eventList, task)
			}
		}
		if workersBusy < nbOfWorkers {
			// We can queue new tasks if its requirements are met!
			for i := 0; i < len(toProcessSortedList); i++ {
				candidateLetter := toProcessSortedList[i]
				prerequisitesMet := checkPrerequisites(candidateLetter, eventList)
				if prerequisitesMet && workersBusy < nbOfWorkers {
					timeRemainingPerTask[candidateLetter] = int([]rune(candidateLetter)[0]) - 5
					workersBusy++
					toProcessSortedList = append(toProcessSortedList[:i], toProcessSortedList[i+1:]...)
					i=-1
				}
			}
			// No more letters can be processed or we don't have enough workers
		}

	}
	fmt.Print(timeElapsed)
}
