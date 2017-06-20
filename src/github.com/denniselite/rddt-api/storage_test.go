package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStorage_Add(t *testing.T) {

	var storage Storage
	Convey("Test: add topic to storage", t, func() {
		topic := Topic{"testAuthor", "testBody", 0}
		id := storage.Add(topic)
		So(id, ShouldEqual, 0)
	})

}

func TestStorage_GetItems(t *testing.T) {

	var storage Storage

	Convey("Test: get all topics from empty storage", t, func() {
		topics := storage.GetItems()
		emptyTopics := make([]Topic, 0)
		So(topics, ShouldResemble, emptyTopics)
	})

	topic := Topic{"testAuthor", "testBody", 0}
	storage.Add(topic)

	Convey("Test: get all topics from non-empty storage", t, func() {
		topics := storage.GetItems()
		comparedTopics := make([]Topic, 0)
		comparedTopics = append(comparedTopics, topic)
		So(topics, ShouldResemble, comparedTopics)
	})

}

func TestStorage_GetTopicById(t *testing.T) {

	var storage Storage
	topic := Topic{"testAuthor", "testBody", 0}
	storage.Add(topic)

	Convey("Test: get topic with wrong ID", t, func() {
		t, err := storage.GetTopicById(1)
		So(err, ShouldNotBeNil)
		emptyTopic := Topic{}
		So(t, ShouldResemble, emptyTopic)
	})

	Convey("Test: get topic with correct ID", t, func() {
		t, err := storage.GetTopicById(0)
		So(err, ShouldBeNil)
		So(t, ShouldResemble, topic)
	})

}

func TestStorage_Update(t *testing.T) {

	var storage Storage
	topic := Topic{"testAuthor", "testBody", 0}
	storage.Add(topic)

	Convey("Test: update topic with wrong ID", t, func() {
		topic := Topic{"testAuthor", "testBody", 1}
		err := storage.Update(1, topic)
		So(err, ShouldNotBeNil)
	})

	Convey("Test: update topic with correct ID", t, func() {
		topic := Topic{"testAuthor", "testBody", 1}
		err := storage.Update(0, topic)
		So(err, ShouldBeNil)
		updatedTopic, _ := storage.GetTopicById(0)
		So(updatedTopic, ShouldResemble, topic)
	})

}