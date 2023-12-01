package main

type Pawn struct {
	BasePiece
}

const (
	// Pawn can move 1 step forward
	PAWN_MOVE_STEP = 8 // 8*8マスなので1マス前進すると座標は8動く
)

func NewPawn(team string) *Pawn {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bp.png"
	} else {
		iconPath = "./image/wp.png"
	}
	return &Pawn{
		BasePiece: BasePiece{
			pieceType: "Pawn",
			pieceTeam: team,
			iconPath:  iconPath,
		},
	}
}

func (p *Pawn) CalcPossibleNextMove(currentTile *ChessTile) []*Move {
	possibleNextMove := []*Move{}
	pos := currentTile.getTileId()
	direction := 1 // 左上から右下に向かって0 -> 63に向かっていくので、黒はポーンが進むに連れ座標の値が大きくなる
	if p.pieceTeam == "white" {
		direction = -1 // 白はポーンが進むに連れ座標の値が小さくなる
	}

	// 初期位置からは2マス
	if isFirstMove(pos) {
		if (currentTile.chessBoard.getChessTile(pos+(PAWN_MOVE_STEP*direction)).getPiece().GetPieceTeam() == "noteam") && (currentTile.chessBoard.getChessTile(pos+(PAWN_MOVE_STEP*direction*2)).getPiece().GetPieceTeam() == "noteam") {
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, pos+(PAWN_MOVE_STEP*direction)*2))
		}
	}
	// 普通に1マス
	if currentTile.chessBoard.getChessTile(pos+(PAWN_MOVE_STEP*direction)).getPiece().GetPieceTeam() == "noteam" {
		possibleNextMove = append(possibleNextMove, NewMove(currentTile, pos+(PAWN_MOVE_STEP*direction)))
	}
	// 斜め前方のマスに敵コマがいれば、そこも候補に入れる
	possibleNextMove = append(possibleNextMove, calcAttackMove(currentTile)...)
	return possibleNextMove
}

func calcAttackMove(currentTile *ChessTile) []*Move {
	move := []*Move{}
	team := currentTile.getPiece().GetPieceTeam()
	pos := currentTile.getTileId()
	if team == "white" {
		leftAttack := pos - 9
		rightAttack := pos - 7
		if !isFirstColumn(pos) && leftAttack >= 0 && currentTile.getChessBoard().getChessTile(leftAttack).getPiece().GetPieceTeam() == "black" {
			move = append(move, NewMove(currentTile, leftAttack))
		}
		if !isEighthColumn(pos) && rightAttack >= 0 && currentTile.getChessBoard().getChessTile(rightAttack).getPiece().GetPieceTeam() == "black" {
			move = append(move, NewMove(currentTile, rightAttack))
		}
	}
	if team == "black" {
		leftAttack := pos + 9
		rightAttack := pos + 7
		if !isEighthColumn(pos) && leftAttack <= 63 && currentTile.getChessBoard().getChessTile(leftAttack).getPiece().GetPieceTeam() == "white" {
			move = append(move, NewMove(currentTile, leftAttack))
		}
		if !isFirstColumn(pos) && rightAttack <= 63 && currentTile.getChessBoard().getChessTile(rightAttack).getPiece().GetPieceTeam() == "white" {
			move = append(move, NewMove(currentTile, rightAttack))
		}
	}
	return move
}

func isFirstColumn(currentPos int) bool {
	firstColumn := [8]int{0, 8, 16, 24, 32, 40, 48, 56}
	for _, s := range firstColumn {
		if currentPos == s {
			return true
		}
	}
	return false
}

func isEighthColumn(currentPos int) bool {
	eighthColumn := [8]int{7, 15, 23, 31, 39, 47, 55, 63}
	for _, s := range eighthColumn {
		if currentPos == s {
			return true
		}
	}
	return false
}

func isFirstMove(currentPos int) bool {
	firstSquare := [16]int{8, 9, 10, 11, 12, 13, 14, 15, 48, 49, 50, 51, 52, 53, 54, 55}
	for _, s := range firstSquare {
		if currentPos == s {
			return true
		}
	}
	return false
}

func isEndSquare(currentPos int) bool {
	endSquare := [16]int{0, 1, 2, 3, 4, 5, 6, 7, 56, 57, 58, 59, 60, 61, 62, 63}
	for _, s := range endSquare {
		if currentPos == s {
			return true
		}
	}
	return false
}
