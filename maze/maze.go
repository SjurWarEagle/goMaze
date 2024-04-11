package maze

import "fmt"

type Maze struct {
	MazeWidth  int
	MazeHeight int
	start      [2]int
	finish     [2]int
	cells      [][]Cell
}

func (m *Maze) Init(width int, height int) {
	m.MazeWidth = width
	m.MazeHeight = height

	newCells := make([][]Cell, m.MazeWidth)
	for x := 0; x < m.MazeWidth; x++ {
		newCells[x] = make([]Cell, m.MazeHeight)
	}
	m.cells = newCells

	for x := 0; x < m.MazeWidth; x++ {
		for y := 0; y < m.MazeHeight; y++ {
			m.cells[x][y] = NewCell(x, y)
		}
	}
	m.fillCellsWithOuterBorder()
}

func (m *Maze) fillCellsWithOuterBorder() {
	for x := 0; x < m.MazeWidth; x++ {
		m.cells[x][0].Walls[NORTH] = true
		m.cells[x][m.MazeHeight-1].Walls[SOUTH] = true
	}
	for y := 0; y < m.MazeHeight; y++ {
		m.cells[0][y].Walls[WEST] = true
		m.cells[m.MazeWidth-1][y].Walls[EAST] = true
	}
}

func (m *Maze) GetStart() *Cell {
	//TODO define start with random field during Init
	return &m.cells[1][1]
}

func (m *Maze) GetFinish() *Cell {
	//TODO define start with random field during Init
	return &m.cells[m.MazeWidth-2][m.MazeHeight-2]
}

func (m *Maze) PrettyPrintAllCells() {
	for x := 0; x < m.MazeWidth; x++ {
		for y := 0; y < m.MazeHeight; y++ {
			pWalls := &m.cells[x][y].Walls

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
func (m *Maze) PrintAllCells() {
	for x := 0; x < m.MazeWidth; x++ {
		for y := 0; y < m.MazeHeight; y++ {
			fmt.Printf("(%2d,%2d): ", x, y)

			pWalls := &m.cells[x][y].Walls
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

func (m *Maze) resetVisitedMarker() {
	for x := 0; x < m.MazeWidth; x++ {
		for y := 0; y < m.MazeHeight; y++ {
			m.cells[x][y].visited = false
		}
	}
}

func (m *Maze) GetCell(x int, y int) (*Cell, error) {
	if x < 0 || x > m.MazeWidth-1 {
		return nil, fmt.Errorf("no cell at X=%d,Y=%d", x, y)
	}
	if y < 0 || y > m.MazeHeight-1 {
		return nil, fmt.Errorf("no cell at X=%d,Y=%d", x, y)
	}
	return &m.cells[x][y], nil
}

func (m *Maze) GetWalkableOrthogonalNeighbours(current *Cell) []*Cell {
	var rc []*Cell

	//get western
	cell, err := m.GetCell(current.X-1, current.Y)
	if err == nil && !cell.blocker && !cell.Walls[EAST] {
		rc = append(rc, cell)
	}

	//get eastern
	cell, err = m.GetCell(current.X+1, current.Y)
	if err == nil && !cell.blocker && !cell.Walls[WEST] {
		rc = append(rc, cell)
	}

	//get northern
	cell, err = m.GetCell(current.X, current.Y-1)
	if err == nil && !cell.blocker && !cell.Walls[SOUTH] {
		rc = append(rc, cell)
	}

	//get southern
	cell, err = m.GetCell(current.X, current.Y+1)
	if err == nil && !cell.blocker && !cell.Walls[NORTH] {
		rc = append(rc, cell)
	}
	return rc
}

func (m *Maze) GetUnvisitedOrthogonalNeighbours(current *Cell) []*Cell {
	var rc []*Cell

	cell, err := m.GetCell(current.X-1, current.Y)
	if err == nil && !cell.visited && !cell.blocker {
		rc = append(rc, cell)
	}

	cell, err = m.GetCell(current.X+1, current.Y)
	if err == nil && !cell.visited && !cell.blocker {
		rc = append(rc, cell)
	}

	cell, err = m.GetCell(current.X, current.Y-1)
	if err == nil && !cell.visited && !cell.blocker {
		rc = append(rc, cell)
	}

	cell, err = m.GetCell(current.X, current.Y+1)
	if err == nil && !cell.visited && !cell.blocker {
		rc = append(rc, cell)
	}
	return rc
}

func (m *Maze) RemoveWalls(current *Cell, next *Cell) {
	if current.X < next.X {
		// cell1 <- cell2
		current.Walls[EAST] = false
		next.Walls[WEST] = false
	} else if current.X > next.X {
		// current -> cell2
		current.Walls[WEST] = false
		next.Walls[EAST] = false
	} else if current.Y > next.Y {
		// cell1 ^ cell2
		current.Walls[NORTH] = false
		next.Walls[SOUTH] = false
	} else if current.Y < next.Y {
		// current V cell2
		current.Walls[SOUTH] = false
		next.Walls[NORTH] = false
	}

}
