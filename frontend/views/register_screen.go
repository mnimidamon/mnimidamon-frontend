package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "mnimidamonbackend/frontend/global"
)

var RegisterScreen fyne.CanvasObject

func init() {
	RegisterScreen = container.NewHBox(widget.NewLabel("Register Screen"))
}