// GOOS=js GOARCH=wasm go build -o static/mouse.wasm src/mouse/main.go

package main

import (
	"fmt"
	"math"
	"math/rand"
	"syscall/js"
	"time"
)

var (
	canvas  js.Value
	context js.Value
	width   int
	height  int
	rng     *rand.Rand
	catX    float64
	catY    float64
	targetX float64
	targetY float64
	speed   float64 = 2.0
	catSize float64 = 24.0
	mouseX  float64
	mouseY  float64
)

func getRandomPosition(max int) float64 {
	return float64(rng.Intn(max-int(catSize*2))) + catSize
}

func moveCat(this js.Value, p []js.Value) interface{} {
	// Calculate the direction vector
	dx := targetX - catX
	dy := targetY - catY
	distance := math.Sqrt(dx*dx + dy*dy)

	// If the cat is close to the target, choose a new target
	if distance < speed {
		targetX = getRandomPosition(width)
		targetY = getRandomPosition(height)
	} else {
		// Normalize the direction vector and move the cat
		catX += dx / distance * speed
		catY += dy / distance * speed
	}

	// Clear the canvas and draw the cat and mouse at the new positions
	context.Call("clearRect", 0, 0, width, height)
	context.Call("fillText", "🐱", catX, catY)
	context.Call("fillText", "🐭", mouseX, mouseY)

	return nil
}

func mouseMoveHandler(this js.Value, args []js.Value) interface{} {
	boundingRect := canvas.Call("getBoundingClientRect")
	mouseX = args[0].Get("clientX").Float() - boundingRect.Get("left").Float()
	mouseY = args[0].Get("clientY").Float() - boundingRect.Get("top").Float()
	return nil
}

func main() {
	fmt.Println("Hello from the Cat and Mouse!")
	js.Global().Get("console").Call("log", "Hello from Cat and Mouse WebAssembly!")

	// Initialize a new random number generator
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Get the canvas and context
	canvas = js.Global().Get("document").Call("getElementById", "myCanvas")
	context = canvas.Call("getContext", "2d")

	// Set canvas dimensions
	width = canvas.Get("width").Int()
	height = canvas.Get("height").Int()

	// Set font size for the cat and mouse emojis
	context.Set("font", "24px serif")

	// Initialize cat position and target position
	catX = getRandomPosition(width)
	catY = getRandomPosition(height)
	targetX = getRandomPosition(width)
	targetY = getRandomPosition(height)

	// Add event listener for mouse movements
	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(mouseMoveHandler))

	// Move the cat emoji every 16 milliseconds (approximately 60 FPS)
	js.Global().Call("setInterval", js.FuncOf(moveCat), 16)

	// Keep the Go program running
	select {}
}
