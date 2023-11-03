package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Plot      string             `bson:"plot" json:"plot"`
	Genres    []string           `bson:"genres" json:"genres"`
	Runtime   int                `bson:"runtime" json:"runtime"`
	Cast      []string           `bson:"cast" json:"cast"`
	Poster    string             `bson:"poster" json:"poster"`
	Title     string             `bson:"title" json:"title"`
	FullPlot  string             `bson:"fullplot" json:"full_plot"`
	Languages []string           `bson:"languages" json:"languages"`
	Released  time.Time          `bson:"released" json:"released"`
	Directors []string           `bson:"directors" json:"directors"`
	Rated     string             `bson:"rated" json:"rated"`
	Awards    struct {
		Wins        int    `bson:"wins" json:"wins"`
		Nominations int    `bson:"nominations" json:"nominations"`
		Text        string `bson:"text" json:"text"`
	} `bson:"awards" json:"awards"`
	LastUpdated string `bson:"lastupdated" json:"last_updated"`
	Year        int    `bson:"year" json:"year"`
	IMDb        struct {
		Rating float64 `bson:"rating" json:"rating"`
		Votes  int     `bson:"votes" json:"votes"`
		ID     int     `bson:"id" json:"id"`
	} `bson:"imdb" json:"im_db"`
	Countries []string `bson:"countries" json:"countries"`
	Type      string   `bson:"type" json:"type"`
	Tomatoes  struct {
		Viewer struct {
			Rating     float64 `bson:"rating" json:"rating"`
			NumReviews int     `bson:"numReviews" json:"num_reviews"`
			Meter      int     `bson:"meter" json:"meter"`
		} `bson:"viewer" json:"viewer"`
		Fresh struct {
			NumReviews int `bson:"numReviews" json:"num_reviews"`
		} `bson:"fresh" json:"fresh"`
		Critic struct {
			Rating     float64 `bson:"rating" json:"rating"`
			NumReviews int     `bson:"numReviews" json:"num_reviews"`
			Meter      int     `bson:"meter" json:"meter"`
		} `bson:"critic" json:"critic"`
		Rotten      int       `bson:"rotten" json:"rotten"`
		LastUpdated time.Time `bson:"lastUpdated" json:"last_updated"`
	} `bson:"tomatoes" json:"tomatoes"`
	NumMflixComments int `bson:"num_mflix_comments" json:"num_mflix_comments"`
}
