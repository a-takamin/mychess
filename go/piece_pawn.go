package main

type Pawn struct {
	BasePiece
}

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

func (p *Pawn) CalcPossibleNextMove(currentPosition int) []*Move {
	return []*Move{}
}
