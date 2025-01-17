// GOOS=js GOARCH=wasm go build -o static/mouse.wasm ./src/mouse

package main

import (
	"fmt"
	"math"
	"math/rand"
	"syscall/js"
)

// Settings
const (
	UpdateRate = 16 // milliseconds around 60 FPS
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
		House:   NewHouse(width, height),
		Mouse:   NewMouse(width, height),
		Cat:     NewCat(width, height),
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

// Cat movement strategies
func (g *Game) circleAroundHouse() {
	radius := HouseSize

	// Calculate the initial angle, if the cat has just started circling around.
	if g.CatAngle == 0 {
		dx := g.Cat.X - g.House.X
		dy := g.Cat.Y - g.House.Y
		g.CatAngle = math.Atan2(dy, dx)
	}

	// Increment the angle
	g.CatAngle = math.Mod(g.CatAngle+0.01, 2*math.Pi)

	// Calculate the new position
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
	// const emojiSize = 24.0
	halfSize := EmojiSize / 2

	e.X = math.Max(halfSize, math.Min(float64(g.Width)-halfSize, e.X))
	e.Y = math.Max(halfSize, math.Min(float64(g.Height)-halfSize, e.Y))
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

	// Create a radial gradient
	gradient := g.Context.Call("createRadialGradient",
		g.House.X, g.House.Y, HouseSize/2, // Start point (inner radius)
		g.House.X, g.House.Y, HouseSize/2+20) // End point (outer radius)
	gradient.Call("addColorStop", 0, "rgba(0, 255, 0, 0.05)") // Start: bright green
	gradient.Call("addColorStop", 1, "rgba(0, 100, 0, 0.75)") // End: dark green

	// Fill the circle with the gradient
	g.Context.Call("beginPath")
	g.Context.Call("arc", g.House.X, g.House.Y, HouseSize/2+20, 0, 2*math.Pi)
	g.Context.Set("fillStyle", gradient)
	g.Context.Call("fill")

	// Draw the circle border
	g.Context.Set("lineWidth", 2)
	g.Context.Set("strokeStyle", "rgba(0, 100, 0, 1.0)") // Dark green border
	g.Context.Call("stroke")

	// Draw the entities
	g.Context.Call("fillText", "🏠", g.House.X-12, g.House.Y+8)
	g.Context.Call("fillText", "🐱", g.Cat.X-12, g.Cat.Y+8)
	g.Context.Call("fillText", "🐭", g.Mouse.X-12, g.Mouse.Y+8)
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
	game := NewGame()

	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(game.mouseMoveHandler))
	js.Global().Call("setInterval", js.FuncOf(game.gameLoop), UpdateRate)

	fmt.Println("Game started!")
	select {}
}
