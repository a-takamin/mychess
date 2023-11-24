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
