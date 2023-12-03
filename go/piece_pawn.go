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
