package http

import (
	"github.com/iamkirillnb/Rentateam/pkg/logging"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServerRouter interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

type Context = echo.Context

type Route = echo.Route

type Group = echo.Group

type ServerHandlerFunc = echo.HandlerFunc

type ServerMiddlewareFunc = echo.MiddlewareFunc

type ServerOption func(s *Server) error

type Server struct {
	cfg *ServerConfig

	baseRouter *echo.Echo
	baseServer *http.Server

	log logging.Logger
}

func NewServer(logger logging.Logger, opts ...ServerOption) (*Server, error) {

	defaultCfg := &ServerConfig{}


	defaultRouter := echo.New()

	srv := &Server{
		cfg:        defaultCfg,
		baseRouter: defaultRouter,
		baseServer: nil,
		log:        logger,
	}

	for _, opt := range opts {
		if err := opt(srv); err != nil {
			return nil, err
		}
	}

	srv.baseServer = &http.Server{
		Addr:         srv.cfg.Address(),
		ReadTimeout:  srv.cfg.ReadTimeout,
		WriteTimeout: srv.cfg.WriteTimeout,
	}

	return srv, nil
}

func (s *Server) Start() error {
	s.log.Infof("starting on %s", s.cfg.Address())

	err := s.baseServer.ListenAndServe()
	if err != nil {
		return err
	}
	s.log.Infof("started on %s", s.cfg.Address())

	return nil
}

func (s *Server) GET(path string, h ServerHandlerFunc) *Route {
	return s.baseRouter.GET(path, h)
}

