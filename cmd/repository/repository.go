package repository

import (
	models "github.com/GoutamVerma/FamPay-Backend/utils"
)

// VideoRepository is an interface that defines the methods for video operations
type VideoRepository interface {
	AddVideo(model models.Video) error
	GetVideo() ([]models.Video, error)
	DeleteAllVideos() error
}
