package services

import "github.com/murilosolino/challenge-backend-7/api/model"

type DestinationSvc struct {
	dm model.DestinationModel
}

func NewDestinationSvc(dm model.DestinationModel) *DestinationSvc {
	return &DestinationSvc{dm: dm}
}

func (svc *DestinationSvc) CreateDestination(d model.DestinationRow) error {
	return svc.dm.Save(d)
}

func (svc *DestinationSvc) ListDestinations() ([]model.DestinationRow, error) {
	return svc.dm.List()
}

func (svc *DestinationSvc) DeleteDestinationById(id int) error {
	return svc.dm.Exclude(id)
}
