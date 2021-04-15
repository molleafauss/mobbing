package main

import (
	"log"
	"os"
	"time"
)

type Grid struct {
	Generation int
	Rows       int
	Cols       int
	Cells      [][]bool
}

func (self Grid) set_alive(row int, col int) {
	self.Cells[row][col] = true
}

func (self Grid) set_dead(row int, col int) {
	self.Cells[row][col] = false
}

func create_grid(rows int, cols int) Grid {
	return Grid{
		Generation: 0,
		Rows:       rows,
		Cols:       cols,
		Cells:      make_grid(rows, cols),
	}
}

func make_grid(rows int, cols int) [][]bool {
	grid := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]bool, cols)
	}
	return grid
}

func (self *Grid) new_generation() bool {
	var changed bool
	changes := 0
	next_gen := make_grid(self.Rows, self.Cols)
	for row := 0; row < self.Rows; row++ {
		for col := 0; col < self.Cols; col++ {
			next_gen[row][col], changed = self.cell_tick(row, col)
			if changed {
				changes++
			}
		}
	}
	self.Cells = next_gen
	self.Generation += 1
	log.Printf("Generation %v, changes found: %v", self.Generation, changes)
	return changes > 0
}

/**
 * returns two booleans: the status of the cell in the next generation ind if it has changed
 */
func (self Grid) cell_tick(row int, col int) (bool, bool) {
	// count the number of alive cells between row -1/+1 and cols -1/+1
	current_status := self.Cells[row][col]
	var new_status bool
	row_start := row - 1
	if row_start < 0 {
		row_start = 0
	}
	row_end := row + 1
	if row_end >= self.Rows {
		row_end -= 1
	}
	col_start := col - 1
	if col_start < 0 {
		col_start = 0
	}
	col_end := col + 1
	if col_end >= self.Cols {
		col_end -= 1
	}
	alive := 0
	for r := row_start; r <= row_end; r++ {
		for c := col_start; c <= col_end; c++ {
			// don't count myself
			if r == row && c == col {
				continue
			}
			if self.Cells[r][c] {
				alive++
			}
		}
	}
	if current_status {
		// Any live cell with two or three neighbors survives.
		new_status = alive == 2 || alive == 3
	} else {
		// Any dead cell with three live neighbors becomes a live cell
		new_status = alive == 3
	}
	// log.Printf("Cell %v,%v: (alive %v) %v -> %v (%v)", row, col, alive, current_status, new_status, new_status != current_status)
	return new_status, new_status != current_status
}

func (self Grid) Display() {
	log.Printf("Generation: %v", self.Generation)
	for r := 0; r < self.Rows; r++ {
		rowText := ""
		for c := 0; c < self.Cols; c++ {
			if self.Cells[r][c] {
				rowText += "* "
			} else {
				rowText += ". "
			}
		}
		log.Println(rowText)
	}
	log.Println()
}

func main() {
	log.SetOutput(os.Stdout)
	grid := create_grid(4, 8)
	grid.set_alive(2, 5)
	grid.set_alive(3, 4)
	grid.set_alive(3, 5)
	grid.Display()

	changes := true
	for changes {
		// flattering Kamil by copy & pasting
		time.Sleep(time.Millisecond * 500)
		changes = grid.new_generation()
		grid.Display()
	}
}
