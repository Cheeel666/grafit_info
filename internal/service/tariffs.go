package service

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"projects/grafit_info/internal/database/mongodb"
	"projects/grafit_info/internal/database/mongodb/models"
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
	var tariffs models.Tariffs

	collection := s.mongoClient.Client.Database("grafit").Collection("tariffs", nil)
	filter := bson.D{}

	// TODO: time to config
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	cur, err := collection.Find(timeoutCtx, filter)
	if cur == nil {
		s.log.Error("Cursor is nil", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if err != nil {
		s.log.Error("Failed to get tariffs", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// TODO: time to config
	timeoutCtx, cancelFunc = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	if err := cur.All(timeoutCtx, &tariffs); err != nil {
		s.log.Error("Failed to decode tariffs", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, tariffs)
}

func (s *TariffService) GetByID(ctx *gin.Context) {

}

func (s *TariffService) Create(ctx *gin.Context) {

}

func (s *TariffService) Update(ctx *gin.Context) {

}

func (s *TariffService) Delete(ctx *gin.Context) {

}
