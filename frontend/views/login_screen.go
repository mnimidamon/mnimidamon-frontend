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

	toolbarLabel := widget.NewLabelWithStyle("mnimidamon login", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})

	toolbar := widget.NewToolbar(
		fragments.NewToolbarLabel(toolbarLabel),
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
			username, _ := usernameEntry.Text, passwordEntry.Text

			// TODO HTTP CALL ON LOGIN.

			//

			// Inform about the configuration confirm.
			events.ConfirmUserConfig.Trigger(global.UserConfig{
				User: models.User{
					UserID:   0,
					Username: username,
				},
				Key: "",
			})

			// Navigate to computer name input.
			events.RequestComputerRegisterView.Trigger()
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
	)
}
