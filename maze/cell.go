package maze

type Cell struct {
	Walls             [4]bool
	X                 int
	Y                 int
	visited           bool
	blocker           bool
	DistanceFromStart int
}

func NewCell(givenX int, givenY int) Cell {
	cell := Cell{Walls: [4]bool{true, true, true, true}, X: givenX, Y: givenY}
	cell.DistanceFromStart = -1
	//cell := Cell{Walls: [4]bool{false, false, false, false}, X: givenX, Y: givenY}
	return cell
}
