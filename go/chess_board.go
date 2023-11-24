package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type ChessBoard struct {
	board  *fyne.Container
	pieces map[int]Piece
	// ToDo
}

func (c *ChessBoard) GetBoard() *fyne.Container {
	return c.board
}

func NewInitialChessBoard() *ChessBoard {
	cbb := NewChessBoardBuilder()

	wr1 := NewChessTile(0, NewRook("white"))
	wn1 := NewChessTile(1, NewKnight("white"))
	wb1 := NewChessTile(2, NewBishop("white"))
	wk := NewChessTile(3, NewKing("white"))
	wq := NewChessTile(4, NewQueen("white"))
	wb2 := NewChessTile(5, NewBishop("white"))
	wn2 := NewChessTile(6, NewKnight("white"))
	wr2 := NewChessTile(7, NewRook("white"))
	wp1 := NewChessTile(8, NewPawn("white"))
	wp2 := NewChessTile(9, NewPawn("white"))
	wp3 := NewChessTile(10, NewPawn("white"))
	wp4 := NewChessTile(11, NewPawn("white"))
	wp5 := NewChessTile(12, NewPawn("white"))
	wp6 := NewChessTile(13, NewPawn("white"))
	wp7 := NewChessTile(14, NewPawn("white"))
	wp8 := NewChessTile(15, NewPawn("white"))

	cbb.Set(wr1).Set(wn1).Set(wb1).Set(wk).Set(wq).Set(wb2).Set(wn2).Set(wr2).Set(wp1).Set(wp2).Set(wp3).Set(wp4).Set(wp5).Set(wp6).Set(wp7).Set(wp8)

	bp1 := NewChessTile(48, NewPawn("black"))
	bp2 := NewChessTile(49, NewPawn("black"))
	bp3 := NewChessTile(50, NewPawn("black"))
	bp4 := NewChessTile(51, NewPawn("black"))
	bp5 := NewChessTile(52, NewPawn("black"))
	bp6 := NewChessTile(53, NewPawn("black"))
	bp7 := NewChessTile(54, NewPawn("black"))
	bp8 := NewChessTile(55, NewPawn("black"))
	br1 := NewChessTile(56, NewRook("black"))
	bn1 := NewChessTile(57, NewKnight("black"))
	bb1 := NewChessTile(58, NewBishop("black"))
	bk := NewChessTile(59, NewKing("black"))
	bq := NewChessTile(60, NewQueen("black"))
	bb2 := NewChessTile(61, NewBishop("black"))
	bn2 := NewChessTile(62, NewKnight("black"))
	br2 := NewChessTile(63, NewRook("black"))

	cbb.Set(bp1).Set(bp2).Set(bp3).Set(bp4).Set(bp5).Set(bp6).Set(bp7).Set(bp8).Set(br1).Set(bn1).Set(bb1).Set(bk).Set(bq).Set(bb2).Set(bn2).Set(br2)

	return cbb.Build()
}

func NewChessBoard(tiles map[int]*ChessTile) *ChessBoard {
	chessBoard := &ChessBoard{}
	chessBoard.pieces = make(map[int]Piece)
	c := container.New(layout.NewGridLayout(8))
	chessBoard.board = c

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			position := row*8 + col
			var tile *ChessTile

			if tiles[position] == nil {
				tile = NewChessTile(position, &NoPiece{})
			} else {
				tile = tiles[position]
			}
			c.Add(tile)
			// chessBoard.pieces[position] = tile
		}
	}
	return chessBoard
}
