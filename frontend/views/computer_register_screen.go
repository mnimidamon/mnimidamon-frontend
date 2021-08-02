package views

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	_ "mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/models"
)

var ComputerRegisterScreen fyne.CanvasObject

func init() {
	computerNameEntry := widget.NewEntry()
	computerNameEntry.SetPlaceHolder("damon")

	computerNameEntry.Validator = func(s string) error {
		if len(s) < 3 {
			return errors.New("should be longer than 3 characters")
		}
		return nil
	}

	toolbarLabel := widget.NewLabelWithStyle("computer registration", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})

	toolbar := widget.NewToolbar(
		fragments.NewToolbarLabel(toolbarLabel),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.RestartAltSvg, func() {
			events.RestartConfiguration.Trigger()
		}),
	)

	form := &widget.Form{
		Items: []*widget.FormItem{
			widget.NewFormItem("Computer name", computerNameEntry),
		},
		OnSubmit: func() {
			computerName := computerNameEntry.Text

			// TODO HTTP CALL ON COMPUTER REGISTER.

			//

			// Inform about the configuration confirm.
			events.ConfirmComputerConfig.Trigger(global.ComputerConfig{
				Computer: models.Computer{
					ComputerID: 0,
					Name:       computerName,
					OwnerID:    0,
				},
				Key: "",
			})

			// Registration is complete, navigate to the main window.
			events.RequestMainView.Trigger()

		},
		OnCancel: func() {
			computerNameEntry.SetText("")
		},
		SubmitText: "Confirm",
		CancelText: "Reset",
	}

	ComputerRegisterScreen = container.NewVBox(
		toolbar,
		form,
	)
}
