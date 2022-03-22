package libraries

import (
	"github.com/lsflk/gig-sdk/libraries"
	"testing"
)

func TestThatFixUrlWorks(t *testing.T) {
	href := "/images/pdf/registered%20suppliers%20and%20service%20%20providers%20for%20the%20year%202019%20-%20g1-office%20stationery.pdf"
	baseUrl := "http://www.buildings.gov.lk/index.php?option=com_content&view=category&layout=blog&id=47&Itemid=128&lang=en"
	result := libraries.FixUrl(href, baseUrl)
	expectedResult := "http://www.buildings.gov.lk" + href
	if result != expectedResult {
		t.Errorf("fix url failed. %s != %s", result, expectedResult)
	}
}

func TestThatFixUrlWorksForAbsoluteUrl(t *testing.T) {
	href := "http://www.buildings.gov.lk/images/pdf/registered%20suppliers%20and%20service%20%20providers%20for%20the%20year%202019%20-%20g1-office%20stationery.pdf"
	baseUrl := "http://www.buildings.gov.lk/index.php?option=com_content&view=category&layout=blog&id=47&Itemid=128&lang=en"
	result := libraries.FixUrl(href, baseUrl)
	if result != href {
		t.Errorf("fix url for absolute url failed. %s != %s", result, href)
	}
}
