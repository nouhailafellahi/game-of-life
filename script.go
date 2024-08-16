package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"image/color"
	"log"
	"math"
)

//array of squares
//value of 0 if empty
//valye of 1 if colored
var squares [50][25]*canvas.Rectangle
var fill[50][25]int

//Struct and declaration of a clickable rectangle. Used as the background of grid
type clickableRectangle struct {
	widget.BaseWidget
	rect    *canvas.Rectangle
	OnTapped func(fyne.Position)
}

func newClickableRectangle() *clickableRectangle {
	r := &clickableRectangle{
		rect: canvas.NewRectangle(color.White),
	}
	r.ExtendBaseWidget(r)
	return r
}

func (r *clickableRectangle) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(r.rect)
}

func (r *clickableRectangle) Tapped(event *fyne.PointEvent) {
	if r.OnTapped != nil {
		r.OnTapped(event.Position)
	}
}






//Methods
func play(grid *fyne.Container) {

	
}

func stop(grid *fyne.Container) {

}

func reset(grid *fyne.Container) {
	stop(grid)

}

func click(grid *fyne.Container, pos fyne.Position){
	xPos := (int(math.Floor(float64(int(pos.X))/20)))
	yPos := (int(math.Floor(float64(int(pos.Y))/20)))
	
	println("Position clicked: ", xPos, yPos)//debug

	if(fill[xPos][yPos] == 0) {
		squares[xPos][yPos].FillColor = color.White
		fill[xPos][yPos] = 1
	} else {
		squares[xPos][yPos].FillColor = color.Transparent
		fill[xPos][yPos] = 0
	}
	grid.Refresh()
		
}



//Main
func main() {
	println("Start build")//debug
	a := app.New()
	w := a.NewWindow("Conway's Game of Life")
	
	//contains grid lines and squares
	grid := container.NewWithoutLayout()

	//initialize array of squares
	for i:=0;i<len(squares);i++ {
		for j:=0;j<len(squares[i]);j++{
			rect := canvas.NewRectangle(color.Transparent)
			rect.Resize(fyne.NewSize(20,20))
			rect.Move(fyne.NewPos(float32(i*20), float32(j*20)))
			squares[i][j] = rect
			grid.Add(squares[i][j])
		}
	}

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

	

	//Properly size grid 
	grid.Resize(fyne.NewSize(1000,500))

	//Create a clickable, transparent screen behind the grid
	screen := newClickableRectangle()
	screen.OnTapped = func(pos fyne.Position) {
		click(grid, pos)
	}
	screen.Resize(fyne.NewSize(1000,500))
	screen.rect.FillColor = color.RGBA{R:100,G:100,B:100,A:30}
	
	//Add grid and screen to the container and adjust positions
	content := container.NewWithoutLayout(screen, grid)
	grid.Move(fyne.NewPos(100,40))
	screen.Move(fyne.NewPos(100,40))


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
		stop(grid)
	} )
	stopBtn.Move(fyne.NewPos(525, 580))
	stopBtn.Resize(fyne.NewSize(150, 50))
	//Add button to window
	content.Add(stopBtn)



	//Create and position Reset button
	resetBtn := widget.NewButton("Reset", func(){
		reset(grid)
	} )
	resetBtn.Move(fyne.NewPos(750, 580))
	resetBtn.Resize(fyne.NewSize(150, 50))
	//Add button to window
	content.Add(resetBtn)
	
	w.SetContent(content)
	w.Resize(fyne.NewSize(1200,670))

	//Use customized window icon
	icon , err := fyne.LoadResourceFromPath("./cat.png")
	if err != nil {
		log.Fatal(err)
	}
	w.SetIcon(icon)




	go func() {
		//run while show and running
	}()

	w.ShowAndRun()	

}