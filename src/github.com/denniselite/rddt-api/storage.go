package main

import (
	"fmt"
	"github.com/denniselite/iris-fixed/core/errors"
)

type Topic struct {
	Author string `json:"author"`
	Body   string `json:"body"`
	Rating int    `json:"rating"`
}

type Storage struct {
	data []Topic
}

// Add topics to storage
func (s *Storage) Add(t Topic) (id int) {
	s.data = append(s.data, t)
	id = len(s.data) - 1
	return
}

// Update topic's data in the storage
func (s *Storage) Update(id int, t Topic) (err error) {
	if id < 0 || id >= len(s.data) {
		err = errors.New(fmt.Sprintf("Topic with ID %d is not exists", id))
		return
	}
	s.data[id] = t
	return
}

// Returns all items of storage; when it's empty -
// returns empty slice of the Topic type
func (s *Storage) GetItems() []Topic {

	// For an empty array in a response instead of null
	if len(s.data) == 0 {
		s.data = make([]Topic, 0)
	}
	return s.data
}

// Return topic by ID
func (s *Storage) GetTopicById(id int) (t Topic, err error) {
	if id < 0 || id >= len(s.data) {
		err = errors.New(fmt.Sprintf("Topic with ID %d is not exists", id))
		return
	}
	t = s.data[id]
	return
}
