package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	parsetorrentname "github.com/middelink/go-parse-torrent-name"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&MovieList{})

	downloadsFolder := os.Getenv("DOWNLOADS_FOLDER")
	fmt.Println("Reading from folder ", downloadsFolder)

	t := NewMovieClient()

	tvFolder := os.Getenv("TV_FOLDER")
	movieFolder := os.Getenv("MOVIE_FOLDER")
	files := GetDirectoryFiles(downloadsFolder)
	for _, f := range files {

		info, _ := parsetorrentname.Parse(f)
		log.Println("Parser found: ", info.Title)
		multi := t.GetMediaTypeByMovieTitle(info.Title)
		mediaType := "unknown"
		if len(multi.Results) > 0 {
			mediaType = multi.Results[0].MediaType
			log.Println("Checking fileType:", mediaType, " Title:", info.Title)
		} else {
			fmt.Println("Video file is unknown. Name used: ", info.Title)
			continue
		}

		var desPath string
		srcPath := fmt.Sprintf("%s/%s", downloadsFolder, f)
		if mediaType == "tv" {
			desPath = fmt.Sprintf("%s/%s", tvFolder, f)
		} else if mediaType == "movie" {
			desPath = fmt.Sprintf("%s/%s", movieFolder, f)
		}
		fmt.Println("FileFound:", srcPath)
		err := CopyFile(srcPath, desPath, 1000)
		if err != nil {
			log.Println("Error: ", err)
		}

	}

}
