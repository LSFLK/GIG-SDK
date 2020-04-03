package tests

import (
	"GIG-SDK"
	"os"
)

func (t *TestSDK) TestThatDownloadFileWorks() {
	os.Remove("app/cache/test.pdf")
	link := "https://s1.q4cdn.com/806093406/files/doc_downloads/test.pdf"
	result := libraries.DownloadFile("app/cache/test.pdf",link)
	t.AssertEqual(nil, result)
}