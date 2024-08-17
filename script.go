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
	"time"
)

//array of squares
//value of 0 if empty
//valye of 1 if colored
var squares [50][25]*canvas.Rectangle
var fill[50][25]int
var neighbors[50][25]int
var playing bool


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
	if !playing {
		playing = true
		for playing {
			countNeighbors()
			for i:=0;i<len(squares);i++ {
				for j:=0;j<len(squares[i]);j++{
	
					//decide wether to keep current status or click()
					if(fill[i][j] == 0) {
						//if the cell is dead and has exactly 3 alive neighbors, click() it
						if(neighbors[i][j] == 3) {
							//println("dead, will click")//debug
							click(grid,i,j)
						}
					} else if (fill[i][j] == 1) {
						//if the cell is alive and has neither 2 nor 3 alive neighbors, click() it
						if(neighbors[i][j] != 2 && neighbors[i][j] != 3) {
							//println("alive, will click")//debug
							click(grid,i,j)
						}
					}
				}
			}
			time.Sleep(2*time.Second)
		}

	}
	
		
	
}

func stop(grid *fyne.Container) {
	playing = false
}

func reset(grid *fyne.Container) {
	stop(grid)

}

func initNeighbors() {
	
	//initialize array of neighbors
	for i:=0;i<len(neighbors);i++{
		for j:=0;j<len(neighbors[i]);j++{
			neighbors[i][j] = 0
		}
	}

}

func countNeighbors() {
	initNeighbors()
	
	//following loop counts the neighbors of each cell
	//then decides wether to click() the cell or not

	//run through every cell of squares[]
	for i:=0;i<len(squares);i++ {
		for j:=0;j<len(squares[i]);j++{

			//count the neighbors for squares[i,j]
			diff := [3]int{-1,0,1}
			for a:=0;a<len(diff);a++{
				for b:=0;b<len(diff);b++{
					
					//index (n,m) of neighbor
					n := i+diff[a]
					m := j+diff[b]

					//if statement to check indexes are in-bound and indexes != [i,j]
					if(n>=0 && n<=49 && m>=0 && m<=24 && ( n!=i || m!=j)) {
						if(fill[n][m] == 1) {
							neighbors[i][j]++
						}
					}
				}
			}
			//println(neighbors[i][j])//debug
		}
	}
			
			
}

func click(grid *fyne.Container, xPos int, yPos int){
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
	playing = true
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

	initNeighbors()

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
		xPos := (int(math.Floor(float64(int(pos.X))/20)))
		yPos := (int(math.Floor(float64(int(pos.Y))/20)))
		
		click(grid, xPos, yPos)
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
