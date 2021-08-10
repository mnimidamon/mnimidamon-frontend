package events

import (
	"fyne.io/fyne/v2"
)

var ProcessStarted processStarted

type ProcessStartedHandler interface {
	HandleProcessStarted(process LoaderProcess)
}


type processStarted struct {
	handlers []ProcessStartedHandler
}

func (e *processStarted) Register(handler ProcessStartedHandler) {
	e.handlers = append(e.handlers, handler)
}


func (e *processStarted) Trigger(process LoaderProcess)  {
	for _, handler := range e.handlers {
		go handler.HandleProcessStarted(process)
	}
}

type LoaderProcess interface {
	AddToParentContainer(parent *fyne.Container) // Add the UI to parent container
	RemoveFromParentContainer()                  // Remove it from the parent UI container and stop refreshing.

	StartRefreshing() // Start refreshing UI
	StopRefreshing()  // Stop refreshing UI
}
