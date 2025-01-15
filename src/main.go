package main

import (
	"fmt"
	"math/rand"
	"syscall/js"

	"time"
)

func drawCircle(this js.Value, args []js.Value) interface{} {
	canvas := js.Global().Get("document").Call("getElementById", "myCanvas")
	ctx := canvas.Call("getContext", "2d")

	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()

	// Clear canvas
	// ctx.Call("clearRect", 0, 0, width, height)

	// Draw circle
	x := rand.Float64()*float64(width-30) + 15
	y := rand.Float64()*float64(height-30) + 15
	color := fmt.Sprintf("hsl(%d, 70%%, 50%%, 0.8)", rand.Intn(360))

	ctx.Call("beginPath")
	ctx.Call("arc", x, y, 15, 0, 2*3.14159)
	ctx.Set("fillStyle", color)
	ctx.Call("fill")

	return nil
}

func main() {
	fmt.Println("Hello from Go!")
	js.Global().Get("console").Call("log", "Hello from Go WebAssembly!")
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Export function to JavaScript
	js.Global().Set("drawCircle", js.FuncOf(drawCircle))

	// Create a ticker that triggers every 5 seconds
	// ticker := time.NewTicker(5 * time.Second)

	// Create a ticker that triggers every `t` seconds
	t := 0.3
	ticker := time.NewTicker(time.Duration(t * float64(time.Second)))
	defer ticker.Stop()

	// Run the drawCircle function every 1 seconds
	go func() {
		for range ticker.C {
			js.Global().Call("drawCircle")
		}
	}()

	// Block to keep the program running
	select {}
}
