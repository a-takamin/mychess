package main

type Queen struct {
	BasePiece
}

func NewQueen(team string) *Queen {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bq.png"
	} else {
		iconPath = "./image/wq.png"
	}
	return &Queen{
		BasePiece: BasePiece{
			pieceType: "Queen",
			pieceTeam: team,
			iconPath:  iconPath,
		},
	}
}
