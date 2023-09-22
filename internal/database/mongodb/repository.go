package mongodb

import "projects/grafit_info/internal/database/mongodb/models"

type Repo interface {
	Find() (models.Tariffs, error)
	FindByID(req *models.FindByIdReq) (*models.Tariff, error)
	Delete(req models.FindByIdReq) error
	Update(models.FindByIdReq, *models.Tariff) error
	Create(*models.Tariff) error
}
