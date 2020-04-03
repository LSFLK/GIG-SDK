package tests

import (
	"github.com/revel/revel/testing"
)

type TestSDK struct {
	testing.TestSuite
}

func (t *TestSDK) Before() {
	println("Set up")
}

func (t *TestSDK) After() {
	println("Tear down")
}
