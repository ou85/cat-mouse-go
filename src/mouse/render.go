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
	g.Context.Set("font", "bold 48px")
	titleX := float64(g.Width)/2 - 150
	titleY := float64(g.Height) / 2.7
	g.Context.Call("fillText", "GAME OVER", titleX, titleY)

	// Draw current score.
	g.Context.Set("font", "36px")
	scoreText := fmt.Sprintf("Your Score: %d", g.Score)
	g.Context.Call("fillText", scoreText, titleX, float64(g.Height)/2.5+70)

	// Draw best score.
	highScoreText := fmt.Sprintf("Best Score: %d", g.TopScore)
	g.Context.Call("fillText", highScoreText, titleX, float64(g.Height)/2.5+130)

	// Draw restart instruction.
	g.Context.Set("font", "22px")
	restartText := "To restart, press R"
	// Center the text by calculating an approximate x position.
	restartX := float64(g.Width)/2 - 150
	restartY := float64(g.Height)/2.5 + 190
	g.Context.Call("fillText", restartText, restartX, restartY)
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
	if g.gameOver {
		g.drawGameOver()
	}
}
