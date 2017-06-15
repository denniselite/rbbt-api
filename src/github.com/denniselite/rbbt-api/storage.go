package main

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
)

type Topic struct {
	Author  string `json:"author"`
	Message string `json:"message"`
	Rating  int    `json:"rating"`
}

type Storage struct {
	data []Topic
}

func (s *Storage) Add(t Topic) (id int) {
	s.data = append(s.data, t)
	return
}

func (s *Storage) Update(id int, t Topic) (err error) {
	if id < 0 || id >= len(s.data) {
		return errors.New(fmt.Sprintf("Topic with ID %d is not exists", id))
	}
	s.data[id] = t
	return nil
}

func (s *Storage) GetItems() []Topic {
	return s.data
}

func (s *Storage) GetTopicById(id int) (t Topic, err error) {
	if id < 0 || id >= len(s.data) {
		return t, errors.New(fmt.Sprintf("Topic with ID %d is not exists", id))
	}
	return s.data[id], nil
}
