package handlers

import (
	"net/http"

	"github.com/GoutamVerma/FamPay-Backend/cmd/repository/sql"
	"github.com/GoutamVerma/FamPay-Backend/cmd/usecase"
	"github.com/labstack/echo"
)

// RegisterHandlers registers the HTTP handlers for the routes
func RegisterHandlers(e *echo.Echo) {
	e.GET("/v1/getVideos", GetVideo)
	e.GET("/v1/deleteVideos", DeleteAllVideos)
	go usecase.StartYouTubeAPICall("gamming", 5)
}

// GetVideo handles the GET /v1/getVideos route
func GetVideo(c echo.Context) error {
	data, err := sql.GetVideos()
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, data)
}

// DeleteAllVideos handles the GET /v1/deleteVideo route
func DeleteAllVideos(c echo.Context) error {
	sql.DeleteAllVideos()
	return c.JSON(http.StatusOK, "Deleted")
}
