package main

import "math/rand"

// Константы
const (
	EmojiSize        = 24.0
	CatSpeed         = 2.0
	MouseSpeed       = 5.0
	AirplaneSpeed    = 10.0
	HouseSize        = 57.0
	CherryDuration   = 5000 // Milliseconds
	AirplaneDuration = 5000 // Milliseconds
)

// Entity - общая структура для всех объектов
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

// Создание мыши
func NewMouse(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:      float64(canvasWidth) / 2,
		Y:      float64(canvasHeight) / 2,
		Speed:  MouseSpeed,
		Emoji:  "🐭",
		Active: true,
	}
}

// Создание кошки
func NewCat(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:      rand.Float64()*(float64(canvasWidth)-40) + 20,
		Y:      rand.Float64()*(float64(canvasHeight)-40) + 20,
		Speed:  CatSpeed,
		Emoji:  "🐱",
		Active: true,
	}
}

// Создание домика
func NewHouse(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:     float64(canvasWidth) / 2,
		Y:     float64(canvasHeight) / 2,
		Size:  HouseSize,
		Emoji: "🏠",
	}
}

// Создание сыра
func NewCheese(canvasWidth, canvasHeight int) Entity {
	return Entity{
		X:     rand.Float64()*(float64(canvasWidth)-40) + 20,
		Y:     rand.Float64()*(float64(canvasHeight)-40) + 20,
		Emoji: "🧀",
	}
}

// Создание вишни
func NewCherry() Entity {
	return Entity{
		Emoji:    "🍒",
		Active:   false,
		Duration: CherryDuration,
	}
}

// Создание самолета
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
