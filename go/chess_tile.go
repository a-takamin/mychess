package main

import (
	"fmt"
	"image/color"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type ChessTile struct {
	widget.BaseWidget // 「埋め込み」と呼ばれる。widget.BaseWidget型のフィールドとメソッドを継承する
	chessBoard        *ChessBoard
	tileId            int
	team              string
	// piece は interface として使ってポリモーフィズムを機能させるのでポインタではなくそのまま構造体を持つ
	piece           Piece
	backgroundColor color.Color
	backgroundImage *canvas.Image
}

func (w *ChessTile) updateImage(path string) {
	w.backgroundImage = canvas.NewImageFromFile(path)
	w.Refresh()
}

func NewChessTile(position int, piece Piece, chessBoard *ChessBoard) *ChessTile {

	tile := &ChessTile{
		chessBoard:      chessBoard,
		tileId:          position,
		team:            piece.GetPieceTeam(),
		piece:           piece,
		backgroundColor: calcTileColor(position),
		backgroundImage: canvas.NewImageFromFile(piece.GetIconPath()),
	}
	tile.ExtendBaseWidget(tile)
	return tile
}

func calcTileColor(position int) color.Color {
	row := position / 8
	col := position % 8
	if (row+col)%2 == 0 {
		return WhiteTileColor()
	}
	return GreenTileColor()
}

// @Override
func (w *ChessTile) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(w.backgroundColor)
	objects := []fyne.CanvasObject{bg}
	if w.backgroundImage != nil {
		objects = append(objects, w.backgroundImage)
	}
	return &chessTileRenderer{objects: objects, background: bg, image: w.backgroundImage}
}

func (tile *ChessTile) Tapped(*fyne.PointEvent) {
	if tile.chessBoard.selectedTile == nil {
		if tile.getPiece().GetPieceTeam() == "noteam" {
			fmt.Printf("NoTeam Tile is selected\n")
			tile.chessBoard.selectedTile = nil
			return
		}
		tile.chessBoard.selectedTile = tile
		fmt.Printf("%s's %s is selected!\n", tile.team, tile.piece.GetPieceType())
		return
	}
	// 以下、2回目のタップということになる
	firstSelectedTile := tile.chessBoard.selectedTile
	secondSelectedTile := tile
	move := NewMove(firstSelectedTile, secondSelectedTile.getTileId())

	if !isValidMove(move, tile.chessBoard.possibleNextMove) {
		fmt.Printf("Invalid move: %s's %s from %d to %d\n", firstSelectedTile.getPiece().GetPieceTeam(), firstSelectedTile.getPiece().GetPieceType(), firstSelectedTile.getTileId(), secondSelectedTile.getTileId())
		tile.chessBoard.selectedTile = nil
		return
	}
	// selectedTileの初期化
	tile.chessBoard.selectedTile = nil
	// ボードの更新（実際に手を実行する）
	renewChessBoard(tile.chessBoard, move)

}

func isValidMove(move *Move, possibleNextMove []*Move) bool {
	for _, m := range possibleNextMove {
		if move.fromTile.getTileId() == m.fromTile.getTileId() && move.to == m.to {
			return true
		}
	}
	return slices.Contains(possibleNextMove, move)
}

// @Override
func (w *ChessTile) MinSize() fyne.Size {
	if w.backgroundImage == nil {
		// 背景画像（コマ）がない場合は何を返せば良いのか。Gridで良い感じになるのでとりあえず(1,1)を返している
		return fyne.NewSize(1, 1)
	}
	return w.backgroundImage.MinSize()
}

type chessTileRenderer struct {
	objects    []fyne.CanvasObject
	background *canvas.Rectangle
	image      *canvas.Image
}

// @Override
func (r *chessTileRenderer) MinSize() fyne.Size {
	return r.background.MinSize()
}

// @Override
func (r *chessTileRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	if r.image != nil {
		r.image.Resize(size)
	}
}

// @Override
func (r *chessTileRenderer) Refresh() {
	r.background.Refresh()
	if r.image != nil {
		// なぜかこれをやらないとアンパッサンのときに画像が更新されない。
		// fyneを理解して後で直す。
		r.image = canvas.NewImageFromFile("")
		r.image.Refresh()
	}
}

// @Override
func (r *chessTileRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

// @Override
func (r *chessTileRenderer) Destroy() {}

func WhiteTileColor() color.RGBA {
	return color.RGBA{237, 241, 215, 255}
}

// 白と緑にした
func GreenTileColor() color.RGBA {
	return color.RGBA{146, 174, 120, 255}
}

func (t *ChessTile) getPiece() Piece {
	return t.piece
}

func (t *ChessTile) getTileId() int {
	return t.tileId
}

func (t *ChessTile) getChessBoard() *ChessBoard {
	return t.chessBoard
}
