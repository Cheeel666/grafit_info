package service

import (
	"projects/grafit_info/internal/database/mongodb/models"
	"projects/grafit_info/internal/database/mongodb/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CommonInfoService - service for work with common info such as headers, links, offers, etc.
type CommonInfoService struct {
	log            *zap.Logger
	CommonInfoRepo repository.CommonInfo
	DocumentDict   repository.DocumentDict
}

func NewCommonInfoService(
	log *zap.Logger,
	commonInfoRepo repository.CommonInfo,
	docRepo repository.DocumentDict,
) *CommonInfoService {
	return &CommonInfoService{
		log:            log,
		CommonInfoRepo: commonInfoRepo,
		DocumentDict:   docRepo,
	}
}

func (s *CommonInfoService) Get(ctx *gin.Context) {
	commonInfo, err := s.CommonInfoRepo.Find()
	if err != nil {
		s.log.Error("GetInfo", zap.String("err", err.Error()))
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, commonInfo)
}

// GetByType - get common info by type. Needet to get info by type from front. TODO: add search for latest in db
func (s *CommonInfoService) GetByType(ctx *gin.Context) {
	docName := ctx.Param("name")
	dictDoc, err := s.DocumentDict.FindByName(&models.FindReq{
		Name: docName,
	})
	if err != nil {
		s.log.Error("GetDoc with field", zap.Error(err))
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	doc, err := s.CommonInfoRepo.FindByType(&models.FindReq{
		Type: dictDoc.Type,
	})
	if err != nil {
		s.log.Error("GetInfo with field", zap.Error(err))
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, doc)
}

func (s *CommonInfoService) GetByID(ctx *gin.Context) {

}

func (s *CommonInfoService) Create(ctx *gin.Context) {

}

func (s *CommonInfoService) Update(ctx *gin.Context) {

}

func (s *CommonInfoService) Delete(ctx *gin.Context) {

}
