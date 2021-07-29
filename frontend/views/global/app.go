package global

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"mnimidamonbackend/frontend/resources"
)

var Log = log.Printf
var App fyne.App
var MainWindow fyne.Window

func init() {
	App = app.NewWithID("frontend.mnimidamon.marmiha.com")
	App.SetIcon(resources.ResourceMnimidamonFrontendIconPng)

	MainWindow = App.NewWindow(" ")

	Log("Fyne app initialized with id %v", App.UniqueID())
}
