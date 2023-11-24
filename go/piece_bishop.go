package main

type Bishop struct {
	pieceType string
	pieceTeam string
	iconPath  string
}

func NewBishop(team string) *Bishop {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bb.png"
	} else {
		iconPath = "./image/wb.png"
	}
	return &Bishop{
		pieceType: "Bishop",
		pieceTeam: team,
		iconPath:  iconPath,
	}
}

func (p *Bishop) Move() int {
	return 1
}

func (p *Bishop) GetPieceType() string {
	return p.pieceType
}

func (p *Bishop) GetPieceTeam() string {
	return p.pieceTeam
}

func (p *Bishop) CalcNextPosition(currentPosition int) []int {
	return []int{1}
}

func (p *Bishop) GetIconPath() string {
	return p.iconPath
}
