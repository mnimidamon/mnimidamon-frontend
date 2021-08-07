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
	membersLabel := widget.NewLabelWithStyle("members", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	backupsToolbarLabel := fragments.NewToolbarObject(backupsLabel)
	invitesToolbarLabel := fragments.NewToolbarObject(invitesLabel)
	membersToolbarLabel := fragments.NewToolbarObject(membersLabel)

	// For reference.
	var bc *backupsInvitedContent

	backupsToolbar := widget.NewToolbar(
		backupsToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Backups.GetAllBackups()
		}),
		widget.NewToolbarAction(resources.DiskSaveSvg, func() {
			dialogCreateNewBackup()
		}),
	)

	membersToolbar := widget.NewToolbar(
			membersToolbarLabel,
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(resources.SyncSvg, func() {
				viewmodels.GroupMembers.GetAllMembers()
				viewmodels.GroupComputers.GetAllGroupComputers()
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
		widget.NewButtonWithIcon("members", resources.AccountMultipleSvg, func() {
			bc.DisplayMembersContent()
		}),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("groups", resources.GroupsSvg, func() {
			events.RequestGroupsContent.Trigger()
		}),
	)

	backupsListContainer := container.NewVBox(NewItalicLabel("loading..."))
	invitesListContainer := container.NewVBox(NewItalicLabel("loading..."))
	membersListContainer := container.NewVBox(NewItalicLabel("loading..."))

	rightContent := container.NewMax()

	mainContainer := container.NewBorder(nil, nil, leftNavigation, nil, rightContent)

	bc = &backupsInvitedContent{
		Container:      mainContainer,
		LeftNavigation: leftNavigation,
		RightContent:   rightContent,

		BackupsRightContent: container.NewBorder(backupsToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(backupsListContainer))),
		InvitesRightContent: container.NewBorder(invitesToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(invitesListContainer))),
		MembersRightContent: container.NewBorder(membersToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(membersListContainer))),

		BackupsListContainer: backupsListContainer,
		InvitesListContainer: invitesListContainer,
		MembersListContainer: membersListContainer,
	}

	// Default content is backups.
	bc.DisplayBackupsContent()

	// Register listeners.
	events.GroupInviteesUpdated.Register(bc)
	events.GroupMembersUpdated.Register(bc)
	events.GroupComputersUpdated.Register(bc)
	events.BackupsUpdated.Register(bc)

	return bc
}

type backupsInvitedContent struct {
	Container      *fyne.Container // The encapsulating container.
	LeftNavigation *fyne.Container // Left split content.
	RightContent   *fyne.Container // Right split content.

	BackupsRightContent *fyne.Container // Content displayed upon Invites navigation.
	InvitesRightContent *fyne.Container // Content displayed upon Backups navigation.
	MembersRightContent *fyne.Container // Content displayed upon Members navigation.

	BackupsListContainer *fyne.Container // Containing the backups list.
	InvitesListContainer *fyne.Container // Containing the invites group list.
	MembersListContainer *fyne.Container // Containing the group members list.

	mu sync.Mutex // Lock when rendering UI elements.
}

func (c *backupsInvitedContent) HandleBackupsUpdate() {
	c.rerenderBackups()
}

func (c *backupsInvitedContent) HandleGroupComputersUpdated() {
	c.rerenderMembers()
}

func (c *backupsInvitedContent) HandleGroupMembersUpdated() {
	c.rerenderMembers()
}

func (c *backupsInvitedContent) HandleGroupInviteesUpdated() {
	c.rerenderInvitees()
}

func (c *backupsInvitedContent) DisplayInvitesContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.InvitesRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) DisplayBackupsContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.BackupsRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) rerenderInvitees() {
	c.mu.Lock()
	global.Log("updating invitees list")
	c.InvitesListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.GroupInvitees.Models) == 0 {
		c.InvitesListContainer.Add(NewItalicLabel("There are no pending invites"))
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

func (c *backupsInvitedContent) rerenderMembers() {
	global.Log("updating members list")

	c.mu.Lock()
	c.MembersListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.GroupMembers.Models) < 1 {
		c.MembersListContainer.Add(NewItalicLabel("loading..."))
		c.MembersListContainer.Refresh()
		c.mu.Unlock()
		return
	}

	if len(viewmodels.GroupMembers.Models) == 1 {
		c.MembersListContainer.Add(NewItalicLabel("You are the only member of the group"))
	}

	for _, m := range viewmodels.GroupMembers.Models {
		c.MembersListContainer.Add(NewMemberCanvasObject(m))
	}

	c.MembersListContainer.Refresh()
	c.mu.Unlock()
}

func (c *backupsInvitedContent) DisplayMembersContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.MembersRightContent}
	c.RightContent.Refresh()
}

func (c *backupsInvitedContent) rerenderBackups() {
	global.Log("updating backups list")

	c.mu.Lock()
	c.BackupsListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.Backups.Models) == 0 {
		c.BackupsListContainer.Add(NewItalicLabel("There are no backups in this group"))
		c.BackupsListContainer.Refresh()
		c.mu.Unlock()
		return
	}

	for _, b := range viewmodels.Backups.Models {
		c.BackupsListContainer.Add(NewBackupCanvasObject(b))
	}

	c.BackupsListContainer.Refresh()
	c.mu.Unlock()
}

func NewBackupCanvasObject(b *models.Backup) fyne.CanvasObject {
	return widget.NewLabel(b.Filename)
}

func NewMemberCanvasObject(m *models.User) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle(m.Username, fyne.TextAlignLeading, fyne.TextStyle{}),
		NewComputersCanvasObject(m.UserID),
	)
}

func NewComputersCanvasObject(userID int64) fyne.CanvasObject {
	c := container.NewVBox()

	for _, gc := range viewmodels.GroupComputers.GetAllOf(userID) {
		c.Add(NewGroupComputerCanvasObject(gc))
	}

	return c
}

func NewGroupComputerCanvasObject(gc *models.GroupComputer) fyne.CanvasObject {
	return widget.NewLabel(fmt.Sprintf("   %v  \t%vMB", gc.Computer.Name, gc.StorageSize / 1024))
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

func NewItalicLabel(msg string) *widget.Label{
	return widget.NewLabelWithStyle(msg, fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
}