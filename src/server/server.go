package server

import (
	"context"
	"fmt"
	"github.com/HuckOps/notify/src/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//func Init() {
//
//}
//
//func init() {
//	se := gin.Default()
//	httpServer.Addr =
//		se.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
//	ServerEngine = se
//}

type server struct {
	context     context.Context
	httpServer  *http.Server
	restartChan chan bool
}

type Server interface {
	Listen()
	Restart()
}

func NewServer() Server {
	return &server{
		context:     context.Background(),
		httpServer:  &http.Server{},
		restartChan: make(chan bool, 1),
	}
}

func (s *server) Listen() {
	for {
		// 获取新的服务实体

		select {
		case <-s.restartChan:
			err := s.httpServer.Shutdown(s.context)
			fmt.Println(err)
		}
		se := gin.Default()
		s.httpServer = &http.Server{Handler: se, Addr: fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port)}

		go func() {
			fmt.Printf("Server listen on %s:%d\n", config.Config.Server.Host, config.Config.Server.Port)
			if err := s.httpServer.ListenAndServe(); err != nil {
				fmt.Printf("Server restart on %s\n", time.Now().Format("2006/01/02 15:04:05"))
			}
		}()
	}
}

func (s *server) Restart() {
	s.restartChan <- true
}
