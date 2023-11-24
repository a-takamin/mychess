package main

type Bishop struct {
	BasePiece
}

func NewBishop(team string) *Bishop {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bb.png"
	} else {
		iconPath = "./image/wb.png"
	}
	return &Bishop{
		BasePiece: BasePiece{
			pieceType: "Bishop",
			pieceTeam: team,
			iconPath:  iconPath,
		},
	}
}
