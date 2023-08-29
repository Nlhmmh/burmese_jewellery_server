package server

import (
	"burmese_jewellery/env"
	"burmese_jewellery/handler"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const (
	serverTimeOut time.Duration = 5 * time.Second
)

type server struct {
	srv *http.Server
}

func NewServer() *server {

	// Logging
	newLog(true)

	// Open Postgres DB
	db, err := newDB(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.Get().DB.Host,
		env.Get().DB.Port,
		env.Get().DB.User,
		env.Get().DB.Password,
		env.Get().DB.Database,
	))
	if err != nil {
		panic(err)
	}

	// Boiler
	boil.SetDB(db)
	boil.DebugMode = true
	boil.DebugWriter = log.Logger

	// Set Server
	if !env.Get().Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(
		gin.RecoveryWithWriter(log.Logger),
		gin.LoggerWithConfig(gin.LoggerConfig{Output: log.Logger}),
	)
	router.SetTrustedProxies(nil)
	routerConfig := cors.DefaultConfig()
	routerConfig.AllowOrigins = env.Get().AllowOrigins
	// config.AllowHeaders = []string{
	// 	"Access-Control-Allow-Headers",
	// 	"Content-Type",
	// 	"Content-Length",
	// 	"Accept-Encoding",
	// 	"X-CSRF-Token",
	// 	"Authorization",
	// }
	router.Use(cors.New(routerConfig))

	// TODO: Comment out swagger because of openAPI check error
	// Swagger
	// swagger, err := handler.GetSwagger()
	// if err != nil {
	// 	panic(err)
	// }
	// router.Use(oapi_middleware.OapiRequestValidatorWithOptions(swagger, &oapi_middleware.Options{
	// 	SilenceServersWarning: true,
	// }))

	// Handler
	handler.RegisterHandlers(router, handler.NewHandler())

	return &server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", env.Get().Http.Port),
			Handler: router,
		},
	}

}

func (s *server) Run() error {

	log.Info().Msgf("Server listening on %s", s.srv.Addr)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	// if err := s.srv.ListenAndServeTLS(
	// 	utils.CertsFile,
	// 	utils.CertsKeyFile,
	// ); err != nil && err != http.ErrServerClosed {
	// 	return err
	// }

	return nil

}

func (s *server) Shutdown() {

	log.Info().Msg("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), serverTimeOut)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("error")
	}

	log.Info().Msg("Server terminated")

}
