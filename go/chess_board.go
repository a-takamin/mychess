package main

import (
	"fmt"
	"log"

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
	enPassantTarget  *ChessTile
	// ToDo
}

func (c *ChessBoard) draw() {
	c.w.SetContent(c.board)
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
	chessBoard.blackKingTile = blackKingTile
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
	chessBoard.possibleNextMove = chessBoard.getAllPossibleMove(chessBoard.trunTeam, false)

	return chessBoard
}

func renewChessBoard(currentChessBoard *ChessBoard, nextMove *Move) {
	newChessBoard := newChessBoard(currentChessBoard, nextMove, true)
	newChessBoard.possibleNextMove = newChessBoard.getAllPossibleMove(newChessBoard.trunTeam, true)
	newChessBoard.draw()
}

func newChessBoard(currentChessBoard *ChessBoard, nextMove *Move, considerCheckMove bool) *ChessBoard {
	currentTiles := currentChessBoard.board.Objects
	newChessBoard := &ChessBoard{}
	nextBoard := container.New(layout.NewGridLayout(8))
	newChessBoard.board = nextBoard
	fromTileId := nextMove.fromTile.getTileId()
	var enPassantAttackedPawn *ChessTile = nil

	for _, tile := range currentTiles {
		tileId := tile.(*ChessTile).tileId

		if tileId == fromTileId {
			// 移動元は空にする
			nextBoard.Add(NewChessTile(tileId, NewNoPiece("noteam"), newChessBoard))

		} else if tileId == nextMove.to {
			// 移動先に移動元の駒を移動させる
			if nextMove.fromTile.getPiece().GetPieceType() == "Pawn" { // 移動駒がポーンの場合
				// 初期として通常移動の場合をセットしておく
				newChessTile := NewChessTile(tileId, nextMove.fromTile.getPiece(), newChessBoard)
				movingPawn := nextMove.fromTile.getPiece().(*Pawn)

				/// ポーン特殊0: 2マス移動時はアンパッサン攻撃対象になる
				if movingPawn.isFirstMove(fromTileId) && (nextMove.to-fromTileId == 16 || nextMove.to-fromTileId == -16) {
					newChessBoard.enPassantTarget = newChessTile
				}

				/// ポーン特殊1: 昇格
				if movingPawn.GetPieceTeam() == "white" && nextMove.to >= 0 && nextMove.to <= 7 {
					newChessTile = NewChessTile(tileId, NewQueen("white"), newChessBoard)
				} else if movingPawn.GetPieceTeam() == "black" && nextMove.to >= 56 && nextMove.to <= 63 {
					newChessTile = NewChessTile(tileId, NewQueen("black"), newChessBoard)
				}

				/// ポーン特殊2: EnPassant攻撃
				if currentChessBoard.enPassantTarget != nil {
					if movingPawn.GetPieceTeam() == "white" && currentChessBoard.enPassantTarget.getPiece().GetPieceTeam() == "black" {
						if currentChessBoard.enPassantTarget.getTileId()-8 == nextMove.to {
							enPassantAttackedPawn = currentChessBoard.enPassantTarget
						}
					}
					if movingPawn.GetPieceTeam() == "black" && currentChessBoard.enPassantTarget.getPiece().GetPieceTeam() == "white" {
						if currentChessBoard.enPassantTarget.getTileId()+8 == nextMove.to {
							enPassantAttackedPawn = currentChessBoard.enPassantTarget
						}
					}
				}

				// ポーンを移動させる
				nextBoard.Add(newChessTile)
			} else {
				nextBoard.Add(NewChessTile(tileId, nextMove.fromTile.getPiece(), newChessBoard))
			}

			// TODO: 特殊: キャスリングの場合は、キングとルークを同時に移動させる
			/// キャスリングができるかどうかの判定はここではやらない
		} else {
			// タイルそのまま
			nextBoard.Add(NewChessTile(tileId, tile.(*ChessTile).piece, newChessBoard))
		}
	}

	// アンパッサン攻撃だった場合は攻撃されたポーンを取り除く
	if enPassantAttackedPawn != nil {
		noPiece := NewNoPiece("noteam")
		newChessBoard.getChessTile(enPassantAttackedPawn.getTileId()).piece = noPiece
		newChessBoard.getChessTile(enPassantAttackedPawn.getTileId()).updateImage("./image/bb.png")
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
	newChessBoard.trunTeam = nextTeam
	newChessBoard.selectedTile = nil
	newChessBoard.whiteKingTile = whiteKingTile
	newChessBoard.blackKingTile = blackKingTile
	newChessBoard.possibleNextMove = newChessBoard.getAllPossibleMove(nextTeam, considerCheckMove)
	if len(newChessBoard.possibleNextMove) == 0 {
		fmt.Println("CheckMate!")
	}
	return newChessBoard
}

func isMoveBeInChecked(currentChessBoard *ChessBoard, nextMove *Move) bool {
	// 仮想的に手を実行
	// チェックがかかる手があってしまわないか判定
	supposedNextChessBoard := newChessBoard(currentChessBoard, nextMove, false)
	for _, enemyPossibleNextMove := range supposedNextChessBoard.GetPossibleNextMove() {
		if enemyPossibleNextMove.to == supposedNextChessBoard.whiteKingTile.getTileId() || enemyPossibleNextMove.to == supposedNextChessBoard.blackKingTile.getTileId() {
			return true
		}
	}
	return false
}

func (c *ChessBoard) getAllPossibleMove(turnTeam string, considerCheck bool) []*Move {
	var possibleNextMove []*Move
	for _, tile := range c.board.Objects {
		t, _ := tile.(*ChessTile)
		if t.team == turnTeam {
			possibleNextMove = append(possibleNextMove, t.getPiece().CalcPossibleNextMove(t)...)
		}
	}

	if !considerCheck {
		return possibleNextMove
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
	// キャスリング
	// 1. kingが初期位置である
	// 2. ルークが初期位置である
	// 3. kingの通り道がチェックされていない
	// 4. kingのキャスリング位置がチェックされていない
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
	log.Fatal("tile is not found")
	return nil
}
