package main

type Pawn struct {
	pieceType string
	pieceTeam string
	iconPath  string
}

func NewPawn(team string) *Pawn {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bp.png"
	} else {
		iconPath = "./image/wp.png"
	}
	return &Pawn{
		pieceType: "Pawn",
		pieceTeam: team,
		iconPath:  iconPath,
	}
}

func (p *Pawn) Move() int {
	return 1
}

func (p *Pawn) GetPieceType() string {
	return p.pieceType
}

func (p *Pawn) GetPieceTeam() string {
	return p.pieceTeam
}

func (p *Pawn) CalcNextPosition(currentPosition int) []int {
	return []int{1}
}

func (p *Pawn) GetIconPath() string {
	return p.iconPath
}
