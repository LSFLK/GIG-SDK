package models

import (
	"GIG-SDK/models"
	"strings"
	"testing"
)

/*
entity set title works
 */
func TestThatEntitySetTitleWorks(t *testing.T) {

	testEntity := models.Entity{}.SetTitle(testValueObj)
	titleAttribute, err := testEntity.GetAttribute("titles")

	titleValue := titleAttribute.GetValue()

	errString := "entity set title failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if testEntity.GetTitle() != formattedValueString {
		t.Errorf("%s %s != %s", errString, testEntity.GetTitle(), formattedValueString)
	}
	if titleValue.GetValueString() != formattedValueString {
		t.Errorf("%s %s != %s", errString, titleValue.GetValueString(), formattedValueString)
	}
	if titleValue.GetType() != valueType {
		t.Errorf("%s %s != %s", errString, titleValue.GetType(), valueType)
	}
	if titleValue.GetDate() != date {
		t.Errorf("%s %s != %s", errString, titleValue.GetDate(), date)
	}
	if titleValue.GetSource() != source {
		t.Errorf("%s %s != %s", errString, titleValue.GetSource(), source)
	}
	if titleValue.GetUpdatedDate().After(date) != true {
		t.Errorf("%s %t != %t", errString, titleValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute work for new attribute
 */
func TestThatEntitySetAttributeWorksForNewAttribute(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj)
	testAttribute, err := testEntity.GetAttribute(testAttributeKey)

	testValue := testAttribute.GetValue()

	errString := "entity set attribute for new attribute failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if testValue.GetValueString() != strings.TrimSpace(valueString) {
		t.Errorf("%s %s != %s", errString, testValue.GetValueString(), strings.TrimSpace(valueString))
	}
	if testValue.GetType() != valueType {
		t.Errorf("%s %s != %s", errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date {
		t.Errorf("%s %s != %s", errString, testValue.GetDate(), date)
	}
	if testValue.GetSource() != source {
		t.Errorf("%s %s != %s", errString, testValue.GetSource(), source)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf("%s %t != %t", errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with same value
 */
func TestThatEntitySetAttributeWorksForExistingAttributeWithSameValue(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValue()
	errString := "entity set attribute for existing attribute with same value failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 1 {
		t.Errorf("%s %d != %d", errString, len(testAttribute.GetValues()), 1)
	}
	if testValue.GetValueString() != strings.TrimSpace(valueString) {
		t.Errorf("%s %s != %s", errString, testValue.GetValueString(), strings.TrimSpace(valueString))
	}
	if testValue.GetType() != valueType {
		t.Errorf("%s %s != %s", errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date {
		t.Errorf("%s %s != %s", errString, testValue.GetDate(), date)
	}
	if testValue.GetSource() != source {
		t.Errorf("%s %s != %s", errString, testValue.GetSource(), source)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf("%s %t != %t", errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with new value after the latest existing value
 */
func TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueAfterLatestExistingValue(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValue()

	errString := "entity set attribute for existing attribute with new value after latest existing value failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 2 {
		t.Errorf("%s %d != %d", errString, len(testAttribute.GetValues()), 2)
	}
	if testValue.GetValueString() != strings.TrimSpace(valueString2) {
		t.Errorf("%s %s != %s", errString, testValue.GetValueString(), strings.TrimSpace(valueString2))
	}
	if testValue.GetType() != valueType {
		t.Errorf("%s %s != %s", errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date2 {
		t.Errorf("%s %s != %s", errString, testValue.GetDate(), date2)
	}
	if testValue.GetSource() != source2 {
		t.Errorf("%s %s != %s", errString, testValue.GetSource(), source2)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf("%s %t != %t", errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with new value in between the values
 */
func TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueInBetweenExistingValues(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValues()[1]

	errString := "entity set attribute for existing attribute with new value between existing values failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 3 {
		t.Errorf("%s %d != %d", errString, len(testAttribute.GetValues()), 3)
	}
	if testValue.GetValueString() != strings.TrimSpace(valueString2) {
		t.Errorf("%s %s != %s", errString, testValue.GetValueString(), strings.TrimSpace(valueString2))
	}
	if testValue.GetType() != valueType {
		t.Errorf("%s %s != %s", errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date2 {
		t.Errorf("%s %s != %s", errString, testValue.GetDate(), date2)
	}
	if testValue.GetSource() != source2 {
		t.Errorf("%s %s != %s", errString, testValue.GetSource(), source2)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf("%s %t != %t", errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with new value before the first value date
 */
func TestThatEntitySetAttributeWorksForExistingAttributeWithNewValueBeforeTheFirstValue(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValues()[0]

	errString := "entity set attribute for existing attribute with new value before the first value failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 3 {
		t.Errorf("%s %d != %d", errString, len(testAttribute.GetValues()), 3)
	}
	if testValue.GetValueString() != strings.TrimSpace(valueString) {
		t.Errorf("%s %s != %s", errString, testValue.GetValueString(), strings.TrimSpace(valueString))
	}
	if testValue.GetType() != valueType {
		t.Errorf("%s %s != %s", errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date {
		t.Errorf("%s %s != %s", errString, testValue.GetDate(), date)
	}
	if testValue.GetSource() != source {
		t.Errorf("%s %s != %s", errString, testValue.GetSource(), source)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf("%s %t != %t", errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with same value string but with zero date in existing value
 */
func TestThatEntitySetAttributeWorksForExistingAttributeWithSameValueButWithZeroDateInExistingValue(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj0)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testValue := testAttribute.GetValues()[0]

	errString := "entity set attribute for existing attribute with same value with zero date failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 3 {
		t.Errorf("%s %d != %d", errString, len(testAttribute.GetValues()), 3)
	}
	if testValue.GetValueString() != strings.TrimSpace(valueString) {
		t.Errorf("%s %s != %s", errString, testValue.GetValueString(), strings.TrimSpace(valueString))
	}
	if testValue.GetType() != valueType {
		t.Errorf("%s %s != %s", errString, testValue.GetType(), valueType)
	}
	if testValue.GetDate() != date {
		t.Errorf("%s %s != %s", errString, testValue.GetDate(), date)
	}
	if testValue.GetSource() != source {
		t.Errorf("%s %s != %s", errString, testValue.GetSource(), source)
	}
	if testValue.GetUpdatedDate().After(date) != true {
		t.Errorf("%s %t != %t", errString, testValue.GetUpdatedDate().After(date), true)
	}
}

/*
set attribute works for existing attribute with different values string but with same date
 */
func TestThatEntitySetAttributeWorksForExistingAttributeWithDifferentValuesButWithSameDate(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj0)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2.SetValueString("different value same date"))

	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testAttributeValues := testAttribute.GetValues()

	errString := "entity set attribute for existing attribute with different values but with same date failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 4 {
		t.Errorf("%s %d != %d", errString, len(testAttribute.GetValues()), 4)
	}
	if testAttributeValues[1].GetValueString() == testAttributeValues[2].GetValueString() {
		t.Errorf("%s %s == %s", errString, testAttributeValues[1].GetValueString(), testAttributeValues[2].GetValueString())
	}
	if testAttributeValues[1].GetDate() != testAttributeValues[2].GetDate() {
		t.Errorf("%s %s != %s", errString, testAttributeValues[1].GetDate(), testAttributeValues[2].GetDate())
	}
}

/*
set attribute works for existing attribute with same value string and new value with past date
 */
func TestThatEntitySetAttributeWorksForExistingAttributeWithSameValueAndNewValueWithPastDate(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj0)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2.SetDate(testValueObj.GetDate()))
	testAttribute, err := testEntity.GetAttribute(testAttributeKey)
	testAttributeValues := testAttribute.GetValues()

	errString := "entity set attribute for existing attribute with same value and new value with past date failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 4 {
		t.Errorf("%s %d != %d", errString, len(testAttribute.GetValues()), 4)
	}
	if testValueObj.GetDate().Before(testValueObj2.GetDate()) != true {
		t.Errorf("%s %t != %t", errString, testValueObj.GetDate().Before(testValueObj2.GetDate()), true)
	}
	if testAttributeValues[1].GetValueString() != testAttributeValues[2].GetValueString() {
		t.Errorf("%s %s != %s", errString, testAttributeValues[1].GetValueString(), testAttributeValues[2].GetValueString())
	}
}

/*
set attribute works for existing attribute with same value string and new value with future date within value lifetime
 */
func TestThatEntitySetAttributeWorksForExistingAttributeWithSameValueAndNewValueWithFutureDateWithinValueLifetime(t *testing.T) {

	testEntity := models.Entity{}.SetAttribute(testAttributeKey, testValueObj3)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj0)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj)
	testEntity = testEntity.SetAttribute(testAttributeKey, testValueObj2.SetDate(date25))
	testAttribute, err := testEntity.GetAttribute(testAttributeKey)

	errString := "entity set attribute for existing attribute with same value and new value with future date withing value lifetime failed."
	if err != nil {
		t.Errorf("%s %s != %#v", errString, err.Error(), nil)
	}
	if len(testAttribute.GetValues()) != 3 {
		t.Errorf("%s %d != %d", errString, len(testAttribute.GetValues()), 3)
	}
	if testValueObj.GetDate().Before(date25) != true {
		t.Errorf("%s %t != %t", errString, testValueObj.GetDate().Before(date25), true)
	}
}
