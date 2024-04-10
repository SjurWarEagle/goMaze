package main

func main() {
	maze := Maze{}
	maze.init(200, 200)
	generator := MazeGenerator{}
	generator.fill(maze)

	//maze.PrettyPrintAllCells()

	maze.resetVisitedMarker()
	solver := SolverDijkstra{}
	solver.Solve(maze)

	renderer := PngRenderer{}
	renderer.render(maze)
}
