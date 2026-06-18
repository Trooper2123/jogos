package main

import (
	"awesomeProject/ui"

	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()

	ui.Build(a)

	a.Run()

}
