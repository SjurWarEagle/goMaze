package main

func main() {
	maze := Maze{}
	maze.init(15, 15)
	generator := MazeGenerator{}
	generator.fill(maze)

	//	maze.resetVisitedMarker()
	//	solver := SolverDijkstra{}
	//	solver.Solve(maze)

	renderer := PngRenderer{}
	renderer.render(maze)
}
