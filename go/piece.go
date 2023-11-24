package main

type Piece interface {
	Move() int // クリック時に呼ばれる
	GetPieceType() string
	GetPieceTeam() string
	CalcNextPosition(currentPosition int) []int
	GetIconPath() string
}
