package server

import (
	"burmese_jewellery/dependency"
	"burmese_jewellery/env"
	"burmese_jewellery/handler"
	"burmese_jewellery/server/middleware"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	gin_middleware "github.com/oapi-codegen/gin-middleware"
	"github.com/rs/zerolog/log"
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

	// Dependency
	if err := dependency.NewDependency(); err != nil {
		panic(err)
	}

	// Set Server
	router := gin.New()
	router.MaxMultipartMemory = 8 << 20 // 8 * 1MB(2 power 20) = 8MB
	router.Use(
		gin.RecoveryWithWriter(log.Logger),
		gin.LoggerWithConfig(gin.LoggerConfig{Output: log.Logger}),
	)

	routerConfig := cors.DefaultConfig()
	routerConfig.AllowOrigins = env.AllowOrigins()
	routerConfig.AddAllowHeaders(
		"Access-Control-Allow-Origin",
		"Authorization",
	)
	router.Use(cors.New(routerConfig))

	router.Use(middleware.Auth())
	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	// Swagger
	swagger, err := handler.GetSwagger()
	if err != nil {
		panic(err)
	}
	openapi3filter.RegisterBodyDecoder("image/png", openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder("image/jpeg", openapi3filter.FileBodyDecoder)
	router.Use(gin_middleware.OapiRequestValidatorWithOptions(swagger, &gin_middleware.Options{
		SilenceServersWarning: true,
	}))

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

	if env.Get().UseTLS {
		if err := s.srv.ListenAndServeTLS(
			certsFile,
			certsKeyFile,
		); err != nil && err != http.ErrServerClosed {
			return err
		}
	} else {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
	}

	return nil
}

func (s *server) Shutdown() {
	log.Info().Msg("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), serverTimeOut)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("")
	}

	log.Info().Msg("Server terminated")
}
