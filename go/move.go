package main

type Move struct {
	fromTile *ChessTile
	to       int
}

func NewMove(target *ChessTile, to int) *Move {
	return &Move{
		fromTile: target,
		to:       to,
	}
}
