package handler

import (
	"fmt"

	"github.com/carlmjohnson/errorx"
	"github.com/mackee/tanukirpc"
	"github.com/mackee/tanukirpc/genclient"
	"github.com/mackee/template-tanukistack/client"
	"github.com/mackee/template-tanukistack/record"
)

//go:generate go tool gentypescript -out ../dist/api.ts ./

func NewHandler(cli *client.Client) *tanukirpc.Router[*Registry] {
	registry := NewRegistry(cli)
	router := tanukirpc.NewRouter(registry)
	router.Route("/api", func(router *tanukirpc.Router[*Registry]) {
		router.Get("/hello", tanukirpc.NewHandler(hello))
		router.Get("/messages", tanukirpc.NewHandler(listMessages))
		router.Post("/messages", tanukirpc.NewHandler(postMessage))
	})

	genclient.AnalyzeTarget(router)

	return router
}

type HelloRequest struct {
	Name string `query:"name" validate:"required"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func hello(_ tanukirpc.Context[*Registry], req HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}

type ListMessagesRequest struct {
	StartID int64 `query:"start_id"`
}

type ListMessagesResponse struct {
	Messages record.Messages `json:"messages"`
}

func listMessages(ctx tanukirpc.Context[*Registry], req ListMessagesRequest) (_resp *ListMessagesResponse, err error) {
	defer errorx.Trace(&err)
	db := ctx.Registry().DB
	messages, err := record.NewMessageTable().ListByMessages(ctx, db, req.StartID)
	if err != nil {
		return nil, fmt.Errorf("failed to list messages: %w", err)
	}

	return &ListMessagesResponse{
		Messages: messages,
	}, nil
}

type PostMessageRequest struct {
	Author string `json:"author" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

type PostMessageResponse struct {
	Message *record.Message `json:"message"`
}

func postMessage(ctx tanukirpc.Context[*Registry], req PostMessageRequest) (_resp *PostMessageResponse, err error) {
	defer errorx.Trace(&err)
	db := ctx.Registry().DB
	message, err := record.NewMessageTable().Create(ctx, db, record.MessageTableCreateInput{
		Author: req.Author,
		Text:   req.Text,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}

	return &PostMessageResponse{
		Message: message,
	}, nil
}
