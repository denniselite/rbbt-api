package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type addTopicResponse struct {
	Id int `json:"id"`
}

type updateTopicRequest struct {
	Id            int  `json:"id"`
	VoteDirection bool `json:"voteDirection"`
}

type emptyResponse struct{}

type getTopicsResponse struct {
	Items []Topic `json:"items"`
}

func addTopicHandler(ctx context.Context) {
	var t Topic
	err := ctx.ReadJSON(&t)
	if err != nil {
		msg := errorMessage{iris.StatusBadRequest, "Invalid JSON"}
		ctx.JSON(msg)
		return
	}
	var res addTopicResponse
	res.Id = storage.Add(t)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(res)
}

func updateTopicHandler(ctx context.Context) {
	var req updateTopicRequest
	err := ctx.ReadJSON(&req)
	if err != nil {
		msg := errorMessage{iris.StatusBadRequest, "Invalid JSON"}
		ctx.JSON(msg)
		return
	}
	t, err := storage.GetTopicById(req.Id)
	if err != nil {
		msg := errorMessage{iris.StatusBadRequest, err.Error()}
		ctx.JSON(msg)
		return
	}
	if req.VoteDirection {
		t.Rating++
	} else {
		t.Rating--
	}
	err = storage.Update(req.Id, t)
	if err != nil {
		msg := errorMessage{iris.StatusBadRequest, err.Error()}
		ctx.JSON(msg)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(emptyResponse{})
}

func getTopicsHandler(ctx context.Context) {
	var res getTopicsResponse
	res.Items = storage.GetItems()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(res)
}
