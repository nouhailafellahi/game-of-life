package main 

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


func main() {
	a := app.New()
	w := a.NewWindow("Conway's Game of Life")


	w.setContent(widget.NewLabel("Hello..."))
	w.ShowAndRun()

}
