package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"image/color"
	"fmt"
)

func main() {
	a := app.New()
	w := a.NewWindow("Conway's Game of Life")
	

	grid := container.NewWithoutLayout()

	//draw horizontal lines
	for i := 0; i < 20; i++ {
		rect := canvas.NewRectangle(color.White)

		rect.Resize(fyne.NewSize(1000,1))
		rect.FillColor = color.White;
		rect.StrokeColor = color.White

		grid.Add(rect)
		rect.Move(fyne.NewPos(0, float32(30*i)))
	}
	
	grid.Resize(fyne.NewSize(1000,600))

	content := container.NewWithoutLayout(grid)
	grid.Move(fyne.NewPos(100,0))

	w.SetContent(content)

	w.Resize(fyne.NewSize(1200,670))
	w.ShowAndRun()	

}
