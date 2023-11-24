package main

type ChessBoardBuilder struct {
	predefinedChessTileMap map[int]*ChessTile
}

func NewChessBoardBuilder() *ChessBoardBuilder {
	return &ChessBoardBuilder{
		// nil map になってしまうので、make() で初期化する
		predefinedChessTileMap: make(map[int]*ChessTile),
	}
}

func (b *ChessBoardBuilder) Build() *ChessBoard {
	return NewChessBoard(b.predefinedChessTileMap)
}

func (b *ChessBoardBuilder) Set(tile *ChessTile) *ChessBoardBuilder {
	b.predefinedChessTileMap[tile.tileId] = tile
	return b
}
