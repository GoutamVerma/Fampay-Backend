package handlers

import (
	"net/http"
	"strconv"

	"github.com/GoutamVerma/FamPay-Backend/cmd/repository/sql"
	"github.com/GoutamVerma/FamPay-Backend/cmd/usecase"
	"github.com/labstack/echo"
)

// RegisterHandlers registers the HTTP handlers for the routes
func RegisterHandlers(e *echo.Echo) {
	e.GET("/v1/getVideos", GetVideo)
	e.GET("/v1/deleteVideos", DeleteAllVideos)
	go usecase.StartYouTubeAPICall()
}

// GetVideo handles the GET /v1/getVideos route
func GetVideo(c echo.Context) error {
	// Extract page number from query parameters
	page, err := strconv.Atoi(c.QueryParam("pageno"))
	if err != nil || page < 1 {
		page = 1
	}

	// Fetch videos for the specified page
	data, err := sql.GetVideos(page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch videos"})
	}

	return c.JSON(http.StatusOK, data)
}

// DeleteAllVideos handles the GET /v1/deleteVideo route
func DeleteAllVideos(c echo.Context) error {
	sql.DeleteAllVideos()
	return c.JSON(http.StatusOK, "Deleted")
}
