package main

type Cell struct {
	walls             [4]bool
	x                 int
	y                 int
	visited           bool
	blocker           bool
	distanceFromStart int
}

func NewCell(givenX int, givenY int) Cell {
	cell := Cell{walls: [4]bool{true, true, true, true}, x: givenX, y: givenY}
	cell.distanceFromStart = -1
	//cell := Cell{walls: [4]bool{false, false, false, false}, x: givenX, y: givenY}
	return cell
}
