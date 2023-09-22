package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/grafit_info/internal/database/mongodb/repository"
)

// CommonInfoService - service for work with common info such as headers, links, offers, etc.
type CommonInfoService struct {
	log            *zap.Logger
	CommonInfoRepo repository.CommonInfo
}

func NewCommonInfoService(log *zap.Logger, repo repository.CommonInfo) *CommonInfoService {
	return &CommonInfoService{
		log:            log,
		CommonInfoRepo: repo,
	}
}

func (s *CommonInfoService) Get(ctx *gin.Context) {

}

func (s *CommonInfoService) GetByID(ctx *gin.Context) {

}

func (s *CommonInfoService) Create(ctx *gin.Context) {

}

func (s *CommonInfoService) Update(ctx *gin.Context) {

}

func (s *CommonInfoService) Delete(ctx *gin.Context) {

}
