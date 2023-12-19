package main

type Knight struct {
	BasePiece
}

func NewKnight(team string) *Knight {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bn.png"
	} else {
		iconPath = "./image/wn.png"
	}
	return &Knight{
		BasePiece: BasePiece{
			pieceType: "Knight",
			pieceTeam: team,
			iconPath:  iconPath,
		},
	}
}

func (p *Knight) CalcPossibleNextMove(currentTile *ChessTile) []*Move {
	possibleNextMove := []*Move{}
	pos := currentTile.getTileId()
	var nextMoveCalcList []int

	if isFirstColumn(pos) {
		nextMoveCalcList = []int{-15, -6, 10, 17}
	} else if isSecondColumn(pos) {
		nextMoveCalcList = []int{-17, -15, -6, 10, 15, 17}
	} else if isSeventhColumn(pos) {
		nextMoveCalcList = []int{-17, -15, -10, 6, 15, 17}
	} else if isEighthColumn(pos) {
		nextMoveCalcList = []int{-17, -10, 6, 15}
	} else {
		nextMoveCalcList = []int{-17, -15, -10, -6, 6, 10, 15, 17}
	}

	for _, n := range nextMoveCalcList {
		possiblePos := pos + n
		if !isValidPos(possiblePos) {
			continue
		}
		if currentTile.chessBoard.getChessTile(possiblePos).getPiece().GetPieceTeam() != currentTile.team {
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, possiblePos))
		}
	}
	return possibleNextMove
}
