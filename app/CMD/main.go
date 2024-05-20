package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"

	server "github.com/SaRgEX/Diplom/internal"
	"github.com/SaRgEX/Diplom/internal/config"
	"github.com/SaRgEX/Diplom/internal/storage/postgres"
	"github.com/SaRgEX/Diplom/pkg/handler"
	"github.com/SaRgEX/Diplom/pkg/repository"
	"github.com/SaRgEX/Diplom/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title           OTK API
// @version         1.0
// @description     API Server for OTK

// @host      localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.MustLoad()

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Info("Starting server")
	logrus.Debug("Debug mode")
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			filename = fmt.Sprintf("%s:%d", filename, f.Line)
			return "", filename
		},
	})

	db, err := postgres.NewPostgresDB(cfg.Database)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db.DB)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.HTTPServer, handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.DB.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
