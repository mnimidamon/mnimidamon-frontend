package views

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"mnimidamonbackend/client/group"
	"mnimidamonbackend/client/group_computer"
	"mnimidamonbackend/client/invite"
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/frontend/global"
	"mnimidamonbackend/frontend/resources"
	"mnimidamonbackend/frontend/services"
	"mnimidamonbackend/frontend/views/fragments"
	"mnimidamonbackend/frontend/views/server"
	"mnimidamonbackend/frontend/views/viewmodels"
	"mnimidamonbackend/models"
	"sync"
)

func NewGroupAndInvitationsContent(processContainer fyne.CanvasObject) *groupsInvitationsContent {

	// Base toolbar elements.
	groupLabel := widget.NewLabelWithStyle("groups", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	inviteLabel := widget.NewLabelWithStyle("invites", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	computersLabel := widget.NewLabelWithStyle("computers", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	groupToolbarLabel := fragments.NewToolbarObject(groupLabel)
	inviteToolbarLabel := fragments.NewToolbarObject(inviteLabel)
	computersToolbarLabel := fragments.NewToolbarObject(computersLabel)

	// For reference.
	var gilc *groupsInvitationsContent

	// Group toolbar.
	groupToolbar := widget.NewToolbar(
		groupToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Groups.GetAllGroups()
			viewmodels.CurrentComputerGroupComputers.GetAllGroupComputers()
		}),
		widget.NewToolbarAction(resources.GroupAddSvg, func() {
			groupAddDialog()
		}),
	)

	// Invite Toolbar.
	inviteToolbar := widget.NewToolbar(
		inviteToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Invites.GetAllInvites()
		}),
	)

	// Invite Toolbar.
	computersToolbar := widget.NewToolbar(
		computersToolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Computers.GetAll()
		}),
	)

	// Left navigation.
	leftNavigation := container.NewVBox(
		widget.NewButtonWithIcon("groups", resources.GroupsSvg, func() {
			gilc.DisplayGroupsContent()
		}),
		widget.NewButtonWithIcon("invites", resources.MailboxUpSvg, func() {
			gilc.DisplayInvitesContent()
		}),
		widget.NewButtonWithIcon("computers", resources.ClipboardEditSvg, func() {
			gilc.DisplayComputersContent()
		}),
	)

	// Right list content.
	groupListContainer := container.NewVBox(NewItalicLabel("loading..."))
	inviteListContainer := container.NewVBox(NewItalicLabel("loading..."))
	computersListContainer := container.NewVBox(NewItalicLabel("loading..."))

	// The center content.
	rightContent := container.NewMax()

	// Border layout, where left is navigation and center content is the right content.
	mainContainer := container.NewBorder(nil, nil, leftNavigation, nil, rightContent)

	gilc = &groupsInvitationsContent{
		Container:      mainContainer,
		LeftNavigation: leftNavigation,
		RightContent:   rightContent,

		GroupRightContent:     container.NewBorder(groupToolbar, processContainer, nil, nil, container.NewVScroll(container.NewPadded(groupListContainer))),
		InviteRightContent:    container.NewBorder(inviteToolbar, processContainer, nil, nil, container.NewVScroll(container.NewPadded(inviteListContainer))),
		ComputersRightContent: container.NewBorder(computersToolbar, processContainer, nil, nil, container.NewVScroll(container.NewPadded(computersListContainer))),

		GroupListContainer:     groupListContainer,
		InviteListContainer:    inviteListContainer,
		ComputersListContainer: computersListContainer,

		mu: sync.Mutex{},
	}

	// Default content is groups.
	gilc.DisplayGroupsContent()

	// Register listeners.
	events.GroupsUpdated.Register(gilc)
	events.InvitesUpdated.Register(gilc)
	events.ComputersUpdated.Register(gilc)
	events.CurrentComputerGroupComputersUpdated.Register(gilc)

	// Resize it to look somewhat better.
	global.MainWindow.Resize(fyne.Size{
		Width:  350,
		Height: 220,
	})

	return gilc
}

type groupsInvitationsContent struct {
	Container      *fyne.Container // The encapsulating container.
	LeftNavigation *fyne.Container // Left split content.
	RightContent   *fyne.Container // Right split content.

	GroupRightContent     *fyne.Container // Content displayed upon Invite navigation.
	InviteRightContent    *fyne.Container // Content displayed upon Group navigation.
	ComputersRightContent *fyne.Container // Content displayed upon computers navigation.

	GroupListContainer     *fyne.Container // Containing the group list.
	InviteListContainer    *fyne.Container // Containing the invite group list.
	ComputersListContainer *fyne.Container // Containing the  computers list.

	mu sync.Mutex // Lock when rendering UI elements.
}

func (c *groupsInvitationsContent) HandleComputersUpdated() {
	c.rerenderComputers()
}

func (c *groupsInvitationsContent) HandleGroupComputersUpdated() {
	c.rerenderGroups()
}

func (c *groupsInvitationsContent) HandleGroupsUpdate() {
	c.rerenderGroups()
}

func (c *groupsInvitationsContent) HandleInvitesUpdate() {
	c.rerenderInvites()
}

func (c *groupsInvitationsContent) rerenderInvites() {
	c.mu.Lock()

	global.Log("updating invites list")
	c.InviteListContainer.Objects = []fyne.CanvasObject{}

	if len(viewmodels.Invites.Models) == 0 {
		c.InviteListContainer.Add(NewItalicLabel("You have no pending invites"))
		c.InviteListContainer.Refresh()
		c.mu.Unlock()
		return
	}

	for _, i := range viewmodels.Invites.Models {
		c.InviteListContainer.Add(NewInviteCanvasObject(i))
	}

	c.InviteListContainer.Refresh()
	c.mu.Unlock()
}

func (c *groupsInvitationsContent) rerenderGroups() {
	c.mu.Lock()
	global.Log("updating groups list")
	c.GroupListContainer.Objects = []fyne.CanvasObject{}

	// If there are no groups.
	if len(viewmodels.Groups.Models) == 0 {
		c.GroupListContainer.Add(NewItalicLabel("Create a group or accept an invite"))
		c.GroupListContainer.Refresh()
		c.mu.Unlock()
		return
	}

	// If there are groups, sort them by whom the computer is member.
	var isMember []*models.Group
	var isNotMember []*models.Group
	for _, g := range viewmodels.Groups.Models {
		if viewmodels.CurrentComputerGroupComputers.IsMemberOf(g) {
			isMember = append(isMember, g)
		} else {
			isNotMember = append(isNotMember, g)
		}
	}

	for _, g := range isMember {
		c.GroupListContainer.Add(NewEnterGroupCanvasObject(g))
	}

	if len(isNotMember) > 0 && len(isMember) > 0 {
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

	c.GroupListContainer.Refresh()
	c.mu.Unlock()
}

func (c *groupsInvitationsContent) rerenderComputers() {
	c.mu.Lock()
	global.Log("updating computers list")
	defer c.mu.Unlock()
	defer c.ComputersListContainer.Refresh()

	c.ComputersListContainer.Objects = []fyne.CanvasObject{}
}

func (c *groupsInvitationsContent) DisplayGroupsContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.GroupRightContent}
	c.RightContent.Refresh()
}
func (c *groupsInvitationsContent) DisplayInvitesContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.InviteRightContent}
	c.RightContent.Refresh()
}

func (c *groupsInvitationsContent) DisplayComputersContent() {
	c.RightContent.Objects = []fyne.CanvasObject{c.ComputersRightContent}
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
		widget.NewToolbar(widget.NewToolbarAction(resources.ClipboardEditSvg, func() {
			DialogJoinComputerToGroup(group)
		})))
}

func NewEnterGroupCanvasObject(group *models.Group) fyne.CanvasObject {
	return container.NewHBox(
		widget.NewLabel(group.Name),
		layout.NewSpacer(),
		widget.NewToolbar(widget.NewToolbarAction(resources.LoginSvg, func() {
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
	// Select the Selected group, this will emit an event which will cause main screen to represent the backups content.
	viewmodels.SelectedGroup.Select(group)
}

func DialogJoinComputerToGroup(group *models.Group) {
	// Dialog for entering the amount of space.
	global.Log("join group requested %v", group.GroupID)

	size := binding.NewInt()
	sizeEntry := widget.NewEntryWithData(binding.IntToString(size))
	dialog.NewForm(fmt.Sprintf("Join %v with %v", group.Name, services.ConfigurationStore.GetConfig().Computer.Name), "Join", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("Storage size in MB", sizeEntry),
		}, func(b bool) {
			if b {
				size, _ := size.Get()
				// Convert to KB
				JoinComputerToGroup(size*1024, group)
			}
		}, global.MainWindow).Show()
}

func JoinComputerToGroup(size int, group *models.Group) {
	go func() {
		resp, err := server.Mnimidamon.GroupComputer.JoinComputerToGroup(&group_computer.JoinComputerToGroupParams{
			Body: &models.CreateGroupComputerPayload{
				Size: int64(size),
			},
			GroupID: group.GroupID,
			Context: server.ApiContext,
		}, viewmodels.CurrentComputer.Auth)

		// Display the error.
		if err != nil {
			if br, ok := err.(*group_computer.JoinComputerToGroupBadRequest); ok {
				infoDialog(br.Payload.Message)
				return
			}
			infoDialog(err.Error())
			return
		}

		// Add the group computers to the view models.
		viewmodels.CurrentComputerGroupComputers.Add(resp.Payload)
	}()
}

func AcceptInvite(i *models.Invite) {
	// Accept the invite and add the group to the group view model.
	go func() {
		resp, err := server.Mnimidamon.Invite.AcceptCurrentUserInvite(&invite.AcceptCurrentUserInviteParams{
			GroupID: i.Group.GroupID,
			Context: server.ApiContext,
		}, viewmodels.CurrentComputer.Auth)

		if err != nil {
			infoDialog(err.Error())
			return
		}

		// Add the group and remove the invite.
		viewmodels.Groups.Add(resp.Payload)
		viewmodels.Invites.Remove(i)
	}()
}

func DeclineInvite(i *models.Invite) {
	// Decline the invite and remove it from the invite view model.
	go func() {
		_, err := server.Mnimidamon.Invite.DeclineCurrentUserInvite(&invite.DeclineCurrentUserInviteParams{
			GroupID: i.Group.GroupID,
			Context: server.ApiContext,
		}, viewmodels.CurrentComputer.Auth)

		if err != nil {
			infoDialog(err.Error())
			return
		}

		// Remove the invite.
		viewmodels.Invites.Remove(i)
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
		}, viewmodels.CurrentComputer.Auth)

		if err != nil {
			if br, ok := err.(*group.CreateGroupBadRequest); ok {
				infoDialog(br.GetPayload().Message)
				return
			}
			infoDialog(err.Error())
			return
		}

		// Add the created group.
		viewmodels.Groups.Add(resp.Payload)
	}()
}
