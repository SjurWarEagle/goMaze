package maze_render

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

type PngRenderer struct {
}

func (r PngRenderer) drawFloors(myImage draw.Image, maze maze.Maze, sizeOfCell int, borderWidth int, strokeThickness int) {
	startPosition := maze.GetStart()
	endPosition := maze.GetFinish()

	var maxDistance int
	for x := 0; x < maze.MazeWidth; x++ {
		for y := 0; y < maze.MazeHeight; y++ {
			cell, err := maze.GetCell(x, y)
			if err != nil {
				panic(err)
			}
			maxDistance = int(math.Max(float64(maxDistance), float64(cell.DistanceFromStart)))
		}
	}

	for x := 0; x < maze.MazeWidth; x++ {
		for y := 0; y < maze.MazeHeight; y++ {
			cell, err := maze.GetCell(x, y)
			if err != nil {
				panic(err)
			}
			imageX := borderWidth + x*sizeOfCell
			imageY := borderWidth + y*sizeOfCell
			var floorColor color.RGBA
			floor := image.Rectangle{}
			delta := sizeOfCell/2 + strokeThickness

			if x == startPosition.X && y == startPosition.Y {
				floorColor = color.RGBA{
					R: 0,
					G: 0,
					B: 200,
					A: 255,
				}
				floor = image.Rect(imageX-delta, imageY-delta, imageX+delta, imageY+delta)
			} else if x == endPosition.X && y == endPosition.Y {
				floorColor = color.RGBA{
					R: 0,
					G: 200,
					B: 0,
					A: 255,
				}
				floor = image.Rect(imageX-delta, imageY-delta, imageX+delta, imageY+delta)
			} else if cell.DistanceFromStart == -1 {
				floorColor = color.RGBA{
					R: 1,
					G: 102,
					B: 177,
					A: 255,
				}
				floor = image.Rect(imageX-delta, imageY-delta, imageX+delta, imageY+delta)
			} else {
				floorColor = color.RGBA{
					R: uint8(cell.DistanceFromStart * 255 / maxDistance),
					A: 255,
				}
				floor = image.Rect(imageX-delta, imageY-delta, imageX+delta, imageY+delta)
			}
			draw.Draw(myImage, floor, &image.Uniform{C: floorColor}, image.Point{}, draw.Src)
		}
	}
}

func (r PngRenderer) drawWalls(myImage draw.Image, givenMaze maze.Maze, sizeOfCell int, borderWidth int, strokeThickness int) {

	for x := 0; x < givenMaze.MazeWidth; x++ {
		for y := 0; y < givenMaze.MazeHeight; y++ {
			cell, err := givenMaze.GetCell(x, y)
			if err != nil {
				panic(err)
			}
			imageX := borderWidth + x*sizeOfCell
			imageY := borderWidth + y*sizeOfCell
			wallColor := color.RGBA{
				R: 0,
				G: 30,
				B: 80,
				A: 255,
			}

			//fmt.Printf("distanceFromStart: %d\n", cell.distanceFromStart)
			if cell.Walls[maze.NORTH] {
				wall := image.Rect(imageX-sizeOfCell/2-strokeThickness, imageY-sizeOfCell/2-strokeThickness, imageX+sizeOfCell/2+strokeThickness, imageY-sizeOfCell/2+strokeThickness)
				draw.Draw(myImage, wall, &image.Uniform{C: wallColor}, image.Point{}, draw.Src)
			}
			if cell.Walls[maze.EAST] {
				wall := image.Rect(imageX+sizeOfCell/2-strokeThickness, imageY-sizeOfCell/2-strokeThickness, imageX+sizeOfCell/2+strokeThickness, imageY+sizeOfCell/2+strokeThickness)
				draw.Draw(myImage, wall, &image.Uniform{C: wallColor}, image.Point{}, draw.Src)
			}
			if cell.Walls[maze.SOUTH] {
				wall := image.Rect(imageX-sizeOfCell/2-strokeThickness, imageY+sizeOfCell/2-strokeThickness, imageX+sizeOfCell/2+strokeThickness, imageY+sizeOfCell/2+strokeThickness)
				draw.Draw(myImage, wall, &image.Uniform{C: wallColor}, image.Point{}, draw.Src)
			}
			if cell.Walls[maze.WEST] {
				wall := image.Rect(imageX-sizeOfCell/2-strokeThickness, imageY-sizeOfCell/2-strokeThickness, imageX-sizeOfCell/2+strokeThickness, imageY+sizeOfCell/2+strokeThickness)
				draw.Draw(myImage, wall, &image.Uniform{C: wallColor}, image.Point{}, draw.Src)
			}
		}
	}
}

func (r PngRenderer) Render(maze maze.Maze) {
	file := "./output/pkg.maze.png"
	sizeOfCell := 20
	border := 15
	strokeThickness := 1

	// https://go.dev/blog/image-draw
	myImage := image.NewRGBA(image.Rect(0, 0, sizeOfCell*(maze.MazeWidth-1)+2*border, sizeOfCell*(maze.MazeHeight-1)+2*border)) // x1,y1,  x2,y2 of background rectangle

	// back-fill entire background surface
	backgroundColor := color.RGBA{
		R: 0,
		G: 30,
		B: 80,
		A: 255,
	}

	draw.Draw(myImage, myImage.Bounds(), &image.Uniform{C: backgroundColor}, image.Point{}, draw.Src)

	r.drawFloors(myImage, maze, sizeOfCell, border, strokeThickness)
	r.drawWalls(myImage, maze, sizeOfCell, border, strokeThickness)

	myFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer func(myFile *os.File) {
		err := myFile.Close()
		if err != nil {

		}
	}(myFile)

	err = png.Encode(myFile, myImage)
	if err != nil {
		panic(err)
	}
}
