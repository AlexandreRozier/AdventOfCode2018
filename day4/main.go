package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Guard struct {
	ID                   int
	totalTimeAsleepInMin int
	sleepDurationPerMin  [60]int
}

func main() {
	//[1518-04-29 23:59] Guard #463 begins shift
	file, err := os.Open("day4/data.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	type timeAndEntry struct {
		time  time.Time
		entry []string
	}

	var orderedDF []timeAndEntry
	for scanner.Scan() {
		line := scanner.Text()
		timestamp := line[1:17]

		splitted := strings.Split(line, " ")

		form := "2006-01-02 15:04"
		t, _ := time.Parse(form, timestamp)
		orderedDF = append(orderedDF, timeAndEntry{t, splitted[2:]})
	}
	file.Close()

	// Sort dataset
	sort.Slice(orderedDF, func(i, j int) bool { return orderedDF[i].time.Before(orderedDF[j].time) })

	guardByID := make(map[int]*Guard)
	currentGuardID := 0
	var fellAsleepTime time.Time
	for _, entry := range orderedDF {
		switch entry.entry[0] {
		case "Guard":
			currentGuardID, _ = strconv.Atoi(entry.entry[1][1:])
			// Initializes every Guard
			if guardByID[currentGuardID] == nil {
				guardByID[currentGuardID] = &Guard{ID: currentGuardID}
			}

		case "falls":
			fellAsleepTime = entry.time

		case "wakes":
			timeElapsed := entry.time.Sub(fellAsleepTime)
			minutesAsleep := int(timeElapsed.Minutes())
			guardByID[currentGuardID].totalTimeAsleepInMin += int(minutesAsleep)

			for min := 0; min < minutesAsleep; min++ {
				idx := (fellAsleepTime.Minute() + min) % 60
				guardByID[currentGuardID].sleepDurationPerMin[idx]++
			}
		}
	}

	// Part 1
	var worstGuard *Guard
	maxSleepDuration := 0
	for _, guard := range guardByID {

		if guard.totalTimeAsleepInMin > maxSleepDuration {
			worstGuard = guard
			maxSleepDuration = guard.totalTimeAsleepInMin
		}
	}
	var worstMin int
	var worstDuration int

	for i, sleepDuration := range worstGuard.sleepDurationPerMin {
		if sleepDuration > worstDuration {
			worstMin = i
			worstDuration = sleepDuration
		}
	}
	fmt.Println(worstMin * worstGuard.ID)

	// Part 2
	maxSleepDuration = 0
	worstGuard = nil
	worstMin = 0
	for _, guard := range guardByID {
		for min, sleepDuration := range  guard.sleepDurationPerMin{
			if sleepDuration > maxSleepDuration {
				maxSleepDuration = sleepDuration
				worstGuard = guard
				worstMin = min
			}
		}
	}
	fmt.Println(worstMin*worstGuard.ID)
}
