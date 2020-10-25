package main

import (
	"fmt"
	"log"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
)

type MovieSearch struct {
	tmdb *tmdb.Client
	key  string
}

func NewMovieClient() *MovieSearch {
	m := MovieSearch{
		key: os.Getenv("API_KEY"),
	}
	tmdbClient, err := tmdb.Init(m.key)
	if err != nil {
		fmt.Println(err)
	}
	tmdbClient.SetClientAutoRetry()
	m.tmdb = tmdbClient

	return &m
}

func (t *MovieSearch) GetMediaTypeByMovieTitle(title string) *tmdb.SearchMulti {

	options := make(map[string]string)
	options["page"] = "1"
	options["include_adult"] = "false"

	multi, err := t.tmdb.GetSearchMulti(title, options)
	if err != nil {
		log.Println(err)
	}

	return multi

}
