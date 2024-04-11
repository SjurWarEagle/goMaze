package main

import (
	"fmt"
	maze_package "goMaze/maze"
	maze_render "goMaze/maze/render"
	"log"
	"net/http"
	"os"
)

func main() {
	generateMaze()
	handleRequests()
}

func generateMaze() {
	//	maze.resetVisitedMarker()
	//	solver := SolverDijkstra{}
	//	solver.Solve(maze)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(res http.ResponseWriter, req *http.Request) {
	maze := maze_package.Maze{}
	maze.Init(15, 15)
	generator := maze_package.MazeGenerator{}
	generator.Fill(maze)
	renderer := maze_render.PngRenderer{}
	renderer.Render(maze)
	file := "./output/maze.png"
	myFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Length", fmt.Sprint(len(myFile)))
	res.Header().Set("Content-Type", "image/png")
	_, err = res.Write(myFile)
	if err != nil {
		panic(err)
	}
}
