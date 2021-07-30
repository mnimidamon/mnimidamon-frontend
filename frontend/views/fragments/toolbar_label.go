package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type toolbarLabel struct {
	*widget.Label
}

func (t toolbarLabel) ToolbarObject() fyne.CanvasObject {
	return t.Label
}

func NewToolbarLabel(label *widget.Label) widget.ToolbarItem {
	return &toolbarLabel{label}
}
