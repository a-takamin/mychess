package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type ChessBoard struct {
	w                fyne.Window
	board            *fyne.Container
	trunTeam         string
	selectedTile     *ChessTile
	possibleNextMove []*Move
	whiteKingTile    *ChessTile
	blackKingTile    *ChessTile
	// ToDo
}

func (c *ChessBoard) draw() {
	c.w.SetContent(c.board)
	// c.w.Canvas().Refresh(c.board)
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
	blackKingTile := NewChessTile(4, NewKing("black"), chessBoard)
	chessBoard.whiteKingTile = blackKingTile
	board.Add(blackKingTile)
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
	whiteKingTile := NewChessTile(60, NewKing("white"), chessBoard)
	chessBoard.whiteKingTile = whiteKingTile
	board.Add(whiteKingTile)
	board.Add(NewChessTile(61, NewBishop("white"), chessBoard))
	board.Add(NewChessTile(62, NewKnight("white"), chessBoard))
	board.Add(NewChessTile(63, NewRook("white"), chessBoard))

	chessBoard.board = board
	chessBoard.trunTeam = "white"
	chessBoard.selectedTile = nil
	chessBoard.possibleNextMove = chessBoard.getAllPossibleMove(chessBoard.trunTeam)

	return chessBoard
}

func renewChessBoard(currentChessBoard *ChessBoard, nextMove *Move) {
	newChessBoard := newChessBoard(currentChessBoard, nextMove)
	newChessBoard.possibleNextMove = newChessBoard.getAllPossibleMove(newChessBoard.trunTeam)
	newChessBoard.draw()
}

func newChessBoard(currentChessBoard *ChessBoard, nextMove *Move) *ChessBoard {
	currentTiles := currentChessBoard.board.Objects
	newChessBoard := &ChessBoard{}
	nextBoard := container.New(layout.NewGridLayout(8))

	for _, tile := range currentTiles {
		tileId := tile.(*ChessTile).tileId
		if tileId == nextMove.fromTile.getTileId() {
			// 移動元は空にする
			nextBoard.Add(NewChessTile(tileId, NewNoPiece("noteam"), newChessBoard))
		} else if tileId == nextMove.to {
			// 移動先に移動元の駒を移動させる
			/// 特殊1: ポーン昇格
			if nextMove.fromTile.getPiece().GetPieceType() == "Pawn" && nextMove.fromTile.getPiece().GetPieceTeam() == "white" && nextMove.to >= 0 && nextMove.to <= 7 {
				nextBoard.Add(NewChessTile(tileId, NewQueen("white"), newChessBoard))
			} else if nextMove.fromTile.getPiece().GetPieceType() == "Pawn" && nextMove.fromTile.getPiece().GetPieceTeam() == "black" && nextMove.to >= 56 && nextMove.to <= 63 {
				nextBoard.Add(NewChessTile(tileId, NewQueen("black"), newChessBoard))
				/// TODO: 特殊2: EnPassantロジック
				/// TODO: 特殊3: キャスリングの場合は、キングとルークを同時に移動させる
			} else {
				nextBoard.Add(NewChessTile(tileId, nextMove.fromTile.getPiece(), newChessBoard))
			}
		} else {
			nextBoard.Add(NewChessTile(tileId, tile.(*ChessTile).piece, newChessBoard))
		}
	}
	// kingの位置を把握する
	var whiteKingTile *ChessTile
	var blackKingTile *ChessTile
	if nextMove.fromTile.getPiece().GetPieceType() == "King" {
		if nextMove.fromTile.getPiece().GetPieceTeam() == "white" {
			whiteKingTile = nextBoard.Objects[nextMove.to].(*ChessTile)
			blackKingTile = currentChessBoard.blackKingTile
		} else {
			whiteKingTile = currentChessBoard.whiteKingTile
			blackKingTile = nextBoard.Objects[nextMove.to].(*ChessTile)
		}
	} else {
		whiteKingTile = currentChessBoard.whiteKingTile
		blackKingTile = currentChessBoard.blackKingTile
	}
	var nextTeam string
	if currentChessBoard.trunTeam == "white" {
		nextTeam = "black"
	} else {
		nextTeam = "white"
	}
	newChessBoard.w = currentChessBoard.w
	newChessBoard.board = nextBoard
	newChessBoard.trunTeam = nextTeam
	newChessBoard.selectedTile = nil
	// newChessBoard.possibleNextMove = newChessBoard.getAllPossibleMove(nextTeam)
	newChessBoard.whiteKingTile = whiteKingTile
	newChessBoard.blackKingTile = blackKingTile
	return newChessBoard
}

func isMoveBeInChecked(currentChessBoard *ChessBoard, nextMove *Move) bool {
	// 仮想的に手を実行
	// チェックがかかる手があってしまわないか判定
	supposedNextChessBoard := newChessBoard(currentChessBoard, nextMove)
	// FIXME: 無限ループが起きている
	for _, enemyPossibleNextMove := range supposedNextChessBoard.getAllPossibleMove(supposedNextChessBoard.trunTeam) {
		if enemyPossibleNextMove.to == supposedNextChessBoard.whiteKingTile.getTileId() || enemyPossibleNextMove.to == supposedNextChessBoard.blackKingTile.getTileId() {
			return true
		}
	}
	return false
}

func (c *ChessBoard) getAllPossibleMove(turnTeam string) []*Move {
	var possibleNextMove []*Move
	for _, tile := range c.board.Objects {
		t, _ := tile.(*ChessTile)
		if t.team == turnTeam {
			possibleNextMove = append(possibleNextMove, t.getPiece().CalcPossibleNextMove(t)...)
		}
	}
	// チェックがかかる手を除外する
	// // 初回のみpossibleNextMoveがnilになるので、その場合はチェックを行わない
	// if c.possibleNextMove == nil {
	// 	return possibleNextMove
	// }
	var checkRemovedPossibleNextMove []*Move
	for _, move := range possibleNextMove {
		if isMoveBeInChecked(c, move) {
			// チェック状態になる手は除外する
			continue
		}
		checkRemovedPossibleNextMove = append(checkRemovedPossibleNextMove, move)
	}
	return checkRemovedPossibleNextMove
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
