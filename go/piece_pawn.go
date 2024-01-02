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
	if p.isFirstMove(pos) {
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
		leftAttackTileId := pos - 9
		rightAttackTileId := pos - 7
		if !isFirstColumn(pos) && leftAttackTileId >= 0 && currentTile.getChessBoard().getChessTile(leftAttackTileId).getPiece().GetPieceTeam() == "black" {
			move = append(move, NewMove(currentTile, leftAttackTileId))
		}
		if !isEighthColumn(pos) && rightAttackTileId >= 0 && currentTile.getChessBoard().getChessTile(rightAttackTileId).getPiece().GetPieceTeam() == "black" {
			move = append(move, NewMove(currentTile, rightAttackTileId))
		}
		if currentTile.chessBoard.enPassantTarget != nil && currentTile.chessBoard.enPassantTarget.getPiece().GetPieceTeam() == "black" {
			// 相手ポーンが今いるマスの1マス後ろがアンパッサン攻撃対象マス
			if currentTile.chessBoard.enPassantTarget.getTileId()-8 == leftAttackTileId {
				move = append(move, NewMove(currentTile, leftAttackTileId))
			}
			if currentTile.chessBoard.enPassantTarget.getTileId()-8 == rightAttackTileId {
				move = append(move, NewMove(currentTile, rightAttackTileId))
			}
		}
	}
	if team == "black" {
		leftAttackTileId := pos + 9
		rightAttackTileId := pos + 7
		if !isEighthColumn(pos) && leftAttackTileId <= 63 && currentTile.getChessBoard().getChessTile(leftAttackTileId).getPiece().GetPieceTeam() == "white" {
			move = append(move, NewMove(currentTile, leftAttackTileId))
		}
		if !isFirstColumn(pos) && rightAttackTileId <= 63 && currentTile.getChessBoard().getChessTile(rightAttackTileId).getPiece().GetPieceTeam() == "white" {
			move = append(move, NewMove(currentTile, rightAttackTileId))
		}
		if currentTile.chessBoard.enPassantTarget != nil && currentTile.chessBoard.enPassantTarget.getPiece().GetPieceTeam() == "white" {
			// 相手ポーンが今いるマスの1マス後ろがアンパッサン攻撃対象マス
			if currentTile.chessBoard.enPassantTarget.getTileId()+8 == leftAttackTileId {
				move = append(move, NewMove(currentTile, leftAttackTileId))
			}
			if currentTile.chessBoard.enPassantTarget.getTileId()+8 == rightAttackTileId {
				move = append(move, NewMove(currentTile, rightAttackTileId))
			}
		}
	}
	return move
}

func (p *Pawn) isFirstMove(currentPos int) bool {
	var firstSquare [8]int
	if p.pieceTeam == "white" {
		firstSquare = [8]int{48, 49, 50, 51, 52, 53, 54, 55}
	} else if p.pieceTeam == "black" {
		firstSquare = [8]int{8, 9, 10, 11, 12, 13, 14, 15}
	}
	for _, s := range firstSquare {
		if currentPos == s {
			return true
		}
	}
	return false
}
