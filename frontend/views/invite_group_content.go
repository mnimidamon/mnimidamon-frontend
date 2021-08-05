package views

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/client/group"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
)

func NewGroupListContent() *groupInviteListContent {
	groupLabel := widget.NewLabelWithStyle("groups", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	inviteLabel := widget.NewLabelWithStyle("invites", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	groupToolbarLabel := fragments.NewToolbarObject(groupLabel)
	inviteToolbarLabel := fragments.NewToolbarObject(inviteLabel)

	var gilc *groupInviteListContent
	// Group toolbar
	groupToolbar := widget.NewToolbar(groupToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Groups.GetAllGroups()
		}),
		widget.NewToolbarAction(resources.GroupAddSvg, func() {
			groupAddDialog()
		}),
	)

	// Invite Toolbar
	inviteToolbar := widget.NewToolbar(inviteToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Invites.GetAllInvites()
		}),
	)

	// Left navigation.
	leftNavigation := container.NewVBox(
		widget.NewButtonWithIcon("groups", resources.GroupsSvg, func() {
			gilc.DisplayGroupsContent()
		}),
		widget.NewButtonWithIcon("invites", resources.InboxSvg, func() {
			gilc.DisplayInvitesContent()
		}),
	)

	// Right list content.
	groupListContainer := container.NewVBox(widget.NewLabel("loading..."))
	inviteListContainer := container.NewVBox(widget.NewLabel("loading..."))

	// The center content.
	rightContent := container.NewMax()

	// Border layout, where left is navigation and center content is the right content.
	mainContainer := container.NewBorder(nil, nil, leftNavigation, nil, rightContent)

	gilc = &groupInviteListContent{
		Container:      mainContainer,
		LeftNavigation: leftNavigation,
		RightContent:   rightContent,

		GroupRightContent:  container.NewBorder(groupToolbar, nil, nil ,nil , container.NewVScroll(groupListContainer)),
		InviteRightContent: container.NewBorder(inviteToolbar, nil, nil, nil,  container.NewVScroll(inviteListContainer)),

		GroupListContainer:  groupListContainer,
		InviteListContainer: inviteListContainer,
	}

	gilc.DisplayGroupsContent()

	events.GroupsUpdated.Register(gilc)
	events.InvitesUpdated.Register(gilc)

	return gilc
}

func groupAddDialog() {
	nameEntry := widget.NewEntry()
	nameEntry.Validator = func(s string) error {
		if len(s) < 3 {
			return errors.New("at least 3 characters long")
		}
		return nil
	}

	dialog.NewForm("New group", "Create", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("Name", nameEntry),
		}, func(b bool) {
			if b {
				err := createNewGroup(nameEntry.Text)
				if err != nil {
					infoDialog(err.Error())
				}
			}
		}, global.MainWindow).Show()
}

func infoDialog(msg string) {
	dialog.NewInformation("", msg, global.MainWindow).Show()
}

func createNewGroup(name string) error {
	resp, err := server.Mnimidamon.Group.CreateGroup(&group.CreateGroupParams{
		Body:    &models.GroupCreatePayload{Name: &name},
		Context: server.ApiContext,
	}, server.CompAuth)

	if err != nil {
		if br, ok := err.(*group.CreateGroupBadRequest); ok {
			return errors.New(br.GetPayload().Message)
		}
		return err
	}

	events.GroupCreated.Trigger(*resp.Payload)
	return nil
}

type groupInviteListContent struct {
	Container      *fyne.Container      // The encapsulating container.
	LeftNavigation *fyne.Container      // Left split content.
	RightContent   *fyne.Container // Right split content.

	GroupRightContent  *fyne.Container // Content displayed upon Invite navigation.
	InviteRightContent *fyne.Container // Content displayed upon Group navigation.

	GroupListContainer  *fyne.Container // Containing the group list.
	InviteListContainer *fyne.Container // Containing the invite group list.
}

func (c *groupInviteListContent) HandleGroupsUpdate() {
	global.Log("updating groups list")
	c.GroupListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.Groups.Models) == 0 {
		c.GroupListContainer.Add(widget.NewLabel("Create a group or accept an invite"))
		return
	}

	for _, g := range viewmodels.Groups.Models {
		c.GroupListContainer.Add(widget.NewLabel(g.Name))
	}

	c.GroupListContainer.Refresh()
}

func (c *groupInviteListContent) HandleInvitesUpdate() {
	global.Log("updating invites list")
	c.InviteListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.Invites.Models) == 0 {
		c.InviteListContainer.Add(widget.NewLabel("You have no pending invites"))
		return
	}

	for _, i := range viewmodels.Invites.Models {
		c.InviteListContainer.Add(widget.NewLabel(i.User.Username + " @ " + i.Date.String()))
	}

	c.InviteRightContent.Refresh()
}

func (c *groupInviteListContent) DisplayGroupsContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.GroupRightContent}
	c.RightContent.Refresh()
}
func (c *groupInviteListContent) DisplayInvitesContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.InviteRightContent}
	c.RightContent.Refresh()
}
