package main

type SolverDijkstra struct {
}

func (solver SolverDijkstra) FillSolutionDistance(maze Maze) {

}
func (solver SolverDijkstra) Solve(maze Maze) {
	current := maze.getStart()
	current.distanceFromStart = 0
	maze.getCell(current.x, current.y)
	itertations := 0
	cellStack := make(cellStack, 0)
	cellStack = cellStack.Push(current)

	var err error

	for !cellStack.IsEmpty() {
		cellStack, current, err = cellStack.Pop()
		if err != nil {
			panic(err)
		}

		unvisitedOrthogonalNeighbours := maze.GetWalkableOrthogonalNeighbours(current)
		for cnt := range len(unvisitedOrthogonalNeighbours) {
			neighbour := unvisitedOrthogonalNeighbours[cnt]
			if neighbour.distanceFromStart == -1 {
				neighbour.distanceFromStart = current.distanceFromStart + 1
				cellStack = cellStack.Push(neighbour)
			}
		}

		itertations++
		if itertations > 1_000_000_000 {
			panic("Too many loops!")
		}
	}
}

func filter(ss []*Cell, test func(*Cell) bool) (ret []*Cell) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
