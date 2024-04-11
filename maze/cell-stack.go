package maze

import "fmt"

type cellStack []*Cell

func (s cellStack) IsEmpty() bool {
	if len(s) > 0 {
		return false
	}
	return true
}

func (s cellStack) Push(v *Cell) cellStack {
	return append(s, v)
}

func (s cellStack) Pop() (cellStack, *Cell, error) {
	// FIXME: What do we do if the cellStack is empty, though?

	l := len(s)
	if l == 0 {
		return nil, nil, fmt.Errorf("no cells remaining")
	}

	return s[:l-1], s[l-1], nil
}
