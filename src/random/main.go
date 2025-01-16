// GOOS=js GOARCH=wasm go build -o static/random.wasm src/random/main.go

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"syscall/js"
	"time"
)

var (
	characters = "aueowrszxcvnm<>aueowrszxcvn"
	colors     = []string{
		"#78dce8",
		"#ffd866",
		"#a9dc76",
		"#ab9df2",
		"#ff6188",
		"#fc9867",
	}
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Get random character from characters string
func getRandomChar() string {
	return string(characters[rng.Intn(len(characters))])
}

// Get random color from colors array
func getRandomColor() string {
	return colors[rng.Intn(len(colors))]

}

// Scroll up 15 lines
func scrollUp(container js.Value) {
	htmlContent := container.Get("innerHTML").String()

	parts := strings.Split(htmlContent, "<br>")

	if len(parts) > 15 {
		newHTML := strings.Join(parts[len(parts)-15:], "<br>")
		container.Set("innerHTML", newHTML)
	}
}

// Generate random characters
func generateRandomChars(this js.Value, p []js.Value) interface{} {
	screenWidth := 15
	line := ""
	container := js.Global().Get("document").Call("createElement", "pre")
	js.Global().Get("document").Get("body").Call("appendChild", container)

	js.Global().Call("setInterval", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		randomChar := getRandomChar()
		randomColor := getRandomColor()
		line += "<span style=\"color: " + randomColor + "\">" + randomChar + "</span>"

		if len(line)/38 >= screenWidth {
			container.Set("innerHTML", container.Get("innerHTML").String()+line+"<br>")
			line = ""
		}
		// Scroll up
		scrollUp(container)
		return nil
	}), 100)

	return nil
}

func main() {
	fmt.Println("Hello from Go!")
	js.Global().Get("console").Call("log", "Hello from Go WebAssembly!")
	// Export function to JavaScript
	js.Global().Set("generateRandomChars", js.FuncOf(generateRandomChars))
	select {}
}
