package server

import (
	"fmt"
	"github.com/journeymidnight/log"
	"time"
	"twins/task"
)

var AsElder bool

type TwinsServer struct {
	InitAsElder bool
	RpcServer   *TwinsRpcServer
	RpcClient   *TwinsRpcClient
	SwitchCh    chan bool
	Logger      log.Logger
}

func NewTwinsServer(confPath string) *TwinsServer {
	conf, err := readConfig(confPath)
	if err != nil {
		panic(err)
	}
	logger := log.NewFileLogger(conf.LogPath, log.ParseLevel(conf.LogLevel))
	logger.Info(fmt.Sprintf("config %+v", conf))

	rpcClient, err := NewTwinsClient(conf.TwinsBind, 5*time.Second, logger)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	rpcServer := NewTwinsRpcServer(conf.Bind, logger)
	s := &TwinsServer{
		InitAsElder: conf.InitAsElder,
		RpcServer:   rpcServer,
		RpcClient:   rpcClient,
		SwitchCh:    make(chan bool),
		Logger:      logger,
	}
	task.TaskConfDir = conf.TaskConfDir
	task.TaskLogDir = conf.TaskLogDir
	return s
}

func (s *TwinsServer) Run() error {
	go s.runSwitcher()
	if s.InitAsElder {
		go s.runAsElder()
	} else {
		go s.runAsLittle()
	}

	err := s.RpcServer.Run()
	if err != nil {
		return err
	}
	return nil
}
