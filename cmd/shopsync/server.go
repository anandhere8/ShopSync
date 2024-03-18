package server

import (
	"fmt"
	"reflect"

	db "github.com/anandhere8/ShopSync/db/sqlc"
	"github.com/anandhere8/ShopSync/internal/app/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	queries *db.Queries
}

func NewServer(querie *db.Queries) Server {
	fmt.Println(reflect.TypeOf(querie))
	nServer := Server{queries: querie}
	router := gin.Default()
	nServer.router = router
	handler.ConfigureRoutes(router)
	return nServer
}

func (s *Server) Start(address string) error {
	err := s.router.Run(address)
	if err != nil {
		return err
	}
	return nil
}
