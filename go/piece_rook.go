package main

type Rook struct {
	BasePiece
}

func NewRook(team string) *Rook {
	var iconPath string
	if team == "black" {
		iconPath = "./image/br.png"
	} else {
		iconPath = "./image/wr.png"
	}
	return &Rook{
		BasePiece: BasePiece{
			pieceType: "Rook",
			pieceTeam: team,
			iconPath:  iconPath,
		},
	}
}

func (p *Rook) CalcPossibleNextMove(currentTile *ChessTile) []*Move {
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
	return possibleNextMove
}
