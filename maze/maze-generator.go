package maze

import "math/rand/v2"

/*
MazeGenerator

	https://en.wikipedia.org/wiki/Maze_generation_algorithm#Iterative_implementation

	A disadvantage of the first approach is a large depth of recursion – in the worst case, the routine may need to recur on every cell of the area being processed, which may exceed the maximum recursion stack depth in many environments. As a solution, the same backtracking method can be implemented with an explicit stack, which is usually allowed to grow much bigger with no harm.

	1. Choose the initial cell, mark it as visited and push it to the stack
	2. While the stack is not empty
	     1. Pop a cell from the stack and make it a current cell
	     2. If the current cell has any neighbours which have not been visited
	         1. Push the current cell to the stack
	         2. Choose one of the unvisited neighbours
	         3. Remove the wall between the current cell and the chosen cell
	         4. Mark the chosen cell as visited and push it to the stack
*/
type MazeGenerator struct {
}

func (generator MazeGenerator) Fill(maze Maze) {
	cellStack := make(cellStack, 0)
	start := maze.GetStart()
	currentCell := start
	cellStack = cellStack.Push(currentCell)

	processStack(maze, cellStack)
}

func processStack(maze Maze, myStack cellStack) {
	var current *Cell
	var err error

	for !myStack.IsEmpty() {
		myStack, current, err = myStack.Pop()
		if err != nil {
		} else {
			unvisitedNeighbours := maze.GetUnvisitedOrthogonalNeighbours(current)
			if len(unvisitedNeighbours) > 0 {
				myStack = myStack.Push(current)
				next := getRandomCell(unvisitedNeighbours)
				maze.RemoveWalls(current, next)
				next.visited = true
				myStack = myStack.Push(next)
			}
		}
	}
}

func getRandomCell(unvisitedNeighbours []*Cell) *Cell {
	// yes 2 lines, cause: readability
	// I'm no fan of too many operations in one line
	index := rand.IntN(len(unvisitedNeighbours))
	return unvisitedNeighbours[index]
}
