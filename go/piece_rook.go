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
			if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			}
		}
	}
	// 下
	if !isEighthRank(pos) {
		for i := pos + downMove; isValidPos(i); i += downMove {
			if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			}
		}
	}
	// 左
	if !isFirstColumn(pos) {
		// 第8カラムにワープするまで
		for i := pos + leftMove; isValidPos(i) && !isEighthColumn(i); i += leftMove {
			if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			}
		}
	}
	// 右
	if !isEighthColumn(pos) {
		// 第1カラムにワープするまで
		for i := pos + rightMove; isValidPos(i) && !isFirstColumn(i); i += rightMove {
			if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			}
		}
	}
	return possibleNextMove
}
