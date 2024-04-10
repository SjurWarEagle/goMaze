package main

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

func (r PngRenderer) drawFloors(myImage draw.Image, maze Maze, sizeOfCell int, borderWidth int, strokeThickness int) {
	var maxDistance int
	for x := 0; x < mazeWidth; x++ {
		for y := 0; y < mazeHeight; y++ {
			cell, err := maze.getCell(x, y)
			if err != nil {
				panic(err)
			}
			maxDistance = int(math.Max(float64(maxDistance), float64(cell.distanceFromStart)))
		}
	}

	for x := 0; x < mazeWidth; x++ {
		for y := 0; y < mazeHeight; y++ {
			cell, err := maze.getCell(x, y)
			if err != nil {
				panic(err)
			}
			imageX := borderWidth + x*sizeOfCell
			imageY := borderWidth + y*sizeOfCell
			floorColor := color.RGBA{
				R: uint8(cell.distanceFromStart * 255 / maxDistance),
				A: 255,
			}

			floor := image.Rect(imageX-sizeOfCell/2-strokeThickness, imageY-sizeOfCell/2-strokeThickness, imageX+sizeOfCell/2+strokeThickness, imageY+sizeOfCell/2+strokeThickness)
			draw.Draw(myImage, floor, &image.Uniform{C: floorColor}, image.Point{}, draw.Src)
		}
	}
}

func (r PngRenderer) drawWalls(myImage draw.Image, maze Maze, sizeOfCell int, borderWidth int, strokeThickness int) {

	for x := 0; x < mazeWidth; x++ {
		for y := 0; y < mazeHeight; y++ {
			cell, err := maze.getCell(x, y)
			if err != nil {
				panic(err)
			}
			imageX := borderWidth + x*sizeOfCell
			imageY := borderWidth + y*sizeOfCell
			wallColor := color.Black

			//fmt.Printf("distanceFromStart: %d\n", cell.distanceFromStart)
			if cell.walls[NORTH] {
				wall := image.Rect(imageX-sizeOfCell/2-strokeThickness, imageY-sizeOfCell/2-strokeThickness, imageX+sizeOfCell/2+strokeThickness, imageY-sizeOfCell/2+strokeThickness)
				draw.Draw(myImage, wall, &image.Uniform{C: wallColor}, image.Point{}, draw.Src)
			}
			if cell.walls[EAST] {
				wall := image.Rect(imageX+sizeOfCell/2-strokeThickness, imageY-sizeOfCell/2-strokeThickness, imageX+sizeOfCell/2+strokeThickness, imageY+sizeOfCell/2+strokeThickness)
				draw.Draw(myImage, wall, &image.Uniform{C: wallColor}, image.Point{}, draw.Src)
			}
			if cell.walls[SOUTH] {
				wall := image.Rect(imageX-sizeOfCell/2-strokeThickness, imageY+sizeOfCell/2-strokeThickness, imageX+sizeOfCell/2+strokeThickness, imageY+sizeOfCell/2+strokeThickness)
				draw.Draw(myImage, wall, &image.Uniform{C: wallColor}, image.Point{}, draw.Src)
			}
			if cell.walls[WEST] {
				wall := image.Rect(imageX-sizeOfCell/2-strokeThickness, imageY-sizeOfCell/2-strokeThickness, imageX-sizeOfCell/2+strokeThickness, imageY+sizeOfCell/2+strokeThickness)
				draw.Draw(myImage, wall, &image.Uniform{C: wallColor}, image.Point{}, draw.Src)
			}
		}
	}
}

func (r PngRenderer) render(maze Maze) {
	file := "./maze.png"
	sizeOfCell := 20
	border := 15
	strokeThickness := 1

	// https://go.dev/blog/image-draw
	myImage := image.NewRGBA(image.Rect(0, 0, sizeOfCell*mazeWidth+2*border, sizeOfCell*mazeHeight+2*border)) // x1,y1,  x2,y2 of background rectangle

	// back-fill entire background surface
	draw.Draw(myImage, myImage.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)

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
