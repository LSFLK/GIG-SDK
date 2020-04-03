package request_handlers

import (
	"GIG-SDK/request_handlers"
)

func (t *TestRequestHandlers) TestThatGetRequestWorks() {
	link := "http://www.buildings.gov.lk/index.php"
	result, _ := request_handlers.GetRequest(link)
	t.AssertNotEqual(result,"")
}