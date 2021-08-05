package views

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-openapi/strfmt"
	"mnimidamonbackend/client/authorization"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	_ "mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/models"
)

var LoginScreen fyne.CanvasObject

func init() {
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	usernameEntry.SetPlaceHolder("username")
	passwordEntry.SetPlaceHolder("password")

	usernameEntry.Validator = func(s string) error {
		if len(s) < 3 {
			return errors.New("should be longer than 3 characters")
		}
		return nil
	}

	passwordEntry.Validator = func(s string) error {
		if len(s) < 3 {
			return errors.New("should be longer than 3 characters")
		}
		return nil
	}

	toolbarLabel := widget.NewLabelWithStyle("Mnimidamon login", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	errorLabel := fragments.NewFlashingLabel()

	toolbar := widget.NewToolbar(
		fragments.NewToolbarObject(toolbarLabel),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.RestartAltSvg, func() {
			events.RestartConfiguration.Trigger()
		}),
		widget.NewToolbarAction(resources.AccountPlusOutlineSvg, func() {
			events.RequestRegisterView.Trigger()
		}),
	)


	form := &widget.Form{
		Items: []*widget.FormItem{
			widget.NewFormItem("Username", usernameEntry),
			widget.NewFormItem("Password", passwordEntry),
		},
		OnSubmit: func() {
			username, password := usernameEntry.Text, passwordEntry.Text

			go func() {
				// TODO LOADING

				// Call the api.
				resp, err := server.Mnimidamon.Authorization.LoginUser(&authorization.LoginUserParams{
					Body: &models.LoginPayload{
						Password: (*strfmt.Password)(&password),
						Username: &username,
					},
					Context: server.ApiContext,
				})

				if err != nil {
					if respErr, ok := err.(*authorization.LoginUserUnauthorized); ok {
						errorLabel.ShowMessage(respErr.Payload.Message)
					} else {
						errorLabel.ShowMessage(err.Error())
					}
					return
				}

				// Inform about the configuration confirm.
				events.ConfirmUserConfig.Trigger(global.UserConfig{
					User: models.User{
						UserID:   resp.Payload.User.UserID,
						Username: resp.Payload.User.Username,
					},
					Key: *resp.Payload.APIKey,
				})

				// Navigate to computer name input.
				events.RequestComputerRegisterView.Trigger()
			}()
		},
		OnCancel: func() {
			usernameEntry.SetText("")
			passwordEntry.SetText("")
		},
		SubmitText: "Login",
		CancelText: "Reset",
	}

	LoginScreen = container.NewVBox(
		toolbar,
		form,
		container.NewCenter(errorLabel),
	)
}
