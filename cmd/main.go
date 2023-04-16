package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	server "github.com/ZaiPeeKann/auth-service_pg"
	"github.com/ZaiPeeKann/auth-service_pg/internal/repository"
	"github.com/ZaiPeeKann/auth-service_pg/internal/service"
	handler "github.com/ZaiPeeKann/auth-service_pg/internal/transport/rest"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("configs/example")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Config init error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("DB init error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHTTPHandler(services)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("Server has been started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("Error occured on server shutting down: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		log.Printf("Error occured on db shutting down: %s", err.Error())
	}
}
