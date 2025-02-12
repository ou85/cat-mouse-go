package main

import (
	"fmt"
	"math"
	"syscall/js"
)

func (g *Game) drawGameOver() {
	if !g.gameOver {
		return
	}

	// Dark overlay
	g.Context.Set("fillStyle", "rgba(0, 0, 0, 0.8)")
	g.Context.Call("fillRect", 0, 0, g.Width, g.Height)
	// White text color
	g.Context.Set("fillStyle", "white")

	// Draw "GAME OVER" title.
	g.Context.Set("font", "bold 48px 'Courier New', monospace")
	titleX := float64(g.Width)/2 - 150
	titleY := float64(g.Height) / 3.3
	g.Context.Call("fillText", "GAME OVER", titleX, titleY)

	// Draw current score.
	g.Context.Set("font", "30px 'Courier New', monospace")
	scoreText := fmt.Sprintf("YOUR SCORE: %d", g.Score)
	g.Context.Call("fillText", scoreText, titleX, float64(g.Height)/2.5+50)

	// Draw best score.
	highScoreText := fmt.Sprintf("BEST SCORE: %d", g.TopScore)
	g.Context.Call("fillText", highScoreText, titleX, float64(g.Height)/2.5+90)

	// Draw restart instruction.
	g.Context.Set("font", "20px 'Courier New', monospace")
	restartText := "press 'R' to restart"
	// Center the text by calculating an approximate x position.
	restartX := float64(g.Width)/2 - 150
	restartY := float64(g.Height)/2.5 + 190
	g.Context.Call("fillText", restartText, restartX, restartY)
}

// Fetches the score_table.json file and renders the score table
func (g *Game) renderScoreTable() {
	fetchPromise := js.Global().Call("fetch", "score_table.json")
	thenFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]
		jsonPromise := response.Call("json")
		thenJSON := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := args[0]
			// Sort the data array in descending order by score.
			data = data.Call("sort", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				a := args[0]
				b := args[1]
				return b.Get("score").Int() - a.Get("score").Int()
			}))

			// Build the HTML table.
			tableHTML := "<table border='0' cellpadding='5' cellspacing='0'><caption>BEST EVER 5</caption><tbody>"
			length := data.Length()
			if length > 5 {
				length = 5
			}
			for i := 0; i < length; i++ {
				entry := data.Index(i)
				name := entry.Get("name").String()
				score := entry.Get("score").Int()
				tableHTML += fmt.Sprintf("<tr><td>%s</td><td style='text-align: right'>%d</td></tr>", name, score)
			}
			tableHTML += "</tbody></table>"

			// Set the inner HTML of the scoreTable div.
			scoreTableDiv := js.Global().Get("document").Call("getElementById", "scoreTable")
			scoreTableDiv.Set("innerHTML", tableHTML)
			return nil
		})
		jsonPromise.Call("then", thenJSON)
		return nil
	})
	fetchPromise.Call("then", thenFunc)
}

func (g *Game) render() {
	g.Context.Call("clearRect", 0, 0, g.Width, g.Height)
	g.Context.Set("font", "24px Arial")

	scoreContainer := js.Global().Get("document").Call("getElementById", "score")
	scoreContainer.Set("innerText", fmt.Sprintf("SCORE: %d", g.Score))

	topScoreContainer := js.Global().Get("document").Call("getElementById", "topScore")
	topScoreContainer.Set("innerText", fmt.Sprintf("HIGH: %d", g.TopScore))

	// Create a radial gradient for the house area.
	gradient := g.Context.Call("createRadialGradient",
		g.House.X, g.House.Y, HouseSize/2,
		g.House.X, g.House.Y, HouseSize/2+20)
	gradient.Call("addColorStop", 0, "rgba(0, 255, 0, 0.05)")
	gradient.Call("addColorStop", 1, "rgba(0, 100, 0, 0.75)")

	// Draw the house with gradient.
	g.Context.Call("beginPath")
	g.Context.Call("arc", g.House.X, g.House.Y, HouseSize/2+20, 0, 2*math.Pi)
	g.Context.Set("fillStyle", gradient)
	g.Context.Call("fill")
	g.Context.Set("lineWidth", 2)
	g.Context.Set("strokeStyle", "rgba(0, 100, 0, 1.0)")
	g.Context.Call("stroke")

	// Draw game entities.
	g.Context.Call("fillText", g.House.Emoji, g.House.X-12, g.House.Y+8)
	g.Context.Call("fillText", g.Mouse.Emoji, g.Mouse.X-12, g.Mouse.Y+8)
	g.Context.Call("fillText", g.Cheese.Emoji, g.Cheese.X-12, g.Cheese.Y+8)
	if g.Cat.Active {
		g.Context.Call("fillText", g.Cat.Emoji, g.Cat.X-12, g.Cat.Y+8)
	}

	// If game over, overlay the Game Over screen.
	// if g.gameOver {
	// 	g.drawGameOver()
	// }
	if g.gameOver {
		g.drawGameOver()
	} else {
		g.renderScoreTable()
	}
}
