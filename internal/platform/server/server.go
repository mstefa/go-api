package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	mooc "github.com/mstefa/go-api/internal"
	"github.com/mstefa/go-api/internal/platform/handler/courses"
	"github.com/mstefa/go-api/internal/platform/handler/health"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		courseRepository: courseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}
