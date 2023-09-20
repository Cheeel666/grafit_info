package mongodb

import "projects/grafit_info/internal/database/mongodb/models"

type Repo interface {
	Find() (models.Tariffs, error)
	FindByID(req *models.TariffRequest) (*models.Tariff, error)
	Delete(req models.TariffRequest) error
	Update(models.TariffRequest, *models.Tariff) error
	Create(*models.Tariff) error
}
