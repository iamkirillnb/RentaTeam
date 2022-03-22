package handlers

import (
	"github.com/iamkirillnb/Rentateam/internal"
	"github.com/iamkirillnb/Rentateam/internal/entities"
	"github.com/iamkirillnb/Rentateam/internal/repos"
	"github.com/iamkirillnb/Rentateam/pkg/logging"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

const (
	defaultServerReadTimeout  = 15 * time.Second
	defaultServerWriteTimeout = 30 * time.Second
)

type server struct {
	*http.Server

	repo repos.DbLaw
	log  logging.Logger
}

func NewHandler(config *internal.Config, repo repos.DbLaw, logger logging.Logger) (*server, error) {
	r := echo.New()
	serv := &http.Server{
		Handler:      echo.New(),
		Addr:         config.Server.Address(),
		ReadTimeout:  defaultServerReadTimeout,
		WriteTimeout: defaultServerWriteTimeout,
	}

	a := &server{
		Server: serv,
		repo:   repo,
		log:    logger,
	}

	r.GET("/", a.list)
	r.POST("/", a.add)

	log.Fatal(r.Start(config.Server.Address()))

	return a, nil
}

func (s *server) list(ctx echo.Context) error {
	data, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	return ctx.JSON(200, data)
}

func (s *server) add(ctx echo.Context) error {
	f := &entities.FormData{}
	_ = ctx.Bind(f)

	err := s.repo.WriteData(f)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, f)
}
