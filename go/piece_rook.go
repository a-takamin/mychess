package main

type Rook struct {
	pieceType string
	pieceTeam string
	iconPath  string
}

func NewRook(team string) *Rook {
	var iconPath string
	if team == "black" {
		iconPath = "./image/br.png"
	} else {
		iconPath = "./image/wr.png"
	}
	return &Rook{
		pieceType: "Rook",
		pieceTeam: team,
		iconPath:  iconPath,
	}
}

func (p *Rook) Move() int {
	return 1
}

func (p *Rook) GetPieceType() string {
	return p.pieceType
}

func (p *Rook) GetPieceTeam() string {
	return p.pieceTeam
}

func (p *Rook) CalcNextPosition(currentPosition int) []int {
	return []int{1}
}

func (p *Rook) GetIconPath() string {
	return p.iconPath
}
