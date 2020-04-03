package tests

import (
	"GIG/sdk"
	"gopkg.in/mgo.v2/bson"
)

func (t *TestSDK) TestThatObjectIdInSliceTestReturnsTrue() {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	t.AssertEqual(sdk.ObjectIdInSlice(testSlice, "2"), true)
}

func (t *TestSDK) TestThatObjectIdInSliceTestReturnsFalse() {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	t.AssertEqual(sdk.ObjectIdInSlice(testSlice, "5"), false)
}
