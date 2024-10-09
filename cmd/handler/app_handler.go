package handler

import (
	"net/http"

	"github.com/ShankaranarayananBR/movies-backend/models"
	"github.com/ShankaranarayananBR/movies-backend/requests"
	"github.com/labstack/echo/v4"
)

// HealthCheck function is used to show the status of the app
func (h *Handler) HealthCheck(c echo.Context) error {
	health_check_struct := struct {
		Health bool `json:"health"`
	}{
		Health: true,
	}
	return c.JSON(200, health_check_struct)
}

// AddMovies
func (h *Handler) AddMovies(c echo.Context) error {
	payload := new(requests.RegisterNewMovieRequest)
	err := (&echo.DefaultBinder{}).Bind(&payload, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payload)
	}
	n_movie := models.Movies{
		MovieName:   payload.MovieName,
		Description: payload.Description,
		MovieGenre:  payload.MovieGenre,
		MovieYear:   payload.MovieYear,
	}
	result := h.DB.Create(&n_movie)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Could not add movie")
	}
	return c.JSON(http.StatusOK, n_movie)
}

func (h *Handler) GetMovieByName(c echo.Context) error {
	movie_name := c.Param("movie_name")
	var movie_model models.Movies
	res := h.DB.Where("movie_name = ?", movie_name).Find(&movie_model)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "record not found")
	}
	return c.JSON(http.StatusOK, movie_model)
}

func (h *Handler) UpdateMovieDetails(c echo.Context) error {
	var movie_model models.Movies
	update := new(requests.UpdateMovieRequest)
	err := (&echo.DefaultBinder{}).Bind(update, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid update request")
	}
	err = h.GetMovieByName(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	movie_model.MovieYear = update.MovieYear
	result := h.DB.Save(&movie_model)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update movie year")
	}

	return c.JSON(http.StatusOK, "success")
}
