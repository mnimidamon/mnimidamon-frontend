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
	"mnimidamonbackend/models"
)

var RegisterScreen fyne.CanvasObject

func init() {
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	passwordRepeatEntry := widget.NewPasswordEntry()

	usernameEntry.SetPlaceHolder("username")
	passwordEntry.SetPlaceHolder("password")
	passwordRepeatEntry.SetPlaceHolder("repeat password")

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

	passwordRepeatEntry.Validator = func(s string) error {
		if passwordRepeatEntry.Text != passwordEntry.Text {
			return errors.New("passwords should match")
		}
		return nil
	}

	toolbarLabel := widget.NewLabelWithStyle("user registration", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
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
			widget.NewFormItem("Username", usernameEntry),
			widget.NewFormItem("Password", passwordEntry),
			widget.NewFormItem("", passwordRepeatEntry),
		},
		OnSubmit: func() {
			username, password := usernameEntry.Text, passwordEntry.Text

			go func() {
				// TODO: LOADING?

				// Call the backend endpoint.
				resp, err := mnimidamon.Authorization.RegisterUser(&authorization.RegisterUserParams{
					Body: &models.RegisterPayload{
						Password: (*strfmt.Password)(&password),
						Username: &username,
					},
					Context: apiContext,
				})

				if err != nil {
					if respError, ok := err.(*authorization.RegisterUserBadRequest); ok {
						errorLabel.ShowMessage(respError.Payload.Code)
					} else {
						errorLabel.ShowMessage(err.Error())
					}
					return
				}

				// Inform about the user configuration.
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
			passwordRepeatEntry.SetText("")
		},
		SubmitText: "Register",
		CancelText: "Reset",
	}

	RegisterScreen = container.NewVBox(
		toolbar,
		form,
		container.NewCenter(errorLabel),
	)

}
