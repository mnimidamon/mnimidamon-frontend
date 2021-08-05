package views

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/client/group"
	"mnimidamonbackend/client/invite"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
	"sync"
)

func NewGroupListContent() *groupInviteListContent {

	// Base toolbar elements.
	groupLabel := widget.NewLabelWithStyle("groups", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	inviteLabel := widget.NewLabelWithStyle("invites", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	groupToolbarLabel := fragments.NewToolbarObject(groupLabel)
	inviteToolbarLabel := fragments.NewToolbarObject(inviteLabel)

	// For reference
	var gilc *groupInviteListContent

	// Group toolbar
	groupToolbar := widget.NewToolbar(groupToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Groups.GetAllGroups()
			viewmodels.GroupComputers.GetAllGroupComputers()
		}),
		widget.NewToolbarAction(resources.GroupAddSvg, func() {
			groupAddDialog()
		}),
	)

	// Invite Toolbar.
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

		GroupRightContent:  container.NewBorder(groupToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(groupListContainer))),
		InviteRightContent: container.NewBorder(inviteToolbar, nil, nil, nil, container.NewVScroll(container.NewPadded(inviteListContainer))),

		GroupListContainer:  groupListContainer,
		InviteListContainer: inviteListContainer,
	}

	// Default content is groups.
	gilc.DisplayGroupsContent()

	// Register listeners.
	events.GroupsUpdated.Register(gilc)
	events.InvitesUpdated.Register(gilc)
	events.GroupComputersUpdated.Register(gilc)

	// Resize it to look somewhat better.
	global.MainWindow.Resize(fyne.Size{
		Width:  350,
		Height: 220,
	})

	return gilc
}

type groupInviteListContent struct {
	Container      *fyne.Container // The encapsulating container.
	LeftNavigation *fyne.Container // Left split content.
	RightContent   *fyne.Container // Right split content.

	GroupRightContent  *fyne.Container // Content displayed upon Invite navigation.
	InviteRightContent *fyne.Container // Content displayed upon Group navigation.

	GroupListContainer  *fyne.Container // Containing the group list.
	InviteListContainer *fyne.Container // Containing the invite group list.

	mu sync.Mutex // Lock when rendering UI elements.
}

func (c *groupInviteListContent) HandleGroupComputersUpdated() {
	c.rerenderGroups()
}

func (c *groupInviteListContent) HandleGroupsUpdate() {
	c.rerenderGroups()
}

func (c *groupInviteListContent) HandleInvitesUpdate() {
	c.rerenderInvites()
}

func (c *groupInviteListContent) rerenderInvites() {
	c.mu.Lock()

	global.Log("updating invites list")
	c.InviteListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.Invites.Models) == 0 {
		c.InviteListContainer.Add(widget.NewLabel("You have no pending invites"))
		c.mu.Unlock()
		return
	}

	for _, i := range viewmodels.Invites.Models {
		c.InviteListContainer.Add(NewInviteCanvasObject(i))
	}

	c.mu.Unlock()
}

func (c *groupInviteListContent) rerenderGroups() {
	c.mu.Lock()
	global.Log("updating groups list")
	c.GroupListContainer.Objects = []fyne.CanvasObject{}

	// If there are no groups.
	if len(viewmodels.Groups.Models) == 0 {
		c.GroupListContainer.Add(widget.NewLabel("Create a group or accept an invite"))
		c.mu.Unlock()
		return
	}

	// If there are groups, sort them by whom the computer is member.
	var isMember []*models.Group
	var isNotMember []*models.Group
	for _, g := range viewmodels.Groups.Models {
		if viewmodels.GroupComputers.IsMemberOf(g) {
			isMember = append(isMember, g)
		} else {
			isNotMember = append(isNotMember, g)
		}
	}

	for _, g := range isMember {
		c.GroupListContainer.Add(NewEnterGroupCanvasObject(g))
	}

	if len(isNotMember) > 0 && len(isMember) > 0{
		separator := container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabelWithStyle("· · ·", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			layout.NewSpacer(),
		)
		c.GroupListContainer.Add(separator)
	}

	for _, g := range isNotMember {
		c.GroupListContainer.Add(NewJoinGroupCanvasObject(g))
	}

	c.mu.Unlock()
}

func (c *groupInviteListContent) DisplayGroupsContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.GroupRightContent}
	c.RightContent.Refresh()
}
func (c *groupInviteListContent) DisplayInvitesContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.InviteRightContent}
	c.RightContent.Refresh()
}

func NewInviteCanvasObject(invite *models.Invite) fyne.CanvasObject {
	return container.NewHBox(
		widget.NewLabel(invite.Group.Name+" @ "+invite.Date.String()),
		layout.NewSpacer(),
		widget.NewToolbar(widget.NewToolbarAction(resources.TrashDeleteSvg, func() {
			DeclineInvite(invite)
		}), widget.NewToolbarAction(resources.DoneCheckSvg, func() {
			AcceptInvite(invite)
		})))
}

func NewJoinGroupCanvasObject(group *models.Group) fyne.CanvasObject {
	return container.NewHBox(
		widget.NewLabel(group.Name),
		layout.NewSpacer(),
		widget.NewToolbar(widget.NewToolbarAction(resources.LoginSvg, func() {
			JoinGroup(group)
		})))
}

func NewEnterGroupCanvasObject(group *models.Group) fyne.CanvasObject {
	return container.NewHBox(
		widget.NewLabel(group.Name),
		layout.NewSpacer(),
		widget.NewToolbar(widget.NewToolbarAction(resources.SubdirectorySvg, func() {
			EnterGroup(group)
		})))
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
				createNewGroup(nameEntry.Text)
			}
		}, global.MainWindow).Show()
}

func EnterGroup(group *models.Group) {
	global.Log("entering group %v", group.GroupID)
}

func JoinGroup(group *models.Group) {
	global.Log("join group requested %v", group.GroupID)
}

func AcceptInvite(i *models.Invite) {
	// Accept the invite and add the group to the group view model.
	go func() {
		resp, err := server.Mnimidamon.Invite.AcceptCurrentUserInvite(&invite.AcceptCurrentUserInviteParams{
			GroupID:    i.Group.GroupID,
			Context:    server.ApiContext,
		}, server.CompAuth)

		if err != nil {
			infoDialog(err.Error())
			return
		}

		// Add the group and remove the invite.
		viewmodels.Groups.AddGroup(resp.Payload)
		viewmodels.Invites.RemoveInvite(i)
	}()
}

func DeclineInvite(i *models.Invite) {
	// Decline the invite and remove it from the invite view model.
	go func() {
		_, err := server.Mnimidamon.Invite.DeclineCurrentUserInvite(&invite.DeclineCurrentUserInviteParams{
			GroupID:    i.Group.GroupID,
			Context:    server.ApiContext,
		}, server.CompAuth)

		if err != nil {
			infoDialog(err.Error())
			return
		}

		// Remove the invite.
		viewmodels.Invites.RemoveInvite(i)
	}()
}

func infoDialog(msg string) {
	dialog.NewInformation("", msg, global.MainWindow).Show()
}

func createNewGroup(name string) {
	go func() {
		resp, err := server.Mnimidamon.Group.CreateGroup(&group.CreateGroupParams{
			Body:    &models.GroupCreatePayload{Name: &name},
			Context: server.ApiContext,
		}, server.CompAuth)

		if err != nil {
			if br, ok := err.(*group.CreateGroupBadRequest); ok {
				infoDialog(br.GetPayload().Message)
				return
			}
			infoDialog(err.Error())
			return
		}

		// Add the created group.
		viewmodels.Groups.AddGroup(resp.Payload)
	}()
}