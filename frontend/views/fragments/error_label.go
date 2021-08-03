package fragments

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"time"
)

var (
	showTime = time.Second * 2
)

type FlashingLabel interface {
	fyne.CanvasObject
	ShowMessage(msg string)
}

type flashingLabel struct {
	*widget.Label
	Bind  binding.String
}

func NewFlashingLabel() FlashingLabel {
	bind := binding.NewString()
	label := widget.NewLabelWithData(bind)
	label.Hide()
	return &flashingLabel{
		Label: label,
		Bind:  bind,
	}
}

func (fl *flashingLabel) ShowMessage(msg string) {
	_ = fl.Bind.Set(msg)
	fl.Label.Show()
	go func() {
		time.Sleep(showTime)
		fl.Label.Hide()
	}()
}
