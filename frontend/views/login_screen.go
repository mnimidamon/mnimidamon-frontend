package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	_ "mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/models"
)

var LoginScreen fyne.CanvasObject

func init() {
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	toolbarLabel := widget.NewLabelWithStyle("mnimidamon login", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})

	toolbar := widget.NewToolbar(
		fragments.NewToolbarLabel(toolbarLabel),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			events.RequestRegisterView.Trigger()
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			events.RestartConfiguration.Trigger()
		}),
	)

	form := &widget.Form{
		Items:      []*widget.FormItem{
			widget.NewFormItem("Username", usernameEntry),
			widget.NewFormItem("Password", passwordEntry),
		},
		OnSubmit: func() {
			username, _ := usernameEntry.Text, passwordEntry.Text

			events.ConfirmUserConfig.Trigger(global.UserConfig{
				User: models.User{
					UserID:   0,
					Username: username,
				},
				Key:  "",
			})
		},
		OnCancel: func() {
			usernameEntry.SetText("")
			passwordEntry.SetText("")
		},
		SubmitText: "Login",
		CancelText: "Cancel",
	}

	LoginScreen = container.NewVBox(
		toolbar,
		form,
	)
}