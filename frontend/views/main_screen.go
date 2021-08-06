package views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/frontend/views/viewmodels"
	_ "mnimidamonbackend/frontend/views/viewmodels"
)

var MainScreen *mainScreen

func init() {
	toolbarContainer := container.NewMax()
	contentContainer := container.NewMax()
	mainContainer := container.NewBorder(toolbarContainer, nil, nil, nil, contentContainer)

	toolbarLabel := widget.NewLabelWithStyle("-@-", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	toolbarBind := binding.NewString()
	toolbarBind.AddListener(binding.NewDataListener(func() {
		str, _ := toolbarBind.Get()
		toolbarLabel.SetText(str)
	}))

	toolbar := widget.NewToolbar(
		fragments.NewToolbarObject(toolbarLabel),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.LogOutSvg, func() {
			showLogOutDialog()
		}),
	)

	toolbarContainer.Add(toolbar)

	// Initialize it to zero values.
	MainScreen = &mainScreen{
		Container:        mainContainer,
		ToolbarContainer: toolbarContainer,
		ContentContainer: contentContainer,
		Config:           nil,
		CurrentToolbar:   toolbar,
		ToolbarBind:      toolbarBind,

		GroupsInvitationsContent: NewGroupAndInvitationsContent(),
		BackupsInvitedContent:    NewBackupsAndInvitedContent(),
	}

	// Get the config if it's present.
	if services.ConfigurationStore.IsStored() {
		MainScreen.Config = services.ConfigurationStore.GetConfig()
	}

	// Handle config changes.
	events.ConfirmConfig.Register(MainScreen)
	events.SelectedGroupUpdated.Register(MainScreen)
	events.RequestGroupsContent.Register(MainScreen)

	// Refresh the UI to the current application state.
	MainScreen.refresh()

	// Set the current content to groups.
	MainScreen.SetGroupsContent()
}

type mainScreen struct {
	Container        *fyne.Container // Main Container including all the main screen content.
	ToolbarContainer *fyne.Container // Toolbar Container for toolbar replacement.
	ContentContainer *fyne.Container // Content Container for different content.

	CurrentToolbar *widget.Toolbar   // CurrentToolbar inside the ToolbarContainer.
	CurrentContent fyne.CanvasObject // CurrentContent inside the ContentContainer.

	ToolbarBind binding.String // Binding for the toolbar label name.
	Config      *global.Config

	GroupsInvitationsContent *groupsInvitationsContent // Content that represents the user groups and invitations to groups.
	BackupsInvitedContent    *backupsInvitedContent    // Content that displays a group backups and the sent invites.
}

// Navigate to groups content.
func (ms *mainScreen) HandleRequestGroupsContent() {
	ms.SetGroupsContent()
}

// When a new group is selected navigate to the backups content.
func (ms *mainScreen) HandleSelectedGroupUpdated() {
	ms.SetBackupsContent()
}

func (ms *mainScreen) SetBackupsContent() {
	ms.SetCurrentContent(ms.BackupsInvitedContent.Container)
}

func (ms *mainScreen) SetGroupsContent() {
	ms.SetCurrentContent(ms.GroupsInvitationsContent.Container)
}

func (ms *mainScreen) HandleConfirmConfig(config global.Config) {
	// Save the config.
	if ms.Config == nil {
		ms.Config = new(global.Config)
	}

	*ms.Config = config
	ms.refresh()
}

// Replaces the current content.
func (ms *mainScreen) SetCurrentContent(content fyne.CanvasObject) {
	ms.ContentContainer.Remove(ms.CurrentContent)
	ms.ContentContainer.Add(content)
	ms.CurrentContent = content
	ms.refresh()
}

// Refresh the UI based on the application state.
func (ms *mainScreen) refresh() {
	ms.refreshToolbar()
}

// Refresh the Toolbar based on the application state.
func (ms *mainScreen) refreshToolbar() {
	if ms.Config != nil {
		switch ms.CurrentContent {
		case ms.BackupsInvitedContent.Container:
			_ = ms.ToolbarBind.Set(ms.Config.Computer.Name + "@" + viewmodels.SelectedGroup.Group.Name)
			break
		case ms.GroupsInvitationsContent.Container:
			_ = ms.ToolbarBind.Set(ms.Config.User.Username + "@" + ms.Config.Computer.Name)
			break
		default:
			_ = ms.ToolbarBind.Set("undefined content")
		}
	}
}

func showLogOutDialog() {
	dialog.NewConfirm("Are you sure?", fmt.Sprintf("This will delete the computer %v and its backups.", services.ConfigurationStore.GetConfig().Computer.Name), func(b bool) {
		if b {
			events.RestartConfiguration.Trigger()
		}
	}, global.MainWindow).Show()
}
