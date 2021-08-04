package views

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	_ "mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"net"
	"os"
	"strconv"
	"time"
)

var StartScreen fyne.CanvasObject

func init() {
	portEntry := widget.NewEntry()
	pathEntry := widget.NewEntry()
	hostEntry := widget.NewEntry()

	portEntry.SetPlaceHolder("1000")
	pathEntry.SetPlaceHolder("select a folder")
	hostEntry.SetPlaceHolder("127.0.0.1")

	// Folder field validator.
	pathEntry.Validator = func(s string) error {
		if _, err := os.Stat(s); os.IsNotExist(err) {
			return errors.New("folder does not exist")
		}
		return nil
	}

	hostEntry.OnChanged = func(s string) {
		if portEntry.Text != "" {
			portEntry.SetText("")
		}
	}

	hostEntry.Validator = func(s string) error {
		if len(s) == 0 {
			return errors.New("required")
		}
		return nil
	}

	portEntry.Validator = func(s string) error {
		if err := checkConnection(hostEntry.Text, s); err != nil {
			return err
		}
		return nil
	}

	// The dialog to get the folder selection.
	selectFolderDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if uri != nil {
			pathEntry.SetText(uri.Path())
		}
	}, global.MainWindow)

	// Show the dialog and its window.
	buttonSelectFolder := widget.NewButtonWithIcon("Select folder", resources.FolderOpenSvg, func() {
		selectFolderDialog.Show()
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			widget.NewFormItem("Server address", hostEntry),
			widget.NewFormItem("Port", portEntry),
			widget.NewFormItem("Storage folder", pathEntry),
			widget.NewFormItem("", buttonSelectFolder),
		},

		OnSubmit: func() {
			folder := pathEntry.Text
			port, _ := strconv.Atoi(portEntry.Text)
			host := hostEntry.Text

			// Distribute the event for configuration.
			events.ConfirmServerConfig.Trigger(global.ServerConfig{
				FolderPath: folder,
				Host:       host,
				Port:       port,
			})


			// Request Navigation to the Login View.
			events.RequestLoginView.Trigger()
		},

		OnCancel: func() {
			hostEntry.SetText("")
			pathEntry.SetText("")
			portEntry.SetText("")
		},
	}

	// The canvas object is our input form.
	StartScreen = container.NewPadded(
		container.NewVBox(
			// Title bar.
			widget.NewLabelWithStyle("Mnimidamon setup", fyne.TextAlignCenter, fyne.TextStyle{
				Monospace: true,
			}),
			form,
		))
}

func checkConnection(host string, port string) error {
	timeout := time.Millisecond * 100
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)

	if err != nil {
		global.Log("Error on host name input: %v", err)
		return errors.New("cannot establish connection")
	}

	if conn != nil {
		defer conn.Close()
	}

	return nil
}
