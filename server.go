package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	v1 "github.com/ucpr/atlas-search-example/proto/go/app/v1"
	"github.com/ucpr/atlas-search-example/proto/go/app/v1/appv1connect"
	"go.mongodb.org/mongo-driver/bson"
)

func NewHandler(mcli *Client) http.Handler {
	mux := http.NewServeMux()

	handler := &Handler{
		mcli:    mcli,
		timeNow: time.Now,
	}
	path, sh := appv1connect.NewServiceHandler(handler)
	log.Println("mounted path:", path)
	mux.Handle(path, sh)

	return mux
}

func NewServer(cfg *Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}

type Handler struct {
	mcli    *Client
	timeNow func() time.Time
}

func (s *Handler) GetPost(ctx context.Context, in *connect.Request[v1.GetPostRequest]) (*connect.Response[v1.GetPostResponse], error) {
	col := s.mcli.Post()
	var post Post
	if err := col.FindOne(ctx, bson.M{"_id": in.Msg.Id}).Decode(&post); err != nil {
		return nil, err
	}

	return &connect.Response[v1.GetPostResponse]{
		Msg: &v1.GetPostResponse{
			Post: &v1.Post{
				Id:        post.ID,
				Title:     post.Title,
				Content:   post.Content,
				Author:    post.Author,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
			},
		},
	}, nil
}

func (s *Handler) ListPosts(ctx context.Context, _ *connect.Request[v1.ListPostsRequest]) (*connect.Response[v1.ListPostsResponse], error) {
	panic("not implemented") // TODO: Implement
}

func (s *Handler) CreatePost(ctx context.Context, in *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error) {
	col := s.mcli.Post()
	post := &Post{
		ID:        uuid.New().String(),
		Title:     in.Msg.Title,
		Content:   in.Msg.Content,
		Author:    in.Msg.Author,
		CreatedAt: s.timeNow().Unix(),
		UpdatedAt: s.timeNow().Unix(),
	}
	if _, err := col.InsertOne(ctx, post); err != nil {
		return nil, err
	}
	return &connect.Response[v1.CreatePostResponse]{
		Msg: &v1.CreatePostResponse{
			Post: &v1.Post{
				Id:        post.ID,
				Title:     post.Title,
				Content:   post.Content,
				Author:    post.Author,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
			},
		},
	}, nil
}
