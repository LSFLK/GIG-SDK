package tests

import (
	"GIG-SDK/libraries"
	"gopkg.in/mgo.v2/bson"
)

func (t *TestSDK) TestThatObjectIdInSliceTestReturnsTrue() {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	t.AssertEqual(libraries.ObjectIdInSlice(testSlice, "2"), true)
}

func (t *TestSDK) TestThatObjectIdInSliceTestReturnsFalse() {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	t.AssertEqual(libraries.ObjectIdInSlice(testSlice, "5"), false)
}
