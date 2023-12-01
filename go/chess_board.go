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

	board.Add(NewChessTile(0, NewRook("black"), chessBoard))
	board.Add(NewChessTile(1, NewKnight("black"), chessBoard))
	board.Add(NewChessTile(2, NewBishop("black"), chessBoard))
	board.Add(NewChessTile(3, NewQueen("black"), chessBoard))
	board.Add(NewChessTile(4, NewKing("black"), chessBoard))
	board.Add(NewChessTile(5, NewBishop("black"), chessBoard))
	board.Add(NewChessTile(6, NewKnight("black"), chessBoard))
	board.Add(NewChessTile(7, NewRook("black"), chessBoard))
	board.Add(NewChessTile(8, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(9, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(10, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(11, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(12, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(13, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(14, NewPawn("black"), chessBoard))
	board.Add(NewChessTile(15, NewPawn("black"), chessBoard))

	for i := 16; i < 48; i++ {
		board.Add(NewChessTile(i, NewNoPiece("noteam"), chessBoard))
	}

	board.Add(NewChessTile(48, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(49, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(50, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(51, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(52, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(53, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(54, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(55, NewPawn("white"), chessBoard))
	board.Add(NewChessTile(56, NewRook("white"), chessBoard))
	board.Add(NewChessTile(57, NewKnight("white"), chessBoard))
	board.Add(NewChessTile(58, NewBishop("white"), chessBoard))
	board.Add(NewChessTile(59, NewQueen("white"), chessBoard))
	board.Add(NewChessTile(60, NewKing("white"), chessBoard))
	board.Add(NewChessTile(61, NewBishop("white"), chessBoard))
	board.Add(NewChessTile(62, NewKnight("white"), chessBoard))
	board.Add(NewChessTile(63, NewRook("white"), chessBoard))

	chessBoard.board = board
	chessBoard.trunTeam = "white"
	chessBoard.selectedTile = nil
	chessBoard.possibleNextMove = chessBoard.getAllPossibleMove(chessBoard.trunTeam)

	return chessBoard
}

func renewChessBoard(chessBoard *ChessBoard, nextMove *Move) {

	currentTiles := chessBoard.board.Objects
	board := chessBoard.board
	board.RemoveAll()

	for _, tile := range currentTiles {
		tileId := tile.(*ChessTile).tileId
		if tileId == nextMove.fromTile.getTileId() {
			// 移動元は空にする
			board.Add(NewChessTile(tileId, NewNoPiece("noteam"), chessBoard))
		} else if tileId == nextMove.to {
			// 移動先に移動元の駒を移動させる
			if nextMove.fromTile.getPiece().GetPieceType() == "Pawn" && nextMove.fromTile.getPiece().GetPieceTeam() == "white" && nextMove.to >= 0 && nextMove.to <= 7 {
				board.Add(NewChessTile(tileId, NewQueen("white"), chessBoard))
			} else if nextMove.fromTile.getPiece().GetPieceType() == "Pawn" && nextMove.fromTile.getPiece().GetPieceTeam() == "black" && nextMove.to >= 56 && nextMove.to <= 63 {
				board.Add(NewChessTile(tileId, NewQueen("black"), chessBoard))
			} else {
				board.Add(NewChessTile(tileId, nextMove.fromTile.getPiece(), chessBoard))
			}
		} else {
			board.Add(NewChessTile(tileId, tile.(*ChessTile).piece, chessBoard))
		}
	}
	// 再描画
	board.Refresh()
	// turnTeamの変更
	if chessBoard.trunTeam == "white" {
		chessBoard.trunTeam = "black"
	} else {
		chessBoard.trunTeam = "white"
	}
	// possibleMoveの更新
	chessBoard.possibleNextMove = chessBoard.getAllPossibleMove(chessBoard.trunTeam)
}

func (c *ChessBoard) getAllPossibleMove(turnTeam string) []*Move {
	var possibleNextMove []*Move
	for _, tile := range c.board.Objects {
		t, _ := tile.(*ChessTile)
		if t.team == turnTeam {
			possibleNextMove = append(possibleNextMove, t.getPiece().CalcPossibleNextMove(t)...)
		}
	}
	// possibleNextMoveのうち、敵駒の存在によって行けないマスを除外する
	// チェックであれば、チェックを回避する手以外を除外する
	return possibleNextMove
}

func (c *ChessBoard) GetPossibleNextMove() []*Move {
	return c.possibleNextMove
}

func (c *ChessBoard) getChessTile(tileId int) *ChessTile {
	for _, tile := range c.board.Objects {
		t, _ := tile.(*ChessTile)
		if t.tileId == tileId {
			return t
		}
	}
	panic("tile is not found")
}
