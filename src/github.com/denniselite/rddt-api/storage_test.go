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