package main

type Piece interface {
	CalcPossibleNextMove(currentTile *ChessTile) []*Move

	GetPieceType() string
	GetPieceTeam() string
	GetIconPath() string
}
