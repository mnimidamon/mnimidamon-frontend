package viewmodels

import (
	"mnimidamonbackend/frontend/events"
	"mnimidamonbackend/models"
)

var SelectedGroup *selectedGroupViewModel

func init() {
	SelectedGroup = &selectedGroupViewModel{
		Model: nil,
	}
}

type selectedGroupViewModel struct {
	Model *models.Group
}

func (vm *selectedGroupViewModel) Select(group *models.Group)  {
	vm.Model = group
	events.SelectedGroupUpdated.Trigger()
}