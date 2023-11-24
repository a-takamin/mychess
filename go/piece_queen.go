package main

type Queen struct {
	pieceType string
	pieceTeam string
	iconPath  string
}

func NewQueen(team string) *Queen {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bq.png"
	} else {
		iconPath = "./image/wq.png"
	}
	return &Queen{
		pieceType: "Queen",
		pieceTeam: team,
		iconPath:  iconPath,
	}
}

func (p *Queen) Move() int {
	return 1
}

func (p *Queen) GetPieceType() string {
	return p.pieceType
}

func (p *Queen) GetPieceTeam() string {
	return p.pieceTeam
}

func (p *Queen) CalcNextPosition(currentPosition int) []int {
	return []int{1}
}

func (p *Queen) GetIconPath() string {
	return p.iconPath
}
