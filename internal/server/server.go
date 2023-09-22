package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"projects/grafit_info/config"
	mongodb "projects/grafit_info/internal/database/mongodb"
	"projects/grafit_info/internal/database/mongodb/repository"
	"projects/grafit_info/internal/service"
)

// Server is a struct that contains all server dependencies.
type Server struct {
	mongo mongodb.Client
	log   *zap.Logger
	app   *gin.Engine

	commonInfo *service.CommonInfoService
	tariff     *service.TariffService
}

// NewServer creates new server instance.
func NewServer(cfg *config.Config, log *zap.Logger, mongoClient *mongodb.Client) *Server {
	tariffRepository := repository.NewTariffRepository(cfg, mongoClient)
	commonInfoRepository := repository.NewCommonInfoRepo(cfg, mongoClient)

	commonInfo := service.NewCommonInfoService(log, commonInfoRepository)
	tariffService := service.NewTariffService(log, cfg, tariffRepository)

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

// initHandlers initializes handlers for server and use CORS middleware.
func (s *Server) initHandlers() {
	s.app.Use(
		CORSMiddleware(),
		gin.Logger(),
		gin.Recovery(),
	)

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

// CORSMiddleware for logging
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
