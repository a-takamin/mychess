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

func isFirstColumn(currentPos int) bool {
	firstColumn := [8]int{0, 8, 16, 24, 32, 40, 48, 56}
	for _, s := range firstColumn {
		if currentPos == s {
			return true
		}
	}
	return false
}

func isEighthColumn(currentPos int) bool {
	eighthColumn := [8]int{7, 15, 23, 31, 39, 47, 55, 63}
	for _, s := range eighthColumn {
		if currentPos == s {
			return true
		}
	}
	return false
}

func isFirstMove(currentPos int) bool {
	firstSquare := [16]int{8, 9, 10, 11, 12, 13, 14, 15, 48, 49, 50, 51, 52, 53, 54, 55}
	for _, s := range firstSquare {
		if currentPos == s {
			return true
		}
	}
	return false
}

// func isEndSquare(currentPos int) bool {
// 	endSquare := [16]int{0, 1, 2, 3, 4, 5, 6, 7, 56, 57, 58, 59, 60, 61, 62, 63}
// 	for _, s := range endSquare {
// 		if currentPos == s {
// 			return true
// 		}
// 	}
// 	return false
// }
