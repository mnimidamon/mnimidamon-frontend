package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/views/fragments"
	"sync"
)

func NewBackupsAndInvitedContent() *backupsInvitedContent {
	backupsLabel := widget.NewLabelWithStyle("backups", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	invitesLabel := widget.NewLabelWithStyle("invited", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	backupsToolbarLabel := fragments.NewToolbarObject(backupsLabel)
	invitesToolbarLabel := fragments.NewToolbarObject(invitesLabel)

	// For reference.
	var bc *backupsInvitedContent

	backupsToolbar := widget.NewToolbar(
		backupsToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			// TODO reload backups
		}),
		widget.NewToolbarAction(resources.DiskSaveSvg, func() {
			dialogCreateNewBackup()
		}),
	)

	invitesToolbar := widget.NewToolbar(
		invitesToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			// TODO reload invitees
		}),
		widget.NewToolbarAction(resources.EmailPlusSvg, func() {
			dialogInviteUser()
		}),
	)

	leftNavigation := container.NewVBox(
		widget.NewButtonWithIcon("backups", resources.DiskStackSvg, func() {
			bc.DisplayBackupsContent()
		}),
		widget.NewButtonWithIcon("invites", resources.EmailMultipleSvg, func() {
			bc.DisplayInvitesContent()
		}),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("groups", resources.GroupsSvg, func() {
			events.RequestGroupsContent.Trigger()
		}),
	)

	backupsListContainer := container.NewVBox(widget.NewLabel("loading..."))
	invitesListContainer := container.NewVBox(widget.NewLabel("loading..."))

	rightContent := container.NewMax()

	mainContainer := container.NewBorder(nil, nil, leftNavigation, nil, rightContent)

	bc = &backupsInvitedContent{
		Container:            mainContainer,
		LeftNavigation:       leftNavigation,
		RightContent:         rightContent,

		BackupRightContent:   container.NewBorder(backupsToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(backupsListContainer))),
		InvitesRightContent:  container.NewBorder(invitesToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(invitesListContainer))),

		BackupListContainer:  backupsListContainer,
		InvitesListContainer: invitesListContainer,
	}

	// Default content is backups.
	bc.DisplayBackupsContent()

	// Register listeners.
	// TODO: updated invites, updated backups

	return bc
}

type backupsInvitedContent struct {
	Container      *fyne.Container // The encapsulating container.
	LeftNavigation *fyne.Container // Left split content.
	RightContent   *fyne.Container // Right split content.

	BackupRightContent  *fyne.Container // Content displayed upon Invites navigation.
	InvitesRightContent *fyne.Container // Content displayed upon Backups navigation.

	BackupListContainer  *fyne.Container // Containing the backups list.
	InvitesListContainer *fyne.Container // Containing the invites group list.

	mu sync.Mutex // Lock when rendering UI elements.
}

func (c *backupsInvitedContent) DisplayInvitesContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.InvitesRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) DisplayBackupsContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.BackupRightContent}
	c.RightContent.Refresh()
}

func dialogCreateNewBackup() {
	// TODO Create new backup dialog.
}

func dialogInviteUser() {
	// TODO dialog invite user.
}