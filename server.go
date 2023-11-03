package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewHandler(mcli *Client) http.Handler {
	mux := http.NewServeMux()

	handler := &Handler{
		mcli: mcli,
	}
	mux.HandleFunc("/search", handler.search)

	return mux
}

func NewServer(cfg *Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}

type Handler struct {
	mcli *Client
}

func (h *Handler) search(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	col := h.mcli.Movies()
	query := r.URL.Query().Get("query")

	searchStage := bson.D{
		{Key: "$search", Value: bson.D{
			{
				Key: "text", Value: bson.D{
					{
						Key: "path", Value: "title",
					},
					{
						Key: "query", Value: query,
					},
				},
			},
		}},
	}
	limitStage := bson.D{
		{Key: "$limit", Value: 2},
	}
	projectStage := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "title", Value: 1},
				{Key: "_id", Value: 1},
			},
		},
	}
	// 検索クエリを実行します。
	opts := options.Aggregate().SetMaxTime(1 * time.Second)
	cursor, err := col.Aggregate(ctx, mongo.Pipeline{searchStage, limitStage, projectStage}, opts)
	if err != nil {
		log.Println("failed to aggregate:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 結果をデコードします。
	var results []Movie
	if err := cursor.All(ctx, &results); err != nil {
		log.Println("failed to decode results:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 結果を表示します。
	log.Println("searched:", query, "results:", len(results))

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		log.Println("failed to encode results:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
