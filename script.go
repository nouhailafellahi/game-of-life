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
	

	grid := container.NewWithoutLayout()

	//draw grid lines
	for i := 0; i < 51; i++ {
		//draw vertical lines
		rect := canvas.NewRectangle(color.RGBA{R:255, G:255, B:255, A:100})

		rect.Resize(fyne.NewSize(1,500))
		rect.FillColor = color.RGBA{R:255, G:255, B:255, A:100};
		rect.StrokeColor = color.RGBA{R:255, G:255, B:255, A:100}

		grid.Add(rect)
		rect.Move(fyne.NewPos(float32(20*i), 0))

		//add horizontal lines
		if (i < 26) {
			rect := canvas.NewRectangle(color.RGBA{R:255, G:255, B:255, A:100})

			rect.Resize(fyne.NewSize(1000,1))
			rect.FillColor = color.RGBA{R:255, G:255, B:255, A:100}
			rect.StrokeColor = color.RGBA{R:255, G:255, B:255, A:100}

			grid.Add(rect)
			rect.Move(fyne.NewPos(0, float32(20*i)))
		}
	}


	grid.Resize(fyne.NewSize(1000,500))

	content := container.NewWithoutLayout(grid)
	grid.Move(fyne.NewPos(100,0))

	w.SetContent(content)

	w.Resize(fyne.NewSize(1200,670))
	w.ShowAndRun()	

}
