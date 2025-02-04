// GOOS=js GOARCH=wasm go build -o mouse.wasm ./src/mouse

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

	Cat    Entity
	Mouse  Entity
	House  Entity
	Cheese Entity

	Score        int
	TopScore     int
	gameInterval js.Value
	gameOver     bool
}

func NewGame() *Game {
	canvas := js.Global().Get("document").Call("getElementById", "myCanvas")
	context := canvas.Call("getContext", "2d")
	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()

	game := &Game{
		Canvas:   canvas,
		Context:  context,
		Width:    width,
		Height:   height,
		House:    NewHouse(width, height),
		Mouse:    NewMouse(width, height),
		Cat:      NewCat(width, height),
		Cheese:   NewCheese(width, height),
		Score:    0,
		TopScore: 0,
	}
	game.spawnCat()
	game.spawnCheese()
	return game
}

func (g *Game) update() {
	if g.gameOver {
		return
	}

	g.updateCatMovement()
	g.constrainToBounds(&g.Cat)
	g.constrainToBounds(&g.Mouse)

	// Only check collision if the cat is active.
	if g.Cat.Active && g.checkCollision(g.Cat, g.Mouse) {
		g.endGame()
		return
	}

	// Check if the mouse has reached the cheese.
	if g.checkCollision(g.Mouse, g.Cheese) {
		g.checkCheeseCollision()
	}
}

// endGame stops the game loop and prompts a restart.
func (g *Game) endGame() {
	g.gameOver = true
	js.Global().Call("clearInterval", g.gameInterval)
}

// Implement the restart logic via a keydown event handler.
func (g *Game) handleRestart(_ js.Value, args []js.Value) interface{} {
	event := args[0]
	key := event.Get("key").String()
	if g.gameOver && (key == "r" || key == "R") {
		js.Global().Get("location").Call("reload")
	}
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
	game.initEventHandlers()

	fmt.Println("Game started!!!")
	select {}
}
