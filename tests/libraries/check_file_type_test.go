package libraries

import (
	"GIG-SDK/libraries"
	"testing"
)

func TestThatFileTypeCheckTrueWorks(t *testing.T) {
	link := "http://www.buildings.gov.lk/index.php"
	result := libraries.FileTypeCheck(link, "php")
	if result != true {
		t.Errorf("file type check true failed. %t != %t", result, true)
	}
}

func TestThatFileTypeCheckFalseWorks(t *testing.T) {
	link := "http://www.buildings.gov.lk/index.php"
	result := libraries.FileTypeCheck(link, "pdf")
	if result != false {
		t.Errorf("file type check false failed. %t != %t", result, false)
	}
}
