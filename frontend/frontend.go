package frontend

import (
	"fyne.io/fyne/v2"
	"mnimidamonbackend/frontend/views"
	"mnimidamonbackend/frontend/views/global"
)

type ApplicationEntryPoint interface {
	Run()
}

type applicationContainer struct {
	App        fyne.App
	MainWindow fyne.Window
}

func (ac *applicationContainer) Run() {
	// Get the first window that we have to show.

	// If the API keys are not present, display InputServerAndFolderPath.
	ac.MainWindow.SetContent(views.StartScreen.CanvasObject)

	// Else display main screen.
	ac.MainWindow.ShowAndRun()
}

func NewEntryPoint() ApplicationEntryPoint {
	return 	&applicationContainer{
		App: global.App,
		MainWindow: global.MainWindow,
	}
}
