package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
	"log"
	"main/maze"
	"math"
	"os"
)

var myMaze = maze.Maze{}
var maxWidthOfMaze = 30
var desiredWidthOfMaze = new(widget.Float)

func main() {
	go createWindow()
	app.Main()
}

func createWindow() {
	w := new(app.Window)
	w.Option(app.Title("Maze Runner"), // app.Decorated(false),
		app.Size(unit.Dp(400), unit.Dp(600)))
	err := run(w)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var widthSlider = material.Slider(th, desiredWidthOfMaze)
	var ops op.Ops
	// generateButton is a clickable widget
	var generateButton widget.Clickable
	generateMaze()

	for {
		switch eventType := w.Event().(type) {
		// and this is sent when the application should exit
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, eventType)

			updateMazeIfNewWidthIsRequested()

			if generateButton.Clicked(gtx) {
				generateMaze()
			} else {
				layout.Flex{
					// Vertical alignment, from top to bottom
					Axis:    layout.Vertical,
					Spacing: layout.SpaceBetween,
				}.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					label := material.H3(th, "Maze")
					return label.Layout(gtx)
				}), layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return widthSlider.Layout(gtx)
				}), //layout.Rigid(
					//	func(gtx layout.Context) layout.Dimensions {
					//		circle := clip.Ellipse{
					//			// Hard coding the x coordinate. Try resizing the window
					//			// Min: image.Pt(80, 0),
					//			// Max: image.Pt(320, 240),
					//			// Soft coding the x coordinate. Try resizing the window
					//			Min: image.Pt(gtx.Constraints.Max.X/2-120, 0),
					//			Max: image.Pt(gtx.Constraints.Max.X/2+120, 240),
					//		}.Op(gtx.Ops)
					//		circleColor := color.NRGBA{R: 200, A: 255}
					//		paint.FillShape(gtx.Ops, circleColor, circle)
					//		d := image.Point{Y: 400}
					//		return layout.Dimensions{Size: d}
					//	},
					//),
					generateMazeVisuals(gtx, myMaze), layout.Rigid(
						// ... then one to hold an empty spacer
						//	The height of the spacer is 25 Device independent pixels
						layout.Spacer{
							Height: unit.Dp(25),
						}.Layout), layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &generateButton, "Generate")
						return btn.Layout(gtx)
					}))
				eventType.Frame(gtx.Ops)
			}
		}
	}
}

func updateMazeIfNewWidthIsRequested() {
	width := int(math.Max(3, float64(desiredWidthOfMaze.Value*float32(maxWidthOfMaze))))
	if width != myMaze.MazeWidth {
		generateMaze()
	}
}

func generateMaze() {
	width := int(math.Max(3, float64(desiredWidthOfMaze.Value*float32(maxWidthOfMaze))))
	height := int(math.Max(3, float64(desiredWidthOfMaze.Value*float32(maxWidthOfMaze))))
	myMaze.Init(width, height)
	myMaze.SetRandomStartFinish()
	generator := maze.MazeGenerator{}
	generator.Fill(myMaze)
}

func generateMazeVisuals(gtx layout.Context, maze maze.Maze) layout.FlexChild {
	cellWidth := determineCellWidth(gtx, maze)
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		renderCells(gtx, maze, cellWidth)
		//paint.FillShape(gtx.Ops, fillColor, circle)
		return layout.Dimensions{Size: image.Point{Y: maze.MazeHeight * cellWidth}}
	})
}

func determineCellWidth(gtx layout.Context, maze maze.Maze) int {
	maxSize := gtx.Constraints.Max
	widthToUse := maze.MazeWidth + 1
	heightToUse := maze.MazeHeight + 1
	cellWidth := int(math.Min(float64(maxSize.X/widthToUse), float64(maxSize.Y/heightToUse)))
	return cellWidth
}

func renderCells(gtx layout.Context, maze maze.Maze, cellWidth int) {
	for mazeX := 0; mazeX < maze.MazeWidth; mazeX++ {
		for mazeY := 0; mazeY < maze.MazeHeight; mazeY++ {
			cell, err := maze.GetCell(mazeX, mazeY)
			if err != nil {
				panic(err)
			}
			renderFloor(gtx, cell, cellWidth)
		}
	}
	for mazeX := 0; mazeX < maze.MazeWidth; mazeX++ {
		for mazeY := 0; mazeY < maze.MazeHeight; mazeY++ {
			cell, err := maze.GetCell(mazeX, mazeY)
			if err != nil {
				panic(err)
			}
			renderWall(gtx, cell, cellWidth)
		}
	}
}

func renderWall(gtx layout.Context, mazeCell *maze.Cell, cellWidth int) {
	wallColor := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	border := cellWidth / 2
	startX := border + cellWidth/2 + cellWidth*mazeCell.X
	startY := border + cellWidth/2 + cellWidth*mazeCell.Y
	wallThickness := 4

	if mazeCell.Walls[maze.NORTH] {
		wall := clip.Rect{
			Min: image.Pt(startX-cellWidth/2, startY-cellWidth/2),
			Max: image.Pt(startX+cellWidth/2, startY-cellWidth/2+wallThickness/2),
		}.Op()
		paint.FillShape(gtx.Ops, wallColor, wall)
	}
	if mazeCell.Walls[maze.SOUTH] {
		wall := clip.Rect{
			Min: image.Pt(startX-cellWidth/2, startY+cellWidth/2),
			Max: image.Pt(startX+cellWidth/2, startY+cellWidth/2+wallThickness/2),
		}.Op()
		paint.FillShape(gtx.Ops, wallColor, wall)
	}
	if mazeCell.Walls[maze.WEST] {
		wall := clip.Rect{
			Min: image.Pt(startX-cellWidth/2-wallThickness/2, startY-cellWidth/2),
			Max: image.Pt(startX-cellWidth/2, startY+cellWidth/2),
		}.Op()
		paint.FillShape(gtx.Ops, wallColor, wall)
	}
	if mazeCell.Walls[maze.EAST] {
		wall := clip.Rect{
			Min: image.Pt(startX+cellWidth/2, startY-cellWidth/2),
			Max: image.Pt(startX+cellWidth/2+wallThickness/2, startY+cellWidth/2),
		}.Op()
		paint.FillShape(gtx.Ops, wallColor, wall)
	}
}
func renderFloor(gtx layout.Context, mazeCell *maze.Cell, cellWidth int) {
	floorColor := color.NRGBA{R: 230, G: 230, B: 230, A: 255}
	startColor := color.NRGBA{R: 0, G: 230, B: 0, A: 255}
	finishColor := color.NRGBA{R: 0, G: 0, B: 230, A: 255}
	border := cellWidth / 2
	startX := border + cellWidth/2 + cellWidth*mazeCell.X
	startY := border + cellWidth/2 + cellWidth*mazeCell.Y

	isStart := mazeCell.X == myMaze.GetStart().X && mazeCell.Y == myMaze.GetStart().Y
	isFinish := mazeCell.X == myMaze.GetFinish().X && mazeCell.Y == myMaze.GetFinish().Y

	cell := clip.Rect{
		Min: image.Pt(startX-cellWidth/2, startY-cellWidth/2),
		Max: image.Pt(startX+cellWidth/2, startY+cellWidth/2),
	}.Op()
	if isStart {
		paint.FillShape(gtx.Ops, startColor, cell)

	} else if isFinish {
		paint.FillShape(gtx.Ops, finishColor, cell)

	} else {
		paint.FillShape(gtx.Ops, floorColor, cell)

	}

}
