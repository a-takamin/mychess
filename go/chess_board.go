package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type ChessBoard struct {
	board            *fyne.Container
	trunTeam         string
	selectedTile     *ChessTile
	possibleNextMove []*Move
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
	chessBoard.selectedTile = nil

	return chessBoard
}

func renewChessBoard(chessBoard *ChessBoard, nextMove *Move) {

	currentTiles := chessBoard.board.Objects
	board := chessBoard.board
	board.RemoveAll()

	for _, tile := range currentTiles {
		tileId := tile.(*ChessTile).tileId
		if tileId == nextMove.fromTile.getTileId() {
			// 空にする
			board.Add(NewChessTile(tileId, NewNoPiece("noteam"), chessBoard))
		} else if tileId == nextMove.to {
			// 移動先に移動元の駒を移動させる
			board.Add(NewChessTile(tileId, tile.(*ChessTile).piece, chessBoard))
		} else {
			board.Add(NewChessTile(tileId, tile.(*ChessTile).piece, chessBoard))
		}
	}

	board.Refresh()
}

func (c *ChessBoard) getAllPossibleMove(turnTeam string) []*Move {
	var possibleNextMove []*Move
	for _, tile := range c.board.Objects {
		t, _ := tile.(*ChessTile)
		if t.team == turnTeam {
			possibleNextMove = append(possibleNextMove, t.getPiece().CalcPossibleNextMove(t.getTileId())...)
		}
	}
	return possibleNextMove
}

func (c *ChessBoard) GetPossibleNextMove() []*Move {
	return c.possibleNextMove
}
