package main

import "math/rand"

const (
	EmojiSize        = 24.0
	CatSpeed         = 2.0
	MouseSpeed       = 5.0
	AirplaneSpeed    = 10.0
	HouseSize        = 57.0
	CherryDuration   = 5000 // Milliseconds
	AirplaneDuration = 5000 // Milliseconds
)

type Entity struct {
	X, Y      float64
	Speed     float64
	Emoji     string
	Active    bool
	Timer     int
	Direction struct {
		X, Y float64
	}
	Angle      float64
	Circling   bool
	SpawnCount int
	Duration   int
	Size       float64
}

func NewMouse(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:      float64(canvasWidth) / 2,
		Y:      float64(canvasHeight) / 2,
		Speed:  MouseSpeed,
		Emoji:  "🐭",
		Active: true,
	}
}

func NewCat(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:      rand.Float64()*(float64(canvasWidth)-40) + 20,
		Y:      rand.Float64()*(float64(canvasHeight)-40) + 20,
		Speed:  CatSpeed,
		Emoji:  "🐱",
		Active: false,
	}
}

func NewHouse(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:     float64(canvasWidth) / 2,
		Y:     float64(canvasHeight) / 2,
		Size:  HouseSize,
		Emoji: "🏠",
	}
}

func NewCheese(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:     rand.Float64()*(float64(canvasWidth)-40) + 20,
		Y:     rand.Float64()*(float64(canvasHeight)-40) + 20,
		Emoji: "🧀",
	}
}

func NewCherry() Entity {
	return Entity{
		Emoji:    "🍒",
		Active:   false,
		Duration: CherryDuration,
	}
}

func NewAirplane(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:        float64(canvasWidth) / 2,
		Y:        float64(canvasHeight) / 2,
		Speed:    AirplaneSpeed,
		Duration: AirplaneDuration,
		Emoji:    "✈️",
		Active:   false,
	}
}
