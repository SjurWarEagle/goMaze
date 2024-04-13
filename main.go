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
	"goMaze/maze"
	"image"
	"image/color"
	"log"
	"os"
)

var myMaze = maze.Maze{}

func main() {
	go createWindow()
	app.Main()
}

func createWindow() {
	w := new(app.Window)
	w.Option(
		app.Title("Maze Runner"),
		// app.Decorated(false),
		app.Size(unit.Dp(400), unit.Dp(600)),
	)
	err := run(w)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	// generateButton is a clickable widget
	var generateButton widget.Clickable
	myMaze := generateMaze()

	for {
		switch eventType := w.Event().(type) {
		// and this is sent when the application should exits
		case app.DestroyEvent:
			os.Exit(0)
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, eventType)

			if generateButton.Clicked(gtx) {
				myMaze = generateMaze()
			} else {
				layout.Flex{
					// Vertical alignment, from top to bottom
					Axis:    layout.Vertical,
					Spacing: layout.SpaceBetween,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							label := material.H3(th, "Maze")
							return label.Layout(gtx)
						},
					),
					//layout.Rigid(
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
					generateMazeVisuals(th, myMaze, ops),
					layout.Rigid(
						// ... then one to hold an empty spacer
						//	The height of the spacer is 25 Device independent pixels
						layout.Spacer{
							Height: unit.Dp(25),
						}.Layout,
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &generateButton, "Generate")
							return btn.Layout(gtx)
						},
					),
				)
				eventType.Frame(gtx.Ops)
			}
		}
	}
}

func generateMaze() maze.Maze {
	myMaze.Init(10, 10)
	generator := maze.MazeGenerator{}
	generator.Fill(myMaze)

	return myMaze
}

func generateMazeVisuals(th *material.Theme, maze maze.Maze, ops op.Ops) layout.FlexChild {
	cellWidth := 50
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			renderCells(gtx, maze, cellWidth, ops)
			//paint.FillShape(gtx.Ops, fillColor, circle)
			return layout.Dimensions{Size: image.Point{Y: maze.MazeHeight * cellWidth}}
		},
	)
}

func renderCells(gtx layout.Context, maze maze.Maze, cellWidth int, ops op.Ops) {
	for mazeX := range maze.MazeWidth {
		for mazeY := range maze.MazeWidth {
			cell, err := maze.GetCell(mazeX, mazeY)
			if err != nil {
				panic(err)
			}
			renderFloor(gtx, cell, cellWidth, ops)
		}
	}
	for mazeX := range maze.MazeWidth {
		for mazeY := range maze.MazeWidth {
			cell, err := maze.GetCell(mazeX, mazeY)
			if err != nil {
				panic(err)
			}
			renderWall(gtx, cell, cellWidth, ops)
		}
	}
}

func renderWall(gtx layout.Context, mazeCell *maze.Cell, cellWidth int, ops op.Ops) {
	wallColor := color.NRGBA{R: 250, G: 0, B: 0, A: 255}
	startX := 100 + cellWidth*mazeCell.X
	startY := 100 + cellWidth*mazeCell.Y
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
func renderFloor(gtx layout.Context, mazeCell *maze.Cell, cellWidth int, ops op.Ops) {
	floorColor := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
	startX := 100 + cellWidth*mazeCell.X
	startY := 100 + cellWidth*mazeCell.Y

	cell := clip.Rect{
		Min: image.Pt(startX-cellWidth/2, startY-cellWidth/2),
		Max: image.Pt(startX+cellWidth/2, startY+cellWidth/2),
	}.Op()
	paint.FillShape(gtx.Ops, floorColor, cell)

}
