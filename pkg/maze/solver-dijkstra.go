package maze

type SolverDijkstra struct {
}

func (solver SolverDijkstra) Solve(maze Maze) {
	current := maze.GetStart()
	current.DistanceFromStart = 0
	//pkg.maze.GetCell(current.X, current.Y)
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
			if neighbour.DistanceFromStart == -1 {
				neighbour.DistanceFromStart = current.DistanceFromStart + 1
				cellStack = cellStack.Push(neighbour)
			}
		}

		iterations++
		if iterations > 1_000_000_000 {
			panic("Too many loops!")
		}
	}
}
