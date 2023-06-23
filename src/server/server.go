package server

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/HuckOps/notify/pkg/RSA"
	"github.com/HuckOps/notify/src/config"
	_ "github.com/HuckOps/notify/src/server/docs"
	"github.com/HuckOps/notify/src/server/route"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"net/http"
	"os"
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
	context            context.Context
	httpServer         *http.Server
	restartChan        chan bool
	PasswordPrivateKey *rsa.PrivateKey
}

type Server interface {
	Listen()
	Restart()
	Kill()
}

func NewServer(privateKeyPath string) Server {
	priKey, err := RSA.ReadRSAPKCS1PrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	return &server{
		context:            context.Background(),
		httpServer:         &http.Server{},
		restartChan:        make(chan bool, 1),
		PasswordPrivateKey: priKey,
	}
}

//go:generate swag init --generalInfo=route/base.go
func (s *server) Listen() {
	for {
		// 获取新的服务实体

		select {
		case <-s.restartChan:
			s.httpServer.Shutdown(s.context)
			fmt.Println("exit")
		}
		gin.ForceConsoleColor()
		gin.DefaultWriter = io.MultiWriter(os.Stdout)
		e := gin.Default()

		route.SetPrivateKeyToContext(e, s.PasswordPrivateKey)
		route.Handler(e)
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		s.httpServer = &http.Server{Handler: e, Addr: fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port)}
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

func (s *server) Kill() {
	if err := s.httpServer.Shutdown(s.context); err != nil {
		panic(err)
	}
	fmt.Println("Kill server")
}
