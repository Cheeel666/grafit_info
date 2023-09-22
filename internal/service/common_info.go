package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/grafit_info/internal/database/mongodb"
)

// CommonInfoService - service for work with common info such as headers, links, offers, etc.
type CommonInfoService struct {
	log         *zap.Logger
	mongoClient *mongodb.Client
}

func NewCommonInfoService(log *zap.Logger, mongo *mongodb.Client) *CommonInfoService {
	return &CommonInfoService{
		log:         log,
		mongoClient: mongo,
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
