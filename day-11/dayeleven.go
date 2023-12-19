package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputData.txt
var input string

type Position struct {
	y, x, id int
}

func main() {

	fmt.Println("Part 1 =", partOne())
	fmt.Println("Part 2 =", partTwo())
}

func partOne() int {
	observation := parse()
	space := expand(observation)

	sum := 0
	galaxies := toPositions(space)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
		}
	}

	return sum
}

func partTwo() int {
	observation := parse()
	
	rows, cols := findSpace(observation, 1000000)

	sum := 0
	galaxies := toPositionsTwo(observation, rows, cols)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
		}
	}
	return sum
}

func findSpace(grid [][]int, scale int) (map[int]int, map[int]int) {
	rows := map[int]int{}
	emptyTop := 0
	for row := 0; row < len(grid); row++ {
		rows[row] = row + emptyTop*(scale-1)
		if isEmptyRow(grid, row) {
			emptyTop++
		}
	}

	cols := map[int]int{}
	emptyLeft := 0
	for col := 0; col < len(grid[0]); col++ {
		cols[col] = col + emptyLeft*(scale-1)
		if isEmptyCol(grid, col) {
			emptyLeft++
		}
	}

	return rows, cols
}
func toPositions(grid [][]int) []Position {
	var result []Position
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 0 {
				continue
			}
			result = append(result, Position{row, col, grid[row][col]})
		}
	}
	return result
}

func toPositionsTwo(grid [][]int, rows, cols map[int]int) []Position {
	var result []Position

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 0 {
				continue
			}
			result = append(result, Position{
				rows[row],
				cols[col],
				grid[row][col],
			})
		}
	}
	return result
}

func expand(grid [][]int) [][]int {
	for row := 0; row < len(grid); row++ {
		if isEmptyRow(grid, row) {
			grid = append(grid[:row+1], grid[row:]...)
			row++
		}
	}

	for col := 0; col < len(grid[0]); col++ {
		if isEmptyCol(grid, col) {
			for row := 0; row < len(grid); row++ {
				grid[row] = append(grid[row][:col+1], grid[row][col:]...)
			}
			col++
		}
	}
	return grid
}

func isEmptyRow(grid [][]int, row int) bool {
	for col := 0; col < len(grid[0]); col++ {
		if grid[row][col] != 0 {
			return false
		}
	}
	return true
}

func isEmptyCol(grid [][]int, col int) bool {
	for row := 0; row < len(grid); row++ {
		if grid[row][col] != 0 {
			return false
		}
	}
	return true
}

func parse() [][]int {
	id := 1
	var result [][]int
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		var row []int
		for _, r := range strings.TrimSpace(s) {
			if r == '.' {
				row = append(row, 0)
			} else {
				row = append(row, id)
				id++
			}
		}
		result = append(result, row)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}