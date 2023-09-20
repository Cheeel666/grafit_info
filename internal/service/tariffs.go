package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/grafit_info/internal/database/mongodb"
)

type TariffService struct {
	log         *zap.Logger
	mongoClient *mongodb.Client
}

func NewTariffService(log *zap.Logger, mongo *mongodb.Client) *TariffService {
	return &TariffService{
		log:         log,
		mongoClient: mongo,
	}
}

func (s *TariffService) Get(ctx *gin.Context) {

}

func (s *TariffService) GetByID(ctx *gin.Context) {

}

func (s *TariffService) Create(ctx *gin.Context) {

}

func (s *TariffService) Update(ctx *gin.Context) {

}

func (s *TariffService) Delete(ctx *gin.Context) {

}
