package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-sql-driver/mysql" //lint:ignore ST1019 this is just example
	_ "github.com/go-sql-driver/mysql"
	"github.com/szpp-dev-team/hands-on-todo-app/api/grpc_server"
	"github.com/szpp-dev-team/hands-on-todo-app/config"
	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent"
	"golang.org/x/exp/slog"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// mysql client
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConfig := &mysql.Config{
		DBName:    conf.DBName,
		User:      conf.DBUser,
		Passwd:    conf.DBPass,
		Addr:      conf.DBAddr,
		Net:       "tcp",
		ParseTime: true,
		Loc:       jst,
	}
	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer entClient.Close()
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	// logger
	logger := slog.Default()

	srv := grpc_server.New(
		grpc_server.WithLogger(logger),
		grpc_server.WithEntClient(entClient),
	)
	lsnr, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()
	go func() {
		logger.Info("server launched")
		if err := srv.Serve(lsnr); err != nil {
			log.Fatal(err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	logger.Info("server is being stopped")
	srv.GracefulStop()
}
