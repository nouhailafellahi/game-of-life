package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Conway's Game of Life")
	

	grid := container.NewGridWithColumns(22)

	for i := 0; i < 572; i++ {
		rect := canvas.NewRectangle(color.Black)
		rect.Resize(fyne.NewSize(20,20))
		rect.FillColor = color.White;
		rect.StrokeColor = color.White
		rect.StrokeWidth = 1
		grid.Add(rect)

	}
	grid.Resize(fyne.NewSize(500,600))

	content := container.NewWithoutLayout(grid)
	grid.Move(fyne.NewPos(100,0))

	w.SetContent(content)

	w.Resize(fyne.NewSize(700,670))
	w.ShowAndRun()	

}
