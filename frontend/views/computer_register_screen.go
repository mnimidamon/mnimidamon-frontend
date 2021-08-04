package views

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/client/authorization"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	_ "mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/frontend/views/server"
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
	errorLabel := fragments.NewFlashingLabel()

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

			go func() {

				resp, err := server.Mnimidamon.Authorization.RegisterComputer(&authorization.RegisterComputerParams{
					Body:       &models.CreateComputerPayload{
						Name: &computerName,
					},
					Context: server.ApiContext,
				}, server.UserAuth)

				if err != nil {
					if respErr, ok := err.(*authorization.RegisterComputerBadRequest); ok {
						errorLabel.ShowMessage(respErr.Payload.Message)
					} else {
						errorLabel.ShowMessage(err.Error())
					}
					return
				}

				// Inform about the configuration confirm.
				events.ConfirmComputerConfig.Trigger(global.ComputerConfig{
					Computer: models.Computer{
						ComputerID: resp.Payload.Computer.ComputerID,
						Name:       resp.Payload.Computer.Name,
						OwnerID:    resp.Payload.Computer.OwnerID,
					},
					Key: resp.Payload.CompKey,
				})

				// Registration is complete, navigate to the main window.
				events.RequestMainView.Trigger()
			}()

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
		container.NewCenter(errorLabel),
	)
}
