package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type toolbarLabel struct {
	fyne.CanvasObject
}

func (t toolbarLabel) ToolbarObject() fyne.CanvasObject {
	return t.CanvasObject
}

func NewToolbarObject(obj fyne.CanvasObject) widget.ToolbarItem {
	return &toolbarLabel{obj}
}
