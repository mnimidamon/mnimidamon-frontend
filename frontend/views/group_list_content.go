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

func NewGroupListContent() *groupListContent {
	groupLabel := widget.NewLabelWithStyle("groups", fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})
	toolbarLabel := fragments.NewToolbarLabel(groupLabel)

	toolbar := widget.NewToolbar(toolbarLabel,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resources.SyncSvg, func() {
			viewmodels.Groups.GetAllGroups()
		}),
		widget.NewToolbarAction(resources.GroupAddSvg, func() {
			groupAddDialog()
		}),
	)

	listContainer := container.NewMax()
	mainContainer := container.NewVBox(toolbar, listContainer)

	groupListContent := &groupListContent{
		Container:     mainContainer,
		Toolbar:       toolbar,
		ListContainer: listContainer,
	}

	events.GroupsUpdated.Register(groupListContent)

	return groupListContent
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

type groupListContent struct {
	Container     *fyne.Container // The encapsulating container.
	Toolbar       *widget.Toolbar // The upper toolbar.
	ListContainer *fyne.Container // The list container containing groups.
}

func (c *groupListContent) HandleGroupsUpdate() {
	global.Log("group list content groups update")
}
