package main

type NoPiece struct {
	BasePiece
}

func NewNoPiece(team string) *NoPiece {
	return &NoPiece{
		BasePiece: BasePiece{
			pieceType: "NoPiece",
			pieceTeam: team,
			iconPath:  "",
		},
	}
}

func (p *NoPiece) CalcPossibleNextMove(currentPosition int) []*Move {
	return []*Move{}
}
