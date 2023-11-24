package main

type NoPiece struct {
	pieceType string
	pieceTeam string
	iconPath  string
}

func NewNoPiece() *NoPiece {
	return &NoPiece{
		pieceType: "NoPiece",
		pieceTeam: "NoTeam",
		iconPath:  "NoIconPath",
	}
}

func (p *NoPiece) Move() int {
	return 1
}

func (p *NoPiece) GetPieceType() string {
	return p.pieceType
}

func (p *NoPiece) GetPieceTeam() string {
	return p.pieceTeam
}

func (p *NoPiece) CalcNextPosition(currentPosition int) []int {
	return []int{1}
}

func (p *NoPiece) GetIconPath() string {
	return p.iconPath
}
