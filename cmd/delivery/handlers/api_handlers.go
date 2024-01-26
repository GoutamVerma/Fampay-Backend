package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// RegisterHandlers registers the HTTP handlers for the routes
func RegisterHandlers(e *echo.Echo) {
	e.GET("/v1/getVideos", GetVideo)
	e.GET("/v1/deleteVideos", DeleteAllVideos)
}

// GetVideo handles the GET /v1/getVideos route
func GetVideo(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get videos")
}

// DeleteAllVideos handles the GET /v1/deleteVideo route
func DeleteAllVideos(c echo.Context) error {
	return c.JSON(http.StatusOK, "Deleted")
}
