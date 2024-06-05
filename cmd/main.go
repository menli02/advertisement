package main

import (
	"Advertisement"
	"Advertisement/configs"
	"Advertisement/internal/handler"
	"Advertisement/internal/repository/mysql"
	"Advertisement/internal/repository/redis"
	"Advertisement/internal/service"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := configs.NewMysqlDB(configs.MySQLConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Database: viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initializing to db: %s", err.Error())
	}
	rdb, err := configs.NewRedisCache(configs.Options{
		Address:  viper.GetString("cache.address"),
		Password: viper.GetString("cache.password"),
		DB:       viper.GetInt("cache.db"),
	})
	mysqlRepo := mysql.NewAdvertisementMysql(db)
	redisCache := redis.NewAdvertisementRedis(rdb)
	services := service.NewService(mysqlRepo, redisCache)
	handlers := handler.NewHandler(services)
	srv := new(Advertisement.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Advertisement Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
