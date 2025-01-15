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
	ctx.Call("clearRect", 0, 0, width, height)

	// Draw circle
	x := rand.Float64()*float64(width-60) + 30
	y := rand.Float64()*float64(height-60) + 30
	color := fmt.Sprintf("hsl(%d, 70%%, 50%%)", rand.Intn(360))

	ctx.Call("beginPath")
	ctx.Call("arc", x, y, 30, 0, 2*3.14159)
	ctx.Set("fillStyle", color)
	ctx.Call("fill")

	return nil
}

func main() {
	fmt.Println("Hello from Go!!")
	js.Global().Get("console").Call("log", "Hello from Go WebAssembly!")
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Export function to JavaScript
	js.Global().Set("drawCircle", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return drawCircle(this, args)
	}))

	// Block to keep the program running
	select {}
}
