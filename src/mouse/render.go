package main

import (
	"math"
)

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
	g.Context.Call("fillText", "🐭", g.Mouse.X-12, g.Mouse.Y+8)
	if g.Cat.Active {
		g.Context.Call("fillText", g.Cat.Emoji, g.Cat.X-12, g.Cat.Y+8)
	}
}
