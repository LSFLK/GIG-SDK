package models

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

var (
	source0              = "source0"
	source               = "source"
	source2              = "source2"
	source3              = "source3"
	valueString          = "~test /tit?le % "
	valueString2         = "~test /tit?le % 2"
	valueString3         = "~test /tit?le % 3"
	date, _              = time.Parse(ReferenceDate, "2010-5-20")
	date2, _             = time.Parse(ReferenceDate, "2010-5-22")
	date25, _            = time.Parse(ReferenceDate, "2010-11-22")
	date3, _             = time.Parse(ReferenceDate, "2011-5-22")
	valueType            = ValueType.String
	formattedValueString = "2test -title"
	testAttributeKey     = "test_attribute"

	testValueObj0 = *new(models.Value).
			SetSource(source0).
			SetValueString(valueString).
			SetType(valueType)

	testValueObj = *new(models.Value).
			SetSource(source).
			SetValueString(valueString).
			SetDate(date).
			SetType(valueType)

	testValueObj2 = *new(models.Value).
			SetSource(source2).
			SetValueString(valueString2).
			SetDate(date2).
			SetType(valueType)

	testValueObj3 = *new(models.Value).
			SetSource(source3).
			SetValueString(valueString3).
			SetDate(date3).
			SetType(valueType)

	testValueObj4 = *new(models.Value).
			SetSource(source2).
			SetValueString(valueString2).
			SetDate(date).
			SetType(valueType)
)
