package main

type BasePiece struct {
	pieceType string
	pieceTeam string
	iconPath  string
}

// 基本的なメソッドを定義する。Pieceインターフェースを満たすわけではない。
func (p *BasePiece) GetPieceType() string {
	return p.pieceType
}

func (p *BasePiece) GetPieceTeam() string {
	return p.pieceTeam
}

func (p *BasePiece) GetIconPath() string {
	return p.iconPath
}
