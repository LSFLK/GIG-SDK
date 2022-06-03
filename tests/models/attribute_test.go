package models

import (
	"github.com/lsflk/gig-sdk/models"
	"testing"
	"time"
)

/*
entity set title works
*/
func TestThatAttributeGetValueByDateWorks(t *testing.T) {

	testEntity := new(models.Entity).
		SetAttribute(testAttributeKey, *testValueObj3).
		SetAttribute(testAttributeKey, *testValueObj2).
		SetAttribute(testAttributeKey, *testValueObj0).
		SetAttribute(testAttributeKey, *testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)

	testDate, _ := time.Parse(ReferenceDate, "2010-5-18")
	testDate1, _ := time.Parse(ReferenceDate, "2010-5-21")
	testDate2, _ := time.Parse(ReferenceDate, "2010-5-22")
	testDate3, _ := time.Parse(ReferenceDate, "2010-8-30")
	testDate4, _ := time.Parse(ReferenceDate, "2012-8-30")

	errString := "attribute get value by date failed."
	if err != nil {
		t.Errorf(FormatSSEV, errString, err.Error(), nil)
	}
	result := testAttribute.GetValueByDate(testDate).GetSource()
	if result != "" {
		t.Errorf(FormatSSUS, errString, result, "")
	}
	result1 := testAttribute.GetValueByDate(testDate1).GetSource()
	if result1 != source {
		t.Errorf(FormatSSUS, errString, result1, source)
	}
	result2 := testAttribute.GetValueByDate(testDate2).GetSource()
	if result2 != source2 {
		t.Errorf(FormatSSUS, errString, result2, source2)
	}
	result3 := testAttribute.GetValueByDate(testDate3).GetSource()
	if result3 != source2 {
		t.Errorf(FormatSSUS, errString, result3, source2)
	}
	result4 := testAttribute.GetValueByDate(testDate4).GetSource()
	if result4 != source3 {
		t.Errorf(FormatSSUS, errString, result4, source3)
	}
}
