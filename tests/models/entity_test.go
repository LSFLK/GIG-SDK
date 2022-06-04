package models

import (
	"github.com/lsflk/gig-sdk/models"
	"strings"
	"testing"
)

/*
entity set title works
*/
func TestThatEntitySetTitleWorks(t *testing.T) {

	testEntity := new(models.Entity).SetTitle(testValueObj)
	titleAttribute, err := testEntity.GetAttribute("titles")

	titleValue := titleAttribute.GetValue()

	errString := "entity set title failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	if testEntity.GetTitle() != formattedValueString {
		t.Errorf(FormatSSUS, errString, testEntity.GetTitle(), formattedValueString)
	}
	if titleValue.GetValueString() != formattedValueString {
		t.Errorf(FormatSSUS, errString, titleValue.GetValueString(), formattedValueString)
	}
	if titleValue.GetType() != valueType {
		t.Errorf(FormatSSUS, errString, titleValue.GetType(), valueType)
	}
	if titleValue.GetDate() != date {
		t.Errorf(FormatSSUS, errString, titleValue.GetDate(), date)
	}
	if titleValue.GetSource() != source {
		t.Errorf(FormatSSUS, errString, titleValue.GetSource(), source)
	}
	if titleValue.GetUpdatedDate().After(date) != true {
		t.Errorf(FormatSTUT, errString, titleValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute work for new attribute
*/
func TestThatEntitySetAttributeWorksForNewAttribute(t *testing.T) {

	testEntity := new(models.Entity).SetAttribute(testAttributeKey, testValueObj)
	testAttribute, err := testEntity.GetAttribute(testAttributeKey)

	testValue := testAttribute.GetValue()

	errString := "entity set attribute for new attribute failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	evaluatedSetAttributeWorks(testValue, t, errString)

}

func evaluatedSetAttributeWorks(testValue models.Value, t *testing.T, errString string) {
	if testValue.GetValueString() != strings.TrimSpace(valueString) {
		t.Errorf(FormatSSUS, errString, testValue.GetValueString(), strings.TrimSpace(valueString))
	}
	if testValue.GetType() != valueType {
		t.Errorf(FormatSSUS, errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date {
		t.Errorf(FormatSSUS, errString, testValue.GetDate(), date)
	}
	if testValue.GetSource() != source {
		t.Errorf(FormatSSUS, errString, testValue.GetSource(), source)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf(FormatSTUT, errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with same value
*/
func TestThatEntitySetAttributeWorksForExistingAttributeWithSameValue(t *testing.T) {

	testEntity := new(models.Entity).SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValue()
	errString := "entity set attribute for existing attribute with same value failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	evaluatedSetAttributeWorks(testValue, t, errString)
}

/*
set attribute works for existing attribute with new value after the latest existing value
*/
func TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueAfterLatestExistingValue(t *testing.T) {

	testEntity := new(models.Entity).SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValue()

	errString := "entity set attribute for existing attribute with new value after latest existing value failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 2 {
		t.Errorf(FormatSDUD, errString, len(testAttribute.GetValues()), 2)
	}
	evaluateEntitySetAttributeWorksForExistingAttributeWithNewValue(testValue, t, errString)
}

/*
set attribute works for existing attribute with new value in between the values
*/
func TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueInBetweenExistingValues(t *testing.T) {

	testEntity := new(models.Entity).SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValues()[1]

	errString := "entity set attribute for existing attribute with new value between existing values failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 3 {
		t.Errorf(FormatSDUD, errString, len(testAttribute.GetValues()), 3)
	}
	evaluateEntitySetAttributeWorksForExistingAttributeWithNewValue(testValue, t, errString)
}

func evaluateEntitySetAttributeWorksForExistingAttributeWithNewValue(testValue models.Value, t *testing.T, errString string) {
	if testValue.GetValueString() != strings.TrimSpace(valueString2) {
		t.Errorf(FormatSSUS, errString, testValue.GetValueString(), strings.TrimSpace(valueString2))
	}
	if testValue.GetType() != valueType {
		t.Errorf(FormatSSUS, errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date2 {
		t.Errorf(FormatSSUS, errString, testValue.GetDate(), date2)
	}
	if testValue.GetSource() != source2 {
		t.Errorf(FormatSSUS, errString, testValue.GetSource(), source2)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf(FormatSTUT, errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with new value before the first value date
*/
func TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueBeforeTheFirstValue(t *testing.T) {

	testEntity := *new(models.Entity).SetAttribute(testAttributeKey, testValueObj3).SetAttribute(testAttributeKey, testValueObj2)
	evaluateEntitySetAttributeWorksForExistingAttribute(testEntity, t)
}

/*
set attribute works for existing attribute with same value string but with zero date in existing value
*/
func TestThatEntitySetAttributeWorksForExistingAttributeWithSameValueButWithZeroDateInExistingValue(t *testing.T) {

	testEntity := *new(models.Entity).
		SetAttribute(testAttributeKey, testValueObj3).
		SetAttribute(testAttributeKey, testValueObj2).
		SetAttribute(testAttributeKey, testValueObj0)
	evaluateEntitySetAttributeWorksForExistingAttribute(testEntity, t)

}

func evaluateEntitySetAttributeWorksForExistingAttribute(testEntity models.Entity, t *testing.T) {
	testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValues()[0]

	errString := "entity set attribute for existing attribute with same value with zero date failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 3 {
		t.Errorf(FormatSDUD, errString, len(testAttribute.GetValues()), 3)
	}
	if testValue.GetValueString() != strings.TrimSpace(valueString) {
		t.Errorf(FormatSSUS, errString, testValue.GetValueString(), strings.TrimSpace(valueString))
	}
	if testValue.GetType() != valueType {
		t.Errorf(FormatSSUS, errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date {
		t.Errorf(FormatSSUS, errString, testValue.GetDate(), date)
	}
	if testValue.GetSource() != source {
		t.Errorf(FormatSSUS, errString, testValue.GetSource(), source)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf(FormatSTUT, errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with different values string but with same date
*/
func TestThatEntitySetAttributeWorksForExistingAttributeWithDifferentValuesButWithSameDate(t *testing.T) {

	testEntity := new(models.Entity).
		SetAttribute(testAttributeKey, testValueObj3).
		SetAttribute(testAttributeKey, testValueObj2).
		SetAttribute(testAttributeKey, testValueObj0).
		SetAttribute(testAttributeKey, testValueObj).
		SetAttribute(testAttributeKey, testValueObj5)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testAttributeValues := testAttribute.GetValues()

	errString := "entity set attribute for existing attribute with different values but with same date failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 4 {
		t.Errorf(FormatSDUD, errString, len(testAttribute.GetValues()), 4)
	}
	if testAttributeValues[1].GetValueString() == testAttributeValues[2].GetValueString() {
		t.Errorf(FormatSSES, errString, testAttributeValues[1].GetValueString(), testAttributeValues[2].GetValueString())
	}
	if testAttributeValues[0].GetDate() != testAttributeValues[1].GetDate() {
		t.Errorf(FormatSSUS, errString, testAttributeValues[0].GetDate(), testAttributeValues[1].GetDate())
	}
}

/*
set attribute works for existing attribute with same value string and new value with past date
*/
func TestThatEntitySetAttributeWorksForExistingAttributeWithSameValueAndNewValueWithPastDate(t *testing.T) {

	testEntity := new(models.Entity).
		SetAttribute(testAttributeKey, testValueObj3).
		SetAttribute(testAttributeKey, testValueObj2).
		SetAttribute(testAttributeKey, testValueObj0).
		SetAttribute(testAttributeKey, testValueObj).
		SetAttribute(testAttributeKey, testValueObj4)
	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testAttributeValues := testAttribute.GetValues()

	errString := "entity set attribute for existing attribute with same value and new value with past date failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 4 {
		t.Errorf(FormatSDUD, errString, len(testAttribute.GetValues()), 4)
	}
	if testValueObj.GetDate().Before(testValueObj2.GetDate()) != true {
		t.Errorf(FormatSTUT, errString, testValueObj.GetDate().Before(testValueObj2.GetDate()), true)
	}
	if testAttributeValues[1].GetValueString() != testAttributeValues[2].GetValueString() {
		t.Errorf(FormatSSUS, errString, testAttributeValues[1].GetValueString(), testAttributeValues[2].GetValueString())
	}
}

/*
set attribute works for existing attribute with same value string and new value with future date within value lifetime
*/
func TestThatEntitySetAttributeWorksForExistingAttributeWithSameValueAndNewValueWithFutureDateWithinValueLifetime(t *testing.T) {

	testEntity := new(models.Entity).SetAttribute(testAttributeKey, testValueObj3).
		SetAttribute(testAttributeKey, testValueObj2).
		SetAttribute(testAttributeKey, testValueObj0).
		SetAttribute(testAttributeKey, testValueObj).
		SetAttribute(testAttributeKey, *testValueObj2.SetDate(date25))
	testAttribute, err := testEntity.GetAttribute(testAttributeKey)

	errString := "entity set attribute for existing attribute with same value and new value with future date withing value lifetime failed."
	if err != nil {
		t.Errorf(FormatSSUV, errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 3 {
		t.Errorf(FormatSDUD, errString, len(testAttribute.GetValues()), 3)
	}
	if testValueObj.GetDate().Before(date25) != true {
		t.Errorf(FormatSTUT, errString, testValueObj.GetDate().Before(date25), true)
	}
}
