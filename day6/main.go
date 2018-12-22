package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x      int
	y      int
	id     int
	burned bool
}


var closestPtsMatrix [m][n][]*point

func hammingDist(pt *point, i int, j int) float64 {
	return math.Abs(float64(pt.x-i)) + math.Abs(float64(pt.y-j))
}

const mMax, nMax = 341, 355
const mMin, nMin = 44, 41
const m = mMax - mMin + 1 // m nb of rows in the matrix
const n = nMax - nMin + 1 // n nb of columns in the matrix

func main() {

	file, err := os.Open("day6/data.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var pts []*point

	i := 0
	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		pts = append(pts, &point{x - mMin, y - nMin, i, false})
		i++
	}

	// Populates closestPtMatrix
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			minDist := float64(m * n)

			for _, pt := range pts {
				dist := hammingDist(pt, i, j)

				if dist == minDist {
					closestPtsMatrix[i][j] = append(closestPtsMatrix[i][j], pt)
				}

				if dist < minDist {

					closestPtsMatrix[i][j] = []*point{pt}
					minDist = dist

				}

			}
		}
	}

	// Part 1
	maxArea := 0
	for _, pt := range pts{
		area :=0
		for i := 0; i< m;i++{
			for j:=0;j<n ;j++  {
				if i==0 || j ==0 || i == m-1 || j==n-1 {
					for _,ptToBurn := range closestPtsMatrix[i][j]{
						ptToBurn.burned = true
					}
				}
				if len(closestPtsMatrix[i][j])==1 && closestPtsMatrix[i][j][0].id == pt.id {
					area++
				}
			}
		}
		if area >maxArea && !pt.burned{
			maxArea = area
		}
	}
	println(maxArea)

	// Part 2
	area :=0
	for i := 0; i< m;i++{
		for j:=0;j<n ;j++  {
			summedDist := 0.
			for _, pt := range pts{
				summedDist += hammingDist(pt,i,j)
			}
			if summedDist<10000 {
				area++
			}
		}
	}
	println(area)
}



