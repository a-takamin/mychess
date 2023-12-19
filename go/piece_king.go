package main

type King struct {
	BasePiece
}

func NewKing(team string) *King {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bk.png"
	} else {
		iconPath = "./image/wk.png"
	}
	return &King{
		BasePiece: BasePiece{
			pieceType: "King",
			pieceTeam: team,
			iconPath:  iconPath,
		},
	}
}

func (p *King) CalcPossibleNextMove(currentTile *ChessTile) []*Move {
	possibleNextMove := []*Move{}
	pos := currentTile.getTileId()
	leftMove := -1
	rightMove := 1
	upMove := -8
	downMove := 8
	leftUpMove := -9
	rightUpMove := -7
	rightDownMove := 9
	leftDownMove := 7

	// 上
	if !isFirstRank(pos) {
		nextPos := pos + upMove
		if isValidPos(nextPos) {
			possibleNextMovePosTeam := currentTile.getChessBoard().getChessTile(nextPos).getPiece().GetPieceTeam()
			if possibleNextMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, nextPos))
			}
		}
	}
	// 下
	if !isEighthRank(pos) {
		nextPos := pos + downMove
		if isValidPos(nextPos) {
			possibleNextMovePosTeam := currentTile.getChessBoard().getChessTile(nextPos).getPiece().GetPieceTeam()
			if possibleNextMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, nextPos))
			}
		}
	}
	// 左
	if !isFirstColumn(pos) {
		nextPos := pos + leftMove
		if isValidPos(nextPos) && !isEighthColumn(nextPos) {
			possibleNextMovePosTeam := currentTile.getChessBoard().getChessTile(nextPos).getPiece().GetPieceTeam()
			if possibleNextMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, nextPos))
			}
		}
	}
	// 右
	if !isEighthColumn(pos) {
		nextPos := pos + rightMove
		if isValidPos(nextPos) && !isFirstColumn(nextPos) {
			possibleNextMovePosTeam := currentTile.getChessBoard().getChessTile(nextPos).getPiece().GetPieceTeam()
			if possibleNextMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, nextPos))
			}
		}
	}
	// 左上
	if !isFirstColumn(pos) {
		nextPos := pos + leftUpMove
		if isValidPos(nextPos) {
			possibleNextMovePosTeam := currentTile.getChessBoard().getChessTile(nextPos).getPiece().GetPieceTeam()
			if possibleNextMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, nextPos))
			}
		}
	}
	// 右上
	if !isEighthColumn(pos) {
		nextPos := pos + rightUpMove
		if isValidPos(nextPos) {
			possibleNextMovePosTeam := currentTile.getChessBoard().getChessTile(nextPos).getPiece().GetPieceTeam()
			if possibleNextMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, nextPos))
			}
		}
	}
	// 右下
	if !isEighthColumn(pos) {
		nextPos := pos + rightDownMove
		if isValidPos(nextPos) {
			possibleNextMovePosTeam := currentTile.getChessBoard().getChessTile(nextPos).getPiece().GetPieceTeam()
			if possibleNextMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, nextPos))
			}
		}
	}
	// 左下
	if !isFirstColumn(pos) {
		nextPos := pos + leftDownMove
		if isValidPos(nextPos) {
			possibleNextMovePosTeam := currentTile.getChessBoard().getChessTile(nextPos).getPiece().GetPieceTeam()
			if possibleNextMovePosTeam != currentTile.team {
				possibleNextMove = append(possibleNextMove, NewMove(currentTile, nextPos))
			}
		}
	}
	return possibleNextMove
}
