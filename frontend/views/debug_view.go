package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/frontend/events"
	_ "mnimidamonbackend/frontend/global"
)

var DebugView fyne.CanvasObject

func init() {
	viewNavigation := []struct {
		Name         string
		ListenerCall func()
	}{
		{Name: "Register View", ListenerCall: events.RequestRegisterView.Trigger},
		{Name: "Computer Register View", ListenerCall: events.RequestComputerRegisterView.Trigger},
		{Name: "Login View", ListenerCall: events.RequestLoginView.Trigger},
		{Name: "Main View", ListenerCall: events.RequestMainView.Trigger},
		{Name: "Restart", ListenerCall: events.RestartConfiguration.Trigger},
	}

	c := container.NewVBox()

	for _, vn := range viewNavigation {
		c.Add(widget.NewButton(vn.Name, vn.ListenerCall))
	}

	DebugView = c
}
