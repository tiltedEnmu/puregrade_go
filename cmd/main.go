package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/repository"
	"github.com/ZaiPeeKann/puregrade/internal/service"
	g "github.com/ZaiPeeKann/puregrade/internal/transport/grpc"
	gh "github.com/ZaiPeeKann/puregrade/internal/transport/grpc/grpchandler"
	handler "github.com/ZaiPeeKann/puregrade/internal/transport/rest"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	viper.AddConfigPath("configs/example")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Config init error: %s", err.Error())
	}

	postgres, err := repository.NewPostgresDB(repository.PGConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		Username: viper.GetString("postgres.username"),
		Password: viper.GetString("postgres.password"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
	})

	redis := repository.NewRedisDB(repository.RedisConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
	})

	if err != nil {
		log.Fatalf("PostgreSQL database init error: %s", err.Error())
	}

	repos := repository.NewRepository(&repository.Databases{
		Postgres: postgres,
		Redis:    redis,
	})
	services := service.NewService(repos)
	handler := handler.NewHTTPHandler(services)
	srv := new(puregrade.Server)

	s := grpc.NewServer()
	grpchandler := g.NewGRPCServer(services)
	gh.RegisterAuthServer(s, grpchandler)

	go func() {
		if err := srv.Run(viper.GetString("httpport"), handler.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("HTTP server has been started")

	go func() {
		l, err := net.Listen("tcp", ":"+viper.GetString("grpcport"))
		if err != nil {
			log.Printf("Error occured while starting grpc server: %s", err.Error())
		}
		if err := s.Serve(l); err != nil {
			log.Printf("Error occured while running grpc server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err = srv.Shutdown(ctx); err != nil {
		log.Printf("Error occured on server shutting down: %s", err.Error())
	}

	if err = postgres.Close(); err != nil {
		log.Printf("Error occured on db shutting down: %s", err.Error())
	}

}
