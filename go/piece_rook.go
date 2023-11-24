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
