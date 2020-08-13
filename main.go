package main

import (
	"ab_division/golbal"
	"ab_division/log"
	"ab_division/rpc/proto"
	"ab_division/util"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/robfig/cron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func init() {
	var err error
	log.DataLogger, err = seelog.LoggerFromConfigAsFile("seelog_request.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = golbal.LoadConfig("prod.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func main() {
	conf := golbal.GetConfig()
	lis, err := net.Listen("tcp", conf.Port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	util.LoadExperiments()
	s := grpc.NewServer()
	proto.RegisterDivisionRPCServer(s, &util.DivisionServer{})

	//方便使用grpcurl工具调试，不使用时可关闭
	reflection.Register(s)

	c := cron.New()
	c.AddFunc("* */10 * * * *", func() {
		util.LoadExperiments()
	})
	c.Start()
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
