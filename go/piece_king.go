package main

type King struct {
	pieceType string
	pieceTeam string
	iconPath  string
}

func NewKing(team string) *King {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bk.png"
	} else {
		iconPath = "./image/wk.png"
	}
	return &King{
		pieceType: "King",
		pieceTeam: team,
		iconPath:  iconPath,
	}
}

func (p *King) Move() int {
	return 1
}

func (p *King) GetPieceType() string {
	return p.pieceType
}

func (p *King) GetPieceTeam() string {
	return p.pieceTeam
}

func (p *King) CalcNextPosition(currentPosition int) []int {
	return []int{1}
}

func (p *King) GetIconPath() string {
	return p.iconPath
}
