package main

import (
	"crud-go/config"
	"crud-go/controller"
	"crud-go/helper"
	"crud-go/model"
	"crud-go/repository"
	"crud-go/router"
	"crud-go/service"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	//Init Repository
	tagRepository := repository.NewTagsRepositoryImpl(db)

	//Init Service
	tagService := service.NewTagServiceImpl(tagRepository, validate)

	//Init controller
	tagController := controller.NewTagController(tagService)

	//Router
	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
