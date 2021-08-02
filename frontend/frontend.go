package frontend

import (
	"fyne.io/fyne/v2"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views"
	"time"
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
	// Open the debug window.
	go func() {
		time.Sleep(time.Second)
		ac.OpenDebugWindowView()
	}()

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
	ac.MainWindow.SetContent(object)
	object.Show()
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

func (ac *applicationContainer) HandleRequestComputerRegisterView() {
	ac.SetMainContent(views.ComputerRegisterScreen)
}

func (ac *applicationContainer) HandleRequestMainView() {
	ac.SetMainContent(views.MainScreen)
}

func (ac *applicationContainer) OpenDebugWindowView() {
	dw := ac.App.NewWindow("Debug")
	dw.SetContent(views.DebugView)
	dw.Show()
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
	events.RequestComputerRegisterView.Register(ac)
	events.RequestMainView.Register(ac)

	return ac
}
