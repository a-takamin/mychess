package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Grid Layout")
	chessBoard := NewInitialChessBoard()

	w.SetContent(chessBoard.GetBoard())
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()

}
