package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"image/color"
	"fmt"
)

//array of squares
//value of 0 if empty
//valye of 1 if colored
var squares [25][50]int


func play(grid *fyne.Container) {
	for i:=0; i<len(squares); i++{
		for j:=0; j<len(squares[i]); j++ {
			rect := canvas.NewRectangle(color.White)
			rect.Resize(fyne.NewSize(20,20))
			rect.Move(fyne.NewPos(float32(j*20), float32(i*20)))
			grid.Add(rect)	
		}	
	}
	
	
}

func main() {
	fmt.Println("Start build")//debug
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
	

	//Properly size+position grid and add to window
	grid.Resize(fyne.NewSize(1000,500))
	content := container.NewWithoutLayout(grid)
	grid.Move(fyne.NewPos(100,50))


	//Create and position Play button
	playBtn := widget.NewButton("Play", func(){
		play(grid)
	})
	playBtn.Move(fyne.NewPos(300, 580))
	playBtn.Resize(fyne.NewSize(150, 50))
	//Add button to window
	content.Add(playBtn)


	//Create and position Stop button
	stopBtn := widget.NewButton("Stop", func(){
		stop()
	} )
	stopBtn.Move(fyne.NewPos(525, 580))
	stopBtn.Resize(fyne.NewSize(150, 50))
	//Add button to window
	content.Add(stopBtn)

	//Create and position Reset button
	resetBtn := widget.NewButton("Reset", func(){
		reset()
	} )
	resetBtn.Move(fyne.NewPos(750, 580))
	resetBtn.Resize(fyne.NewSize(150, 50))
	//Add button to window
	content.Add(resetBtn)
	

	w.SetContent(content)
	w.Resize(fyne.NewSize(1200,670))

	go func() {
		//run while show and running
	}()

	w.ShowAndRun()	

}

func stop() {

}

func reset() {
	stop()


}