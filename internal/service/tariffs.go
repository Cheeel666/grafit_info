package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/grafit_info/config"
	"projects/grafit_info/internal/database/mongodb/models"
	"projects/grafit_info/internal/database/mongodb/repository"
)

type TariffService struct {
	log        *zap.Logger
	TariffRepo repository.Tariff
	cfg        *config.Config
}

func NewTariffService(log *zap.Logger, cfg *config.Config, repo repository.Tariff) *TariffService {
	return &TariffService{
		log:        log,
		TariffRepo: repo,
		cfg:        cfg,
	}
}

func (s *TariffService) Get(ctx *gin.Context) {
	tariffs, err := s.TariffRepo.Find()
	if err != nil {
		s.log.Error("GetTariffs", zap.String("err", err.Error()))
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, tariffs)
}

func (s *TariffService) GetByID(ctx *gin.Context) {
	tariff, err := s.TariffRepo.FindByID(&models.FindReq{
		ID: ctx.Param("id"),
	})
	if err != nil {
		s.log.Error("GetTariff", zap.Error(err))
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, tariff)

}

func (s *TariffService) Create(ctx *gin.Context) {
	var tariff *models.Tariff
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(jsonData, &tariff)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.TariffRepo.Create(tariff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tariff created successfully!"})
}

func (s *TariffService) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	var tariff *models.Tariff
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(jsonData, &tariff)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.TariffRepo.Update(models.FindReq{
		ID: id,
	},
		tariff,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tariff updated successfully!"})
}

func (s *TariffService) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	err := s.TariffRepo.Delete(models.FindReq{
		ID: id,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tariff deleted successfully!"})
}
