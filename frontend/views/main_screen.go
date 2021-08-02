package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var MainScreen fyne.CanvasObject

func init() {
	MainScreen = widget.NewLabel("Main Screen")
}