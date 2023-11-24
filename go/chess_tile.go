package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type ChessTile struct {
	widget.BaseWidget // 「埋め込み」と呼ばれる。widget.BaseWidget型のフィールドとメソッドを継承する
	tileId            int
	team              string
	backgroundColor   color.Color
	backgroundImage   *canvas.Image
}

// func NewChessTile(position int, bgColor color.Color, bgImage *canvas.Image) *ChessTile {
// 	w := &ChessTile{
// 		tileId:          position,
// 		team:            "",
// 		backgroundColor: bgColor,
// 		backgroundImage: bgImage,
// 	}
// 	w.ExtendBaseWidget(w)
// 	return w
// }

func NewChessTile(position int, piece Piece) *ChessTile {

	tile := &ChessTile{
		tileId:          position,
		team:            piece.GetPieceTeam(),
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

func (w *ChessTile) Tapped(*fyne.PointEvent) {
	fmt.Println("Tapped")
}

// @Override
func (w *ChessTile) MinSize() fyne.Size {
	if w.backgroundImage == nil {
		return fyne.NewSize(1, 1) // 暫定的…
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
	// r.background.FillColor = r.background.FillColor
	r.background.Refresh()
	if r.image != nil {
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
