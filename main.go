package main

import (
	"context"
	"os"

	"github.com/ShankaranarayananBR/movies-backend/cmd/handler"
	"github.com/ShankaranarayananBR/movies-backend/cmd/handler/middleware"
	"github.com/ShankaranarayananBR/movies-backend/database"
	"github.com/ShankaranarayananBR/movies-backend/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func init() {
	// Loading the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file:%v", err)
		return
	}
}

func main() {
	e := echo.New()
	newmovie := models.Movies{}
	db, err := database.MySQLConn(context.TODO(), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Printf("Error while connecting to DB:%v", err)
	}
	// AutoMigrate
	db.AutoMigrate(&newmovie)
	// Declaration of Handler
	h := handler.Handler{
		DB: db,
	}
	// checking health
	e.GET("/", h.HealthCheck, middleware.CustomMiddleware)
	e.POST("/newmovies", h.AddMovies, middleware.CustomMiddleware)
	e.GET("/getmovie/movie_name/:movie_name", h.GetMovieByName, middleware.CustomMiddleware)
	e.PUT("/updatemovie/movie_name/:movie_name", h.UpdateMovieDetails, middleware.CustomMiddleware)
	e.Logger.Fatal(e.Start(":3000"))
}
