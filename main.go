package main

import (
	"workoutGo/internal/app"
)

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}
	app.Logger.Println("app is running...")
}
