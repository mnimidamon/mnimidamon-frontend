package viewmodels

import (
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/models"
)

var SelectedGroup *selectedGroupViewModel

func init() {
	SelectedGroup = &selectedGroupViewModel{
		Group: nil,
	}
}

type selectedGroupViewModel struct {
	Group *models.Group
}

func (vm *selectedGroupViewModel) Select(group *models.Group)  {
	vm.Group = group
	events.SelectedGroupUpdated.Trigger()
}