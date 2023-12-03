package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("takamin chess by golang")
	chessBoard := NewInitialChessBoard()
	chessBoard.w = w
	w.SetContent(chessBoard.GetBoard())
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
