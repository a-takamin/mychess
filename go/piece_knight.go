package main

type Knight struct {
	pieceType string
	pieceTeam string
	iconPath  string
}

func NewKnight(team string) *Knight {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bn.png"
	} else {
		iconPath = "./image/wn.png"
	}
	return &Knight{
		pieceType: "Knight",
		pieceTeam: team,
		iconPath:  iconPath,
	}
}

func (p *Knight) Move() int {
	return 1
}

func (p *Knight) GetPieceType() string {
	return p.pieceType
}

func (p *Knight) GetPieceTeam() string {
	return p.pieceTeam
}

func (p *Knight) CalcNextPosition(currentPosition int) []int {
	return []int{1}
}

func (p *Knight) GetIconPath() string {
	return p.iconPath
}
