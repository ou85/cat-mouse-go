package main

import (
	"math"
	"math/rand"
	"syscall/js"
	"time"
)

func (g *Game) spawnCat() {
	time.AfterFunc(5*time.Second, func() {
		g.Cat.Active = true

		g.Cat.X = rand.Float64()*(float64(g.Width)-40) + 20
		g.Cat.Y = rand.Float64()*(float64(g.Height)-40) + 20
		g.Cat.Angle = math.Atan2(g.Cat.Y-g.House.Y, g.Cat.X-g.House.X)

	})
}

func (g *Game) spawnCheese() {
	g.Mouse.X = rand.Float64()*(float64(g.Width)-40) + 20
	g.Mouse.Y = rand.Float64()*(float64(g.Height)-40) + 20
	g.Mouse.Timer = 0
	g.Mouse.Active = true
}

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
	if g.Cat.Angle == 0 {
		dx := g.Cat.X - g.House.X
		dy := g.Cat.Y - g.House.Y
		g.Cat.Angle = math.Atan2(dy, dx)
	}

	// Increment the angle
	g.Cat.Angle = math.Mod(g.Cat.Angle+0.01, 2*math.Pi)

	// Calculate the new position
	g.Cat.X = g.House.X + math.Cos(g.Cat.Angle)*(radius+20)
	g.Cat.Y = g.House.Y + math.Sin(g.Cat.Angle)*(radius+20)
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
	halfSize := EmojiSize / 2

	e.X = math.Max(halfSize, math.Min(float64(g.Width)-halfSize, e.X))
	e.Y = math.Max(halfSize, math.Min(float64(g.Height)-halfSize, e.Y))
}

// Event Handlers
func (g *Game) mouseMoveHandler(this js.Value, args []js.Value) interface{} {
	boundingRect := g.Canvas.Call("getBoundingClientRect")
	g.Mouse.X = args[0].Get("clientX").Float() - boundingRect.Get("left").Float()
	g.Mouse.Y = args[0].Get("clientY").Float() - boundingRect.Get("top").Float()
	return nil
}

// Initialize Event Handlers
func (g *Game) initEventHandlers() {
	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(g.mouseMoveHandler))
	js.Global().Call("setInterval", js.FuncOf(g.gameLoop), UpdateRate)
}

// Check for Mouse-Cheese collision
func (g *Game) checkCheeseCollision() {
	if g.checkCollision(g.Mouse, g.Cheese) {
		g.Score++
		if g.Score > g.TopScore {
			g.TopScore = g.Score
		}
		g.Cheese = NewCheese(g.Width, g.Height)
	}
}

func (g *Game) checkCollision(e1, e2 Entity) bool {
	dx := e1.X - e2.X
	dy := e1.Y - e2.Y
	distance := math.Hypot(dx, dy)
	return distance < EmojiSize
}
