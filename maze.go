package main

import "fmt"

type Maze struct{}

var mazeWidth int
var mazeHeight int
var cells [][]Cell

func (m Maze) init(width int, height int) {
	mazeWidth = width
	mazeHeight = height

	newCells := make([][]Cell, mazeWidth)
	for x := range mazeWidth {
		newCells[x] = make([]Cell, mazeHeight)
	}
	cells = newCells

	for x := range width {
		for y := range height {
			cells[x][y] = NewCell(x, y)
		}
	}
	m.fillCellsWithOuterBorder()
}

func (m Maze) fillCellsWithOuterBorder() {
	for x := 0; x < mazeWidth; x++ {
		cells[x][0].walls[NORTH] = true
		cells[x][mazeHeight-1].walls[SOUTH] = true
	}
	for y := 0; y < mazeHeight; y++ {
		cells[0][y].walls[WEST] = true
		cells[mazeWidth-1][y].walls[EAST] = true
	}
}

func (m Maze) getStart() *Cell {
	//TODO define start with random field during init
	return &cells[1][1]
}

func (m Maze) PrettyPrintAllCells() {
	for x := 0; x < mazeWidth; x++ {
		for y := 0; y < mazeHeight; y++ {
			pWalls := &cells[x][y].walls

			if pWalls[NORTH] {
				fmt.Print("N")
			} else if pWalls[NORTH] {
				fmt.Print("_")
			}
			if pWalls[EAST] {
				fmt.Print("E")
			} else if !pWalls[EAST] {
				fmt.Print("_")
			}
			if pWalls[SOUTH] {
				fmt.Print("S")
			} else if !pWalls[SOUTH] {
				fmt.Print("_")
			}
			if pWalls[WEST] {
				fmt.Print("W")
			} else if !pWalls[WEST] {
				fmt.Print("_")
			}
			fmt.Print(",")
		}
		fmt.Print("\n")
	}
}
func (m Maze) PrintAllCells() {
	for x := 0; x < mazeWidth; x++ {
		for y := 0; y < mazeHeight; y++ {
			fmt.Printf("(%2d,%2d): ", x, y)

			pWalls := &cells[x][y].walls
			var wall string
			if pWalls[NORTH] {
				wall = "|"
			} else {
				wall = " "
			}
			fmt.Printf("N:%s", wall)

			if pWalls[EAST] {
				wall = "|"
			} else {
				wall = " "
			}
			fmt.Printf(",E:%s", wall)

			if pWalls[SOUTH] {
				wall = "|"
			} else {
				wall = " "
			}
			fmt.Printf(",S:%s", wall)

			if pWalls[WEST] {
				wall = "|"
			} else {
				wall = " "
			}
			fmt.Printf(",W:%s,\n", wall)
		}
	}
}

func (m Maze) resetVisitedMarker() {
	for x := range mazeWidth {
		for y := range mazeHeight {
			cells[x][y].visited = false
		}
	}
}

func (m Maze) getCell(x int, y int) (*Cell, error) {
	if x < 0 || x > mazeWidth-1 {
		return nil, fmt.Errorf("no cell at x=%d,y=%d", x, y)
	}
	if y < 0 || y > mazeHeight-1 {
		return nil, fmt.Errorf("no cell at x=%d,y=%d", x, y)
	}
	return &cells[x][y], nil
}

func (m Maze) GetWalkableOrthogonalNeighbours(current *Cell) []*Cell {
	var rc []*Cell

	//get western
	cell, err := m.getCell(current.x-1, current.y)
	if err == nil && !cell.blocker && !cell.walls[EAST] {
		rc = append(rc, cell)
	}

	//get eastern
	cell, err = m.getCell(current.x+1, current.y)
	if err == nil && !cell.blocker && !cell.walls[WEST] {
		rc = append(rc, cell)
	}

	//get northern
	cell, err = m.getCell(current.x, current.y-1)
	if err == nil && !cell.blocker && !cell.walls[SOUTH] {
		rc = append(rc, cell)
	}

	//get southern
	cell, err = m.getCell(current.x, current.y+1)
	if err == nil && !cell.blocker && !cell.walls[NORTH] {
		rc = append(rc, cell)
	}
	return rc
}

func (m Maze) GetUnvisitedOrthogonalNeighbours(current *Cell) []*Cell {
	var rc []*Cell

	cell, err := m.getCell(current.x-1, current.y)
	if err == nil && !cell.visited && !cell.blocker {
		rc = append(rc, cell)
	}

	cell, err = m.getCell(current.x+1, current.y)
	if err == nil && !cell.visited && !cell.blocker {
		rc = append(rc, cell)
	}

	cell, err = m.getCell(current.x, current.y-1)
	if err == nil && !cell.visited && !cell.blocker {
		rc = append(rc, cell)
	}

	cell, err = m.getCell(current.x, current.y+1)
	if err == nil && !cell.visited && !cell.blocker {
		rc = append(rc, cell)
	}
	return rc
}

func (m Maze) RemoveWalls(current *Cell, next *Cell) {
	if current.x < next.x {
		// cell1 <- cell2
		current.walls[EAST] = false
		next.walls[WEST] = false
	} else if current.x > next.x {
		// current -> cell2
		current.walls[WEST] = false
		next.walls[EAST] = false
	} else if current.y > next.y {
		// cell1 ^ cell2
		current.walls[NORTH] = false
		next.walls[SOUTH] = false
	} else if current.y < next.y {
		// current V cell2
		current.walls[SOUTH] = false
		next.walls[NORTH] = false
	}

}
