package sql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GoutamVerma/FamPay-Backend/config"
	models "github.com/GoutamVerma/FamPay-Backend/utils"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// Read database configuration from config package
	username := config.ReadUserName()
	password := config.ReadPassWord()
	hostname := config.ReadHostName()
	port := config.ReadPort()
	databaseName := config.ReadDatabaseName()

	// Create the Data Source Name (DSN) string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, databaseName)

	// Open database connection
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Create the database if it does not exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName)
	if err != nil {
		panic(err)
	}

	// Use the database
	_, err = db.Exec("USE " + databaseName)
	if err != nil {
		panic(err)
	}

	// Create the youtubeVideos table if it does not exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS youtubeVideos (
		Title VARCHAR(255),
		Descripition VARCHAR(255),
		Thumbnails VARCHAR(255),
		PublishedAt  TIMESTAMP
	)`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully!")
}

// AddVideos inserts a new video into the youtubeVideos table
func AddVideos(model models.Video) error {
	_, err := db.Exec("INSERT INTO youtubeVideos (Title, Descripition, Thumbnails, PublishedAt) VALUES (?, ?, ?, ?)", model.Title, model.Description, model.Thumbnails, model.PublishedAt)
	if err != nil {
		return err
	}
	fmt.Println("Record inserted successfully")
	return nil
}

// GetVideos retrieves all videos from the youtubeVideos table
func GetVideos(page int) ([]models.Video, error) {
	pageSize := config.PageSize()
	var videos []models.Video
	query := "SELECT * FROM youtubeVideos LIMIT ? OFFSET ?"
	rows, err := db.Query(query, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Title string
		var Description string
		var Thumbnails string
		var PublishedAt string
		err := rows.Scan(&Title, &Description, &Thumbnails, &PublishedAt)
		if err != nil {
			return nil, nil
		}
		addedAt, err := time.Parse("2006-01-02 15:04:05", PublishedAt)
		if err != nil {
			return nil, nil
		}
		video := models.Video{
			Title:       Title,
			Description: Description,
			Thumbnails:  Thumbnails,
			PublishedAt: addedAt,
		}
		videos = append(videos, video)
	}

	return videos, nil
}

// DeleteAllVideos deletes all videos from the youtubeVideos table
func DeleteAllVideos() error {
	_, err := db.Exec("DELETE FROM youtubeVideos")
	if err != nil {
		return err
	}
	fmt.Println("All Videos deleted successfully")
	return nil
}

// GetTotalRecords retrieves the total number of records in the youtubeVideos table
func GetTotalRecords() (int, error) {
	var totalRecords int

	// Execute a COUNT query to get the total number of records
	err := db.QueryRow("SELECT COUNT(*) FROM youtubeVideos").Scan(&totalRecords)
	if err != nil {
		return 0, err
	}

	return totalRecords, nil
}
