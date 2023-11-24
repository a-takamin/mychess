package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type ChessBoard struct {
	board    *fyne.Container
	trunTeam string
	// ToDo
}

func (c *ChessBoard) GetBoard() *fyne.Container {
	return c.board
}

func NewInitialChessBoard() *ChessBoard {

	chessBoard := &ChessBoard{}
	board := container.New(layout.NewGridLayout(8))

	board.Add(NewChessTile(0, NewRook("white"), chessBoard))
	board.Add(NewChessTile(1, NewKnight("white"), chessBoard))
	board.Add(NewChessTile(2, NewBishop("white"), chessBoard))
	board.Add(NewChessTile(3, NewKing("white"), chessBoard))
	board.Add(NewChessTile(4, NewQueen("white"), chessBoard))
	board.Add(NewChessTile(5, NewBishop("white"), chessBoard))
	board.Add(NewChessTile(6, NewKnight("white"), chessBoard))
	board.Add(NewChessTile(7, NewRook("white"), chessBoard))
	board.Add(NewChessTile(8, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(9, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(10, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(11, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(12, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(13, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(14, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(15, NewPawn("white"), chessBoard))

	for i := 16; i < 48; i++ {
		board.Add(NewChessTile(i, NewNoPiece("noteam"), chessBoard))
	}

	board.Add(NewChessTile(48, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(49, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(50, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(51, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(52, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(53, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(54, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(55, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(56, NewRook("black"), chessBoard))
	board.Add(NewChessTile(57, NewKnight("black"), chessBoard))
	board.Add(NewChessTile(58, NewBishop("black"), chessBoard))
	board.Add(NewChessTile(59, NewKing("black"), chessBoard))
	board.Add(NewChessTile(60, NewQueen("black"), chessBoard))
	board.Add(NewChessTile(61, NewBishop("black"), chessBoard))
	board.Add(NewChessTile(62, NewKnight("black"), chessBoard))
	board.Add(NewChessTile(63, NewRook("black"), chessBoard))

	chessBoard.board = board
	chessBoard.trunTeam = "white"

	return chessBoard
}

func NewChessBoard(currentChessBoard *ChessBoard, nextMove Move) *ChessBoard {

	tiles := currentChessBoard.board.Objects
	newChessBoard := &ChessBoard{}
	board := container.New(layout.NewGridLayout(8))
	newChessBoard.trunTeam = currentChessBoard.trunTeam

	for _, tile := range tiles {
		tileId := tile.(*ChessTile).tileId
		if tileId == nextMove.from {
			board.Add(NewChessTile(tileId, NewNoPiece("noteam"), newChessBoard))
		} else if tileId == nextMove.to {
			board.Add(NewChessTile(tileId, tile.(*ChessTile).piece, newChessBoard))
		} else {
			board.Add(NewChessTile(tileId, tile.(*ChessTile).piece, newChessBoard))
		}
	}

	newChessBoard.board = board
	return newChessBoard
}
