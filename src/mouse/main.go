// GOOS=js GOARCH=wasm go build -o static/mouse.wasm src/mouse/main.go

// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"syscall/js"
// 	"time"
// )

// // DOM Elements
// var (
// 	canvas  js.Value
// 	context js.Value
// 	width   int
// 	height  int
// )

// // Game Variables
// var (
// 	catX      float64
// 	catY      float64
// 	catAngle  float64
// 	speed     float64 = 2.0
// 	mouseX    float64
// 	mouseY    float64
// 	houseX    float64
// 	houseY    float64
// 	houseSize float64 = 57.0
// )

// // Functions
// func spawnCat() {
// 	catX = rand.Float64()*(float64(width)-40) + 20
// 	catY = rand.Float64()*(float64(height)-40) + 20
// 	catAngle = math.Atan2(catY-houseY, catX-houseX)
// }

// func isPlayerInHouse() bool {
// 	dx := mouseX - houseX
// 	dy := mouseY - houseY
// 	distance := math.Hypot(dx, dy)
// 	return distance < houseSize/2
// }

// func isCatNearHouse() bool {
// 	dx := catX - houseX
// 	dy := catY - houseY
// 	distance := math.Hypot(dx, dy)
// 	return distance < (houseSize/2 + 20)
// }

// func avoidHouse() {
// 	dx := catX - houseX
// 	dy := catY - houseY
// 	distance := math.Hypot(dx, dy)

// 	if distance > 0 {
// 		catX += (dx / distance) * speed
// 		catY += (dy / distance) * speed
// 	} else {
// 		catX += (rand.Float64() - 0.5) * speed * 2
// 		catY += (rand.Float64() - 0.5) * speed * 2
// 	}
// }

// func circleAroundHouse() {
// 	radius := houseSize
// 	catAngle += 0.01
// 	catX = houseX + math.Cos(catAngle)*(radius+20)
// 	catY = houseY + math.Sin(catAngle)*(radius+20)
// }

// func handleCatBounds() {
// 	if catX < 0 {
// 		catX = 0
// 	}
// 	if catX > float64(width) {
// 		catX = float64(width)
// 	}
// 	if catY < 0 {
// 		catY = 0
// 	}
// 	if catY > float64(height) {
// 		catY = float64(height)
// 	}
// }

// func handleMouseBounds() {
// 	if mouseX < 0 {
// 		mouseX = 0
// 	}
// 	if mouseX > float64(width) {
// 		mouseX = float64(width)
// 	}
// 	if mouseY < 0 {
// 		mouseY = 0
// 	}
// 	if mouseY > float64(height) {
// 		mouseY = float64(height)
// 	}
// }

// func update() {
// 	if isPlayerInHouse() {
// 		circleAroundHouse()
// 	} else if isCatNearHouse() {
// 		avoidHouse()
// 	} else {
// 		dx := mouseX - catX
// 		dy := mouseY - catY
// 		distance := math.Sqrt(dx*dx + dy*dy)

// 		if distance > speed {
// 			catX += dx / distance * speed
// 			catY += dy / distance * speed
// 		}
// 	}

// 	handleCatBounds()
// 	handleMouseBounds()
// }

// func draw() {
// 	// Clear the canvas and draw the cat, mouse, and house at the new positions
// 	context.Call("clearRect", 0, 0, width, height)
// 	context.Call("fillText", "🏠", houseX-houseSize/2, houseY-houseSize/2)
// 	context.Call("fillText", "🐱", catX, catY)
// 	context.Call("fillText", "🐭", mouseX, mouseY)
// }

// func gameLoop(this js.Value, p []js.Value) interface{} {
// 	update()
// 	draw()
// 	return nil
// }

// func mouseMoveHandler(this js.Value, args []js.Value) interface{} {
// 	boundingRect := canvas.Call("getBoundingClientRect")
// 	mouseX = args[0].Get("clientX").Float() - boundingRect.Get("left").Float()
// 	mouseY = args[0].Get("clientY").Float() - boundingRect.Get("top").Float()
// 	return nil
// }

// func initializeGame() {
// 	width = canvas.Get("width").Int()
// 	height = canvas.Get("height").Int()

// 	context.Set("font", "24px serif")

// 	houseX = float64(width) / 2
// 	houseY = float64(height) / 2

// 	spawnCat()
// }

// func main() {
// 	fmt.Println("Hello from the Cat and Mouse!")
// 	js.Global().Get("console").Call("log", "Hello from Cat and Mouse WebAssembly!")

// 	rand.New(rand.NewSource(time.Now().UnixNano()))

// 	canvas = js.Global().Get("document").Call("getElementById", "myCanvas")
// 	context = canvas.Call("getContext", "2d")

// 	initializeGame()

// 	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(mouseMoveHandler))

// 	js.Global().Call("setInterval", js.FuncOf(gameLoop), 16)

// 	select {}
// }
//\
//

// =========================================================
//
//

package main

import (
	"fmt"
	"math"
	"math/rand"
	"syscall/js"
)

// Constants
const (
	CatSpeed   = 2.0
	HouseSize  = 57.0
	UpdateRate = 16 // milliseconds
)

// Game Elements
type Game struct {
	Canvas  js.Value
	Context js.Value
	Width   int
	Height  int

	Cat      Entity
	Mouse    Entity
	House    Entity
	CatAngle float64
}

type Entity struct {
	X, Y float64
}

// Initialize the game
func NewGame() *Game {
	canvas := js.Global().Get("document").Call("getElementById", "myCanvas")
	context := canvas.Call("getContext", "2d")
	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()

	game := &Game{
		Canvas:  canvas,
		Context: context,
		Width:   width,
		Height:  height,
		House:   Entity{X: float64(width) / 2, Y: float64(height) / 2},
		Mouse:   Entity{X: float64(width) / 2, Y: float64(height) / 2},
	}
	game.spawnCat()
	return game
}

func (g *Game) spawnCat() {
	g.Cat.X = rand.Float64()*(float64(g.Width)-40) + 20
	g.Cat.Y = rand.Float64()*(float64(g.Height)-40) + 20
	g.CatAngle = math.Atan2(g.Cat.Y-g.House.Y, g.Cat.X-g.House.X)
}

// Game Logic
func (g *Game) isInHouse(e Entity) bool {
	dx := e.X - g.House.X
	dy := e.Y - g.House.Y
	return math.Hypot(dx, dy) < HouseSize/2
}

func (g *Game) isCatNearHouse() bool {
	dx := g.Cat.X - g.House.X
	dy := g.Cat.Y - g.House.Y
	return math.Hypot(dx, dy) < (HouseSize/2 + 20)
}

func (g *Game) avoidHouse() {
	dx := g.Cat.X - g.House.X
	dy := g.Cat.Y - g.House.Y
	distance := math.Hypot(dx, dy)

	if distance > 0 {
		g.Cat.X += (dx / distance) * CatSpeed
		g.Cat.Y += (dy / distance) * CatSpeed
	}
}

func (g *Game) circleAroundHouse() {
	radius := HouseSize
	g.CatAngle = math.Mod(g.CatAngle+0.01, 2*math.Pi)
	g.Cat.X = g.House.X + math.Cos(g.CatAngle)*(radius+20)
	g.Cat.Y = g.House.Y + math.Sin(g.CatAngle)*(radius+20)
}

func (g *Game) updateCatMovement() {
	if g.isInHouse(g.Mouse) {
		g.circleAroundHouse()
	} else if g.isCatNearHouse() {
		g.avoidHouse()
	} else {
		dx := g.Mouse.X - g.Cat.X
		dy := g.Mouse.Y - g.Cat.Y
		distance := math.Hypot(dx, dy)

		if distance > CatSpeed {
			g.Cat.X += dx / distance * CatSpeed
			g.Cat.Y += dy / distance * CatSpeed
		}
	}
}

func (g *Game) constrainToBounds(e *Entity) {
	if e.X < 0 {
		e.X = 0
	}
	if e.X > float64(g.Width) {
		e.X = float64(g.Width)
	}
	if e.Y < 0 {
		e.Y = 0
	}
	if e.Y > float64(g.Height) {
		e.Y = float64(g.Height)
	}
}

// Update and Render
func (g *Game) update() {
	g.updateCatMovement()
	g.constrainToBounds(&g.Cat)
	g.constrainToBounds(&g.Mouse)
}

func (g *Game) render() {
	g.Context.Call("clearRect", 0, 0, g.Width, g.Height)
	g.Context.Set("font", "24px Arial")
	// g.Context.Call("fillText", "🏠", g.House.X-HouseSize/2, g.House.Y-HouseSize/2)
	// g.Context.Call("fillText", "🐱", g.Cat.X, g.Cat.Y)
	// g.Context.Call("fillText", "🐭", g.Mouse.X, g.Mouse.Y)
	g.Context.Call("fillText", "🏠", g.House.X-12, g.House.Y+8)
	g.Context.Call("fillText", "🐱", g.Cat.X-12, g.Cat.Y+8)
	g.Context.Call("fillText", "🐭", g.Mouse.X-12, g.Mouse.Y+8)

	// Draw safe circle around the house
	g.Context.Call("beginPath")
	g.Context.Call("arc", g.House.X, g.House.Y, HouseSize/2+20, 0, 2*math.Pi)
	g.Context.Set("lineWidth", 2)
	g.Context.Set("strokeStyle", "rgba(173, 216, 230, 0.7)") // Light blue with transparency

	g.Context.Call("stroke")
}

// Event Handlers
func (g *Game) mouseMoveHandler(this js.Value, args []js.Value) interface{} {
	boundingRect := g.Canvas.Call("getBoundingClientRect")
	g.Mouse.X = args[0].Get("clientX").Float() - boundingRect.Get("left").Float()
	g.Mouse.Y = args[0].Get("clientY").Float() - boundingRect.Get("top").Float()
	return nil
}

func (g *Game) gameLoop(this js.Value, args []js.Value) interface{} {
	g.update()
	g.render()
	return nil
}

// Main
func main() {
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	game := NewGame()

	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(game.mouseMoveHandler))
	js.Global().Call("setInterval", js.FuncOf(game.gameLoop), UpdateRate)

	fmt.Println("Game started!")
	select {}
}
