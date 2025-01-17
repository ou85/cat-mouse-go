// GOOS=js GOARCH=wasm go build -o static/mouse.wasm ./src/mouse

package main

import (
	"fmt"
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

	Cat   Entity
	Mouse Entity
	House Entity
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
	// game := &Game{

	// 	House: NewHouse(),
	// 	Mouse: NewMouse(),
	// 	Cat:   NewCat(),
	// }
	// game.initCanvas()
	game.spawnCat()
	return game
}

// Update and Render
func (g *Game) update() {
	g.updateCatMovement()
	g.constrainToBounds(&g.Cat)
	g.constrainToBounds(&g.Mouse)
}

func (g *Game) gameLoop(this js.Value, args []js.Value) interface{} {
	g.update()
	g.render()
	return nil
}

// Main
func main() {
	game := NewGame()
	game.initEventHandlers()

	fmt.Println("Game started!")
	select {}
}
