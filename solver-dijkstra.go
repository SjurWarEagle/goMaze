package main

type SolverDijkstra struct {
}

func (solver SolverDijkstra) Solve(maze Maze) {
	current := maze.getStart()
	current.distanceFromStart = 0
	//maze.getCell(current.x, current.y)
	iterations := 0
	cellStack := make(cellStack, 0)
	cellStack = cellStack.Push(current)

	var err error

	for !cellStack.IsEmpty() {
		cellStack, current, err = cellStack.Pop()
		if err != nil {
			panic(err)
		}

		unvisitedOrthogonalNeighbours := maze.GetWalkableOrthogonalNeighbours(current)
		for cnt := 0; cnt < len(unvisitedOrthogonalNeighbours); cnt++ {
			neighbour := unvisitedOrthogonalNeighbours[cnt]
			if neighbour.distanceFromStart == -1 {
				neighbour.distanceFromStart = current.distanceFromStart + 1
				cellStack = cellStack.Push(neighbour)
			}
		}

		iterations++
		if iterations > 1_000_000_000 {
			panic("Too many loops!")
		}
	}
}
