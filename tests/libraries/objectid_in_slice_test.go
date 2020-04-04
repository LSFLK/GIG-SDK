package libraries

import (
	"GIG-SDK/libraries"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestThatObjectIdInSliceReturnsTrue(t *testing.T) {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	result := libraries.ObjectIdInSlice(testSlice, "2")
	if result != true {
		t.Errorf("objectId in slice true failed. %t != %t", result, true)
	}
}

func TestThatObjectIdInSliceReturnsFalse(t *testing.T) {
	testSlice := []bson.ObjectId{"1", "2", "3", "4"}
	result:=libraries.ObjectIdInSlice(testSlice, "5")
	if result != false {
		t.Errorf("objectId in slice false failed. %t != %t", result, false)
	}
}
