package views

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/client/group"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
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
			viewmodels.GroupInvitees.GetAllInvitees()
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
		Container:      mainContainer,
		LeftNavigation: leftNavigation,
		RightContent:   rightContent,

		BackupRightContent:  container.NewBorder(backupsToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(backupsListContainer))),
		InvitesRightContent: container.NewBorder(invitesToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(invitesListContainer))),

		BackupListContainer:  backupsListContainer,
		InvitesListContainer: invitesListContainer,
	}

	// Default content is backups.
	bc.DisplayBackupsContent()

	// Register listeners.
	// TODO: updated backups
	events.GroupInviteesUpdated.Register(bc)

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

func (c *backupsInvitedContent) HandleGroupInviteesUpdated() {
	c.rerenderInvitees()
}

func (c *backupsInvitedContent) DisplayInvitesContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.InvitesRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) DisplayBackupsContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.BackupRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) rerenderInvitees() {
	c.mu.Lock()
	global.Log("updating invitees list")
	c.InvitesListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.GroupInvitees.Models) == 0 {
		c.InvitesListContainer.Add(widget.NewLabel("There are no pending invites"))
		c.InvitesListContainer.Refresh()
		c.mu.Unlock()
		return
	}

	for _, i := range viewmodels.GroupInvitees.Models {
		c.InvitesListContainer.Add(NewInviteeCanvasObject(i))
	}

	c.InvitesListContainer.Refresh()
	c.mu.Unlock()
}

func NewInviteeCanvasObject(i *models.Invite) fyne.CanvasObject {
	return container.NewHBox(
		widget.NewLabel(fmt.Sprintf("%v @ %v", i.User.Username, i.Date)),
	)
}

func dialogCreateNewBackup() {
	// TODO Create new backup dialog.
}

func dialogInviteUser() {
	nameEntry := widget.NewEntry()
	nameEntry.Validator = func(s string) error {
		if len(s) < 3 {
			return errors.New("at least 3 characters long")
		}
		return nil
	}

	dialog.NewForm("Invite user to " + viewmodels.SelectedGroup.Model.Name, "Send", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("Name", nameEntry),
		}, func(b bool) {
			if b {
				inviteUserToGroup(nameEntry.Text)
			}
		}, global.MainWindow).Show()
}

func inviteUserToGroup(name string) {
	go func() {
		resp, err := server.Mnimidamon.Group.InviteUserToGroup(&group.InviteUserToGroupParams{
			Body:       &models.InviteUserPayload{Username: &name},
			GroupID:    viewmodels.SelectedGroup.Model.GroupID,
			Context:    server.ApiContext,
		}, viewmodels.CurrentComputer.Auth)

		if err != nil {
			if br, ok := err.(*group.InviteUserToGroupBadRequest); ok {
				infoDialog(br.GetPayload().Message)
				return
			}
			infoDialog(err.Error())
			return
		}

		// Add the created invitations to the group invitees.
		viewmodels.GroupInvitees.Add(resp.Payload)
	}()
}
