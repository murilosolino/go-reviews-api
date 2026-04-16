package services

import "github.com/murilosolino/challenge-backend-7/api/model"

type DestinationSvc struct {
	dm model.DestinationModel
}

func NewDestinationSvc(dm model.DestinationModel) *DestinationSvc {
	return &DestinationSvc{dm: dm}
}

func (svc *DestinationSvc) CreateDestination(m map[string]interface{}) error {
	return svc.dm.Bm.Save(m)
}

func (svc *DestinationSvc) ListDestinations() ([]model.DestinationRow, error) {
	return svc.dm.ListAllDestinations()
}

func (svc *DestinationSvc) DeleteDestinationById(id int) error {
	return svc.dm.Bm.Exclude(id)
}

func (svc *DestinationSvc) UpdateDestination(id int, m map[string]interface{}) error {
	return svc.dm.Bm.Update(id, m)
}
