package main

func main() {
	maze := Maze{}
	maze.init(10, 10)
	generator := MazeGenerator{}
	generator.fill(maze)

	//	maze.resetVisitedMarker()
	//	solver := SolverDijkstra{}
	//	solver.Solve(maze)

	renderer := PngRenderer{}
	renderer.render(maze)
}
