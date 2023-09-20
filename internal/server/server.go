package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/grafit_info/config"
	mongodb "projects/grafit_info/internal/database/mongodb"
	"projects/grafit_info/internal/service"
)

type Server struct {
	mongo mongodb.Client
	log   *zap.Logger
	app   *gin.Engine

	commonInfo *service.CommonInfoService
	tariff     *service.TariffService
}

func NewServer(cfg *config.Config, log *zap.Logger) *Server {
	mongoClient := mongodb.NewMongoDB(cfg, log)

	commonInfo := service.NewCommonInfoService(log, mongoClient)
	tariffService := service.NewTariffService(log, mongoClient)

	var s = &Server{
		mongo: *mongoClient,
		log:   log,
		app:   gin.Default(),

		commonInfo: commonInfo,
		tariff:     tariffService,
	}

	s.initHandlers()

	return s
}

func (s *Server) initHandlers() {
	tariffs := s.app.Group("/tariffs")
	{
		tariffs.GET("/", s.tariff.Get)
		tariffs.GET("/:id", s.tariff.GetByID)
		tariffs.POST("/", s.tariff.Create)
		tariffs.PUT("/:id", s.tariff.Update)
		tariffs.DELETE("/:id", s.tariff.Delete)
	}

	commonInfo := s.app.Group("/common_info")
	{
		commonInfo.GET("/", s.commonInfo.Get)
		commonInfo.POST("/", s.commonInfo.Create)
		commonInfo.PUT("/:id", s.commonInfo.Update)
		commonInfo.DELETE("/:id", s.commonInfo.Delete)
	}

	//pages := s.app.Group("/pages")
	//{
	//	pages.GET("/", s.getPages)
	//	pages.GET("/:id", s.getPage)
	//	pages.POST("/", s.createPage)
	//	pages.PUT("/:id", s.updatePage)
	//	pages.DELETE("/:id", s.deletePage)
	//}
	//
	//trainers := s.app.Group("/trainers")
	//{
	//	trainers.GET("/", s.getTrainers)
	//	trainers.GET("/:id", s.getTrainer)
	//	trainers.POST("/", s.createTrainer)
	//	trainers.PUT("/:id", s.updateTrainer)
	//	trainers.DELETE("/:id", s.deleteTrainer)
	//}
}

func (s *Server) Run() error {
	return s.app.Run()
}
