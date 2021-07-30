package frontend

import (
	"fyne.io/fyne/v2"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views"
)

type ApplicationEntryPoint interface {
	Run()
}

type applicationContainer struct {
	App         fyne.App
	MainWindow  fyne.Window
	MainContent *fyne.Container
}

func (ac *applicationContainer) Run() {
	// Get the first window that we have to show.
	// If the configuration has not been stored yet then initialize the setup.
	// Otherwise launch the main window.

	if !services.ConfigurationStore.IsStored() {
		ac.SetMainContent(views.StartScreen)
	} else {
		ac.SetMainContent(views.LoginScreen)
	}

	// Else display main screen.
	ac.MainWindow.ShowAndRun()
}

func (ac *applicationContainer) SetMainContent(object fyne.CanvasObject) {
	object.Show()
	ac.MainWindow.SetContent(object)
}

// Event handlers.
func (ac *applicationContainer) HandleRequestLoginView() {
	// Routing to request login view handler. Replace the MainWindow content to the LoginView
	ac.SetMainContent(views.LoginScreen)
}

func (ac *applicationContainer) HandleRestartConfigurationHandler() {
	ac.SetMainContent(views.StartScreen)
}

func (ac *applicationContainer) HandleRequestRegisterView() {
	ac.SetMainContent(views.RegisterScreen)
}

// Constructor.
func NewEntryPoint() ApplicationEntryPoint {
	ac := &applicationContainer{
		App:        global.App,
		MainWindow: global.MainWindow,
	}

	// Listener setup.
	events.RequestLoginView.Register(ac)
	events.RequestRegisterView.Register(ac)
	events.RestartConfiguration.Register(ac)

	return ac
}
