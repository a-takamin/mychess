package main

type Queen struct {
	BasePiece
}

func NewQueen(team string) *Queen {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bq.png"
	} else {
		iconPath = "./image/wq.png"
	}
	return &Queen{
		BasePiece: BasePiece{
			pieceType: "Queen",
			pieceTeam: team,
			iconPath:  iconPath,
		},
	}
}

func (p *Queen) CalcPossibleNextMove(currentTile *ChessTile) []*Move {
	possibleNextMove := []*Move{}
	pos := currentTile.getTileId()
	leftMove := -1
	rightMove := 1
	upMove := -8
	downMove := 8
	// 上
	if !isFirstRank(pos) {
		for i := pos + upMove; isValidPos(i); i += upMove {
			possibleMovePosTeam := currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam()
			if possibleMovePosTeam == "noteam" {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			} else if possibleMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
				break
			} else {
				break
			}
		}
	}
	// 下
	if !isEighthRank(pos) {
		for i := pos + downMove; isValidPos(i); i += downMove {
			possibleMovePosTeam := currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam()
			if possibleMovePosTeam == "noteam" {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			} else if possibleMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
				break
			} else {
				break
			}
		}
	}
	// 左
	if !isFirstColumn(pos) {
		// 第8カラムにワープするまで
		for i := pos + leftMove; isValidPos(i) && !isEighthColumn(i); i += leftMove {
			possibleMovePosTeam := currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam()
			if possibleMovePosTeam == "noteam" {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			} else if possibleMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
				break
			} else {
				break
			}
		}
	}
	// 右
	if !isEighthColumn(pos) {
		// 第1カラムにワープするまで
		for i := pos + rightMove; isValidPos(i) && !isFirstColumn(i); i += rightMove {
			possibleMovePosTeam := currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam()
			if possibleMovePosTeam == "noteam" {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			} else if possibleMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
				break
			} else {
				break
			}
		}
	}

	leftUpMove := -9
	rightUpMove := -7
	leftDownMove := 7
	rightDownMove := 9
	// 左上
	for i := pos + leftUpMove; i >= 0; i += leftUpMove {
		if isEighthColumn(i) {
			// 動いた結果8列目にいる = 左上Moveとしてはありえない（もともと左端の1列目にいた）
			break
		}
		if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() == "noteam" {
			// 空きマス
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
		} else if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != p.pieceTeam {
			// 敵駒マスも移動候補。でもそれ以上は進めない。
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			break
		} else {
			break
		}
	}
	// 右上
	for i := pos + rightUpMove; i >= 0; i += rightUpMove {
		if isFirstColumn(i) {
			// 動いた結果1列目にいる = 右上Moveとしてはありえない（もともと右端の8列目にいた）
			break
		}
		if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() == "noteam" {
			// 空きマス
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
		} else if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != p.pieceTeam {
			// 敵駒マスも移動候補。でもそれ以上は進めない。
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			break
		} else {
			break
		}
	}
	// 左下
	for i := pos + leftDownMove; i <= 63; i += leftDownMove {
		if isEighthColumn(i) {
			// 動いた結果8列目にいる = 左下Moveとしてはありえない（もともと左端の1列目にいた）
			break
		}
		if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() == "noteam" {
			// 空きマス
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
		} else if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != p.pieceTeam {
			// 敵駒マスも移動候補。でもそれ以上は進めない。
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			break
		} else {
			break
		}
	}
	// 右下
	for i := pos + rightDownMove; i <= 63; i += rightDownMove {
		if isFirstColumn(i) {
			// 動いた結果1列目にいる = 右下Moveとしてはありえない（もともと右端の8列目にいた）
			break
		}
		if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() == "noteam" {
			// 空きマス
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
		} else if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != p.pieceTeam {
			// 敵駒マスも移動候補。でもそれ以上は進めない。
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			break
		} else {
			break
		}
	}

	return possibleNextMove
}
