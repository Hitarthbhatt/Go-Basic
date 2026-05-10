package main

import "fmt"

func main() {
	// --- didSet equivalent ---
	var p Player = Player{}
	p.SetScore(50)  // score changed: 0 → 50
	p.SetScore(100) // score changed: 50 → 100
	fmt.Println("Final score:", p.Score())

	// --- 1. var declaration (explicit type) ---
	var name string = "Hitarth"
	var age int = 25
	var pi float64 = 3.14
	var isLearning bool = true

	fmt.Println(name, age, pi, isLearning)

	// --- 2. Short declaration := (type inferred, like Swift's var/let) ---
	city := "Ahmedabad"
	score := 100
	fmt.Println(city, score)

	// --- 3. Zero values (Go's default, Swift has no equivalent) ---
	var zeroInt int       // 0
	var zeroStr string    // ""
	var zeroBool bool     // false
	var zeroFloat float64 // 0.0
	fmt.Println(zeroInt, zeroStr, zeroBool, zeroFloat)

	// --- 4. Multiple assignment ---
	x, y := 10, 20
	x, y = y, x
	fmt.Println("x, y value - ", x, y)

	// --- 5. Constants ---
	const maxScore = 100
	const appName = "GoLearn"
	fmt.Println(maxScore, appName)

	// --- 6. Type conversion (explicit, no implicit — unlike Swift) ---
	var intVal = 42
	floatVal := float64(intVal) // must cast manually
	fmt.Println(floatVal)
}
