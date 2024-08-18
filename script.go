package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"image/color"
	"log"
	"math"
	"time"
	"fmt"
)

//array of cells
var cells [50][25]*canvas.Rectangle

//value of 0 if empty
//valye of 1 if colored
var fill[50][25]int

//count of neighbors for each cell
var neighbors[50][25]int

//state variable indicating whether currently running game
var playing bool

//number of generations
var genCount int




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

//play() runs the game and updates the GUI 
func play(content *fyne.Container, generation *widget.Label) {
	playing = true
	genCount = 0

	for playing && !countNeighbors() {
		
		for i:=0;i<len(cells);i++ {
			for j:=0;j<len(cells[i]);j++{

				//decide wether to keep current status or click()
				if(fill[i][j] == 0) {
					//if the cell is dead and has exactly 3 alive neighbors, click() it
					if(neighbors[i][j] == 3) {
						click(content,i,j)
					}
				} else if (fill[i][j] == 1) {
					//if the cell is alive and has neither 2 nor 3 alive neighbors, click() it
					if(neighbors[i][j] != 2 && neighbors[i][j] != 3) {
						click(content,i,j)
					}
				}
			}
		}
		genCount++
		//update generation label
		updateGeneration(generation)

		time.Sleep(1*time.Second)
	}
	stop()
	
}

//Stop game
func stop() {
	playing = false
}

//Stop and reset game
func reset(grid *fyne.Container, generation *widget.Label) {
	stop()
	genCount = 0
	updateGeneration(generation)
	for i:=0;i<len(fill);i++{
		for j:=0;j<len(fill[i]);j++{
			if(fill[i][j] == 1) {
				//if the cell is alive, kill it.
				click(grid,i,j)
			}
		}
	}

}

//Update the Generation label on GUI
func updateGeneration(generation *widget.Label) {
	text := fmt.Sprintf("%s %d", "Generation ", genCount)
	generation.SetText(text)
}


//initialize array of neighbors
func initNeighbors() {
	for i:=0;i<len(neighbors);i++{
		for j:=0;j<len(neighbors[i]);j++{
			neighbors[i][j] = 0
		}
	}
}

//Count neighbors of each cell and store in neighbors[][]
func countNeighbors() bool {
	initNeighbors()
	lastGen := true;

	//run through every cell of cells[]
	for i:=0;i<len(cells);i++ {
		for j:=0;j<len(cells[i]);j++{

			//count the neighbors for cells[i,j]
			for a:=-1;a<2;a++{
				for b:=-1;b<2;b++{
					
					//index of neighbor (n,m)
					n := i+a
					m := j+b

					//if statement to check indexes are in-bound and indexes != [i,j]
					if(n>=0 && n<=49 && m>=0 && m<=24 && ( n!=i || m!=j)) {
						if(fill[n][m] == 1) {
							neighbors[i][j]++
						}
					}
				}
			}
			if neighbors[i][j] > 0 {
				lastGen = false
			}
		}
	}		
	return lastGen		
}

//if cell @ [i,j] is live, kill it and vice versa.
func click(grid *fyne.Container, xPos int, yPos int){
	if(fill[xPos][yPos] == 0) {
		cells[xPos][yPos].FillColor = color.White
		fill[xPos][yPos] = 1
	} else {
		cells[xPos][yPos].FillColor = color.Transparent
		fill[xPos][yPos] = 0
	}
	grid.Refresh()		
}


//Main/Driver
func main() {
	playing = false
	a := app.New()
	w := a.NewWindow("Conway's Game of Life")

	//initialize neighbors[][]
	initNeighbors()
	
	//contains grid lines and cells
	grid := container.NewWithoutLayout()

	//initialize array of cells
	for i:=0;i<len(cells);i++ {
		for j:=0;j<len(cells[i]);j++{
			rect := canvas.NewRectangle(color.Transparent)
			rect.Resize(fyne.NewSize(20,20))
			rect.Move(fyne.NewPos(float32(i*20), float32(j*20)))
			cells[i][j] = rect
			grid.Add(cells[i][j])
		}
	}

	//Create and draw grid lines
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

	//Create a clickable, transparent screen behind the grid
	screen := newClickableRectangle()
	screen.OnTapped = func(pos fyne.Position) {
		xPos := (int(math.Floor(float64(int(pos.X))/20)))
		yPos := (int(math.Floor(float64(int(pos.Y))/20)))
		
		click(grid, xPos, yPos)
	}
	screen.Resize(fyne.NewSize(1000,500))
	screen.rect.FillColor = color.RGBA{R:100,G:100,B:100,A:30}
	

	//Create Generation count label
	generation := widget.NewLabel("Generation 0")
	generation.Move(fyne.NewPos(980,0))


	//Add grid, screen and generation to the container and adjust positions
	content := container.NewWithoutLayout(screen, grid, generation)
	grid.Move(fyne.NewPos(100,40))
	screen.Move(fyne.NewPos(100,40))

	//Create and position Play button
	playBtn := widget.NewButton("Play", func(){
		//run independent from main()
		go func() {
			if !playing {
				play(content, generation)
			}
		}()
		
	})
	playBtn.Move(fyne.NewPos(300, 580))
	playBtn.Resize(fyne.NewSize(150, 50))
	//Add button to window
	content.Add(playBtn)


	//Create and position Stop button
	stopBtn := widget.NewButton("Stop", func(){
		//run independent of main()
		go func() {
			stop()
		}()
		
	} )
	stopBtn.Move(fyne.NewPos(525, 580))
	stopBtn.Resize(fyne.NewSize(150, 50))
	//Add button to window
	content.Add(stopBtn)


	//Create and position Reset button
	resetBtn := widget.NewButton("Reset", func(){
		//run independent from main()
		go func(){
			reset(grid, generation)
		}()
	} )
	resetBtn.Move(fyne.NewPos(750, 580))
	resetBtn.Resize(fyne.NewSize(150, 50))
	//Add button to window
	content.Add(resetBtn)
	
			
	//Add content to window
	w.SetContent(content)
	w.Resize(fyne.NewSize(1200,670))

	//Use customized window icon
	icon , err := fyne.LoadResourceFromPath("./cat.png")
	if err != nil {
		log.Fatal(err)
	}
	w.SetIcon(icon)

	//Run app
	w.ShowAndRun()	
}
