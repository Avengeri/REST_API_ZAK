package main

import (
	srv "Interface_droch_3"
	"Interface_droch_3/internal/handl"
	"Interface_droch_3/internal/repository"
	"Interface_droch_3/internal/repository/postgres"
	"Interface_droch_3/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

// @title REST_API_ZAK
// @version 0.0.1
// @description Программа для обучения RETS API

// @host localhost:8080
// @BasePath /
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка инициализации конфига: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки переменной окружения: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("Ошибка создания Postgres: %s", err.Error())
	}

	//rdb, err := redis_storage.NewRedisClient(redis_storage.Config{
	//	Addr: viper.GetString("rdb.address"),
	//})
	//
	//if err != nil {
	//	log.Fatalf("Ошибка создания Redis: %s", err.Error())
	//}

	//repo := repository.NewStorageUsersRedis(rdb)
	repo := repository.NewStorageUsersPostgres(db)
	services := service.NewServiceUsers(repo)
	handlers := handl.NewHandler(services)

	serv := new(srv.Server)
	if err = serv.Run(viper.GetString("srv_port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Не удалось запустить сервер: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
