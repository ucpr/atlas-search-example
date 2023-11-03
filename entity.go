package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Plot      string             `bson:"plot" json:"plot,omitempty"`
	Genres    []string           `bson:"genres" json:"genres,omitempty"`
	Runtime   int                `bson:"runtime" json:"runtime,omitempty"`
	Cast      []string           `bson:"cast" json:"cast,omitempty"`
	Poster    string             `bson:"poster" json:"poster,omitempty"`
	Title     string             `bson:"title" json:"title,omitempty"`
	FullPlot  string             `bson:"fullplot" json:"full_plot,omitempty"`
	Languages []string           `bson:"languages" json:"languages,omitempty"`
	Released  time.Time          `bson:"released" json:"released,omitempty"`
	Directors []string           `bson:"directors" json:"directors,omitempty"`
	Rated     string             `bson:"rated" json:"rated,omitempty"`
	Awards    struct {
		Wins        int    `bson:"wins" json:"wins,omitempty"`
		Nominations int    `bson:"nominations" json:"nominations,omitempty"`
		Text        string `bson:"text" json:"text,omitempty"`
	} `bson:"awards" json:"awards,omitempty"`
	LastUpdated string `bson:"lastupdated" json:"last_updated,omitempty"`
	Year        int    `bson:"year" json:"year,omitempty"`
	IMDb        struct {
		Rating float64 `bson:"rating" json:"rating,omitempty"`
		Votes  int     `bson:"votes" json:"votes,omitempty"`
		ID     int     `bson:"id" json:"id,omitempty"`
	} `bson:"imdb" json:"im_db,omitempty"`
	Countries []string `bson:"countries" json:"countries,omitempty"`
	Type      string   `bson:"type" json:"type,omitempty"`
	Tomatoes  struct {
		Viewer struct {
			Rating     float64 `bson:"rating" json:"rating,omitempty"`
			NumReviews int     `bson:"numReviews" json:"num_reviews,omitempty"`
			Meter      int     `bson:"meter" json:"meter,omitempty"`
		} `bson:"viewer" json:"viewer,omitempty"`
		Fresh  int `bson:"fresh" json:"fresh,omitempty"`
		Critic struct {
			Rating     float64 `bson:"rating" json:"rating,omitempty"`
			NumReviews int     `bson:"numReviews" json:"num_reviews,omitempty"`
			Meter      int     `bson:"meter" json:"meter,omitempty"`
		} `bson:"critic" json:"critic,omitempty"`
		Rotten      int       `bson:"rotten" json:"rotten,omitempty"`
		LastUpdated time.Time `bson:"lastUpdated" json:"last_updated,omitempty"`
	} `bson:"tomatoes" json:"tomatoes,omitempty"`
	NumMflixComments int `bson:"num_mflix_comments" json:"num_mflix_comments,omitempty"`
}
