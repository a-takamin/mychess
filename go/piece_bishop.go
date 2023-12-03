package main

type Bishop struct {
	BasePiece
}

func NewBishop(team string) *Bishop {
	var iconPath string
	if team == "black" {
		iconPath = "./image/bb.png"
	} else {
		iconPath = "./image/wb.png"
	}
	return &Bishop{
		BasePiece: BasePiece{
			pieceType: "Bishop",
			pieceTeam: team,
			iconPath:  iconPath,
		},
	}
}

func (p *Bishop) CalcPossibleNextMove(currentTile *ChessTile) []*Move {
	possibleNextMove := []*Move{}
	pos := currentTile.getTileId()
	leftUpMove := -9
	rightUpMove := -7
	leftDownMove := 7
	rightDownMove := 9
	// 左上
	for i := pos + leftUpMove; i >= 0; i += leftUpMove {
		if isEighthColumn(i) {
			// 動いた結果8列目にいる = 左上Moveとしてはありえない（もともと左端の1列目にいた）
			break
		}
		if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() == "noteam" {
			// 空きマス
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
		} else if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != p.pieceTeam {
			// 敵駒マスも移動候補。でもそれ以上は進めない。
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			break
		} else {
			break
		}
	}
	// 右上
	for i := pos + rightUpMove; i >= 0; i += rightUpMove {
		if isFirstColumn(i) {
			// 動いた結果1列目にいる = 右上Moveとしてはありえない（もともと右端の8列目にいた）
			break
		}
		if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() == "noteam" {
			// 空きマス
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
		} else if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != p.pieceTeam {
			// 敵駒マスも移動候補。でもそれ以上は進めない。
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			break
		} else {
			break
		}
	}
	// 左下
	for i := pos + leftDownMove; i <= 63; i += leftDownMove {
		if isEighthColumn(i) {
			// 動いた結果8列目にいる = 左下Moveとしてはありえない（もともと左端の1列目にいた）
			break
		}
		if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() == "noteam" {
			// 空きマス
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
		} else if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != p.pieceTeam {
			// 敵駒マスも移動候補。でもそれ以上は進めない。
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			break
		} else {
			break
		}
	}
	// 右下
	for i := pos + rightDownMove; i <= 63; i += rightDownMove {
		if isFirstColumn(i) {
			// 動いた結果1列目にいる = 右下Moveとしてはありえない（もともと右端の8列目にいた）
			break
		}
		if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() == "noteam" {
			// 空きマス
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
		} else if currentTile.chessBoard.getChessTile(i).getPiece().GetPieceTeam() != p.pieceTeam {
			// 敵駒マスも移動候補。でもそれ以上は進めない。
			possibleNextMove = append(possibleNextMove, NewMove(currentTile, i))
			break
		} else {
			break
		}
	}
	return possibleNextMove
}
