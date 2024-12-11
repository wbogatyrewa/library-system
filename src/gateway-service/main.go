package main

import (
	"lab2/src/gateway-service/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	handler := handler.NewHandler()

	router := gin.Default()

	router.Use(cors.Default())

	// общие методы, для пользователя
	router.GET("/api/v1/libraries", handler.GetLibrariesByCity)               // получить список библиотек
	router.GET("/api/v1/libraries/:uid/books/", handler.GetBooksByLibraryUid) // получить список книг выбранной библиотеки
	router.POST("/api/v1/reservations", handler.CreateReservation)            // забронировать книгу в библиотеке

	// приватные методы, для библиотекаря
	router.GET("/api/v1/reservations", handler.GetReservations)         // получить список забронированных книг пользователя
	router.POST("/api/v1/reservations/:uid/return", handler.ReturnBook) // получить книгу от пользователя, оценив ее состояние
	router.GET("/api/v1/rating/", handler.GetRating)                    // получить рейтинг пользователя

	// сервисные методы
	router.GET("/manage/health", handler.GetHealth)

	router.Run()
}
