package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

<<<<<<< HEAD
	server "github.com/SaRgEX/Diplom"
=======
	migrations "github.com/SaRgEX/Diplom/db"
	server "github.com/SaRgEX/Diplom/internal"
	"github.com/SaRgEX/Diplom/internal/config"
	"github.com/SaRgEX/Diplom/internal/storage/postgres"
>>>>>>> 3ca8a5dc40f54cedc3d6c5ac3e8dc0fb0a0b87fd
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

	fmt.Println(cfg)

	logrus.SetFormatter(new(logrus.JSONFormatter))
<<<<<<< HEAD
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	// migrations.MigrateSQL(db, "postgres")
=======
	logrus.Info("Starting server")
	logrus.Debug("Debug mode")
>>>>>>> 3ca8a5dc40f54cedc3d6c5ac3e8dc0fb0a0b87fd

	db, err := postgres.NewPostgresDB(cfg.Database)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	if err := migrations.MigrateSQL(db.DB, "postgres"); err != nil {
		logrus.Fatalf("failed to apply migrations: %s", err.Error())
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
