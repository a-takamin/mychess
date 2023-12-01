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
	return []*Move{}
}
