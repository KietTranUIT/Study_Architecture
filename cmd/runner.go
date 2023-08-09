package main

import (
	"log"
	"user-service/internal/controller"
	_ "user-service/internal/core/model/request"
	"user-service/internal/core/service"
	"user-service/internal/infra/repository"
	"user-service/internal/infra/repository/config"
	"user-service/internal/server"
)

func main() {
	conf := config.ConfigDatabase{
		Driver:   "mysql",
		Username: "kiettran",
		Password: "Kiet@123456",
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "TEST",
	}
	db, err := repository.NewDB(conf)

	if err != nil {
		panic(err.Error())
	}

	log.Println("Connect to database successfully!")

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userControll := controller.NewUserControll(userService)

	userControll.Router()

	var server server.HTTPServer
	server.NewHTTPServer(userControll.Mux, 8585, "127.0.0.1")

	server.Start()
}
