package flex

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stretchr/testify/assert"
)

func TestDataConversionsString(t *testing.T) {
	// nil slice
	lv, d := StringSliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToStringSlice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []string{}
	lv, d = StringSliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	expectedSlice = []string{"foo", "bar"}
	lv, d = StringSliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	actualSlice, d := ListValueToStringSlice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)

	lvNull := basetypes.NewListNull(types.StringType)
	assert.True(t, lvNull.IsNull())
	actualSlice, d = ListValueToStringSlice(lvNull)
	assert.False(t, d.HasError())
	assert.Nil(t, actualSlice)
}

func TestDataConversionsAny(t *testing.T) {
	// nil value
	tfValue := AnyToStringValue(nil)
	assert.NotNil(t, tfValue)
	assert.True(t, tfValue.IsNull())

	// string pointer
	s := "Foo"
	tfValue = AnyToStringValue(&s)
	assert.NotNil(t, tfValue)
	assert.Equal(t, s, tfValue.ValueString())

	// string
	tfValue = AnyToStringValue(s)
	assert.NotNil(t, tfValue)
	assert.Equal(t, s, tfValue.ValueString())

	// boolean
	tfValue = AnyToStringValue(false)
	assert.NotNil(t, tfValue)
	assert.Equal(t, "false", tfValue.ValueString())

	// nil value
	sdkValue := StringValueToAny(tfValue)
	assert.NotNil(t, sdkValue)
	assert.Equal(t, "false", sdkValue.(string))

	// nil slice
	lv, d := AnySliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	slice, d := ListValueToAnySlice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []any{}
	lv, d = AnySliceToListValue(expectedSlice)
	assert.NotNil(t, lv)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	anySlice := []any{"foo", "bar", 1, 2, true}
	expectedSlice = []any{"foo", "bar", "1", "2", "true"}
	lv, d = AnySliceToListValue(anySlice)
	assert.NotNil(t, lv)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	actualSlice, d := ListValueToAnySlice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)

	lvNull := basetypes.NewListNull(types.StringType)
	assert.True(t, lvNull.IsNull())
	actualSlice, d = ListValueToAnySlice(lvNull)
	assert.False(t, d.HasError())
	assert.Nil(t, actualSlice)
}

func TestDataConversionsDate(t *testing.T) {
	// null string value
	nullStringValue := types.StringNull()
	date, d := StringValueToDate(nullStringValue)
	assert.False(t, d.HasError())
	assert.Nil(t, date)

	// nil date value
	tfValue := DateToStringValue(nil)
	assert.NotNil(t, tfValue)
	assert.True(t, tfValue.IsNull())

	// valid date value.
	dateString := "2025-01-01"
	date, d = StringValueToDate(types.StringValue(dateString))
	assert.False(t, d.HasError())
	assert.NotNil(t, date)
	assert.Equal(t, dateString, date.String())

	tfValue = DateToStringValue(date)
	assert.NotNil(t, tfValue)
	assert.Equal(t, dateString, tfValue.ValueString())

	// invalid date value.
	dateString = "bad date"
	date, d = StringValueToDate(types.StringValue(dateString))
	assert.True(t, d.HasError())
	assert.Nil(t, date)

	// nil slice
	lv, d := DateSliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToDateSlice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []strfmt.Date{}
	lv, d = DateSliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	// Create a list containing our expected date strings.
	dateStrings := []string{"2024-01-01", "2025-01-01"}
	lv, d = StringSliceToListValue(dateStrings)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)

	// Test list -> date slice.
	actualDates, d := ListValueToDateSlice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualDates)
	assert.Equal(t, len(dateStrings), len(actualDates))
	for i, _ := range dateStrings {
		assert.Equal(t, dateStrings[i], actualDates[i].String())
	}

	// Test date slice -> list.
	lv, d = DateSliceToListValue(actualDates)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)
	assert.Equal(t, len(actualDates), len(lv.Elements()))
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	actualStrings, d := ListValueToStringSlice(lv)
	assert.False(t, d.HasError())
	for i, _ := range actualStrings {
		assert.Equal(t, dateStrings[i], actualStrings[i])
	}
}

func TestDataConversionsDateTime(t *testing.T) {
	// null string value
	nullStringValue := types.StringNull()
	dateTime, d := StringValueToDateTime(nullStringValue)
	assert.False(t, d.HasError())
	assert.Nil(t, dateTime)

	// nil datetime value
	tfValue := DateTimeToStringValue(nil)
	assert.NotNil(t, tfValue)
	assert.True(t, tfValue.IsNull())

	// valid datetime value.
	dtString := "2025-01-20T12:00:00.198Z"
	dateTime, d = StringValueToDateTime(types.StringValue(dtString))
	assert.False(t, d.HasError())
	assert.NotNil(t, dateTime)
	assert.Equal(t, dtString, dateTime.String())

	tfValue = DateTimeToStringValue(dateTime)
	assert.NotNil(t, tfValue)
	assert.Equal(t, dtString, tfValue.ValueString())

	// invalid datetime value.
	dtString = "bad datetime"
	dateTime, d = StringValueToDateTime(types.StringValue(dtString))
	assert.True(t, d.HasError())
	assert.Nil(t, dateTime)

	// nil slice
	lv, d := DateTimeSliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToDateTimeSlice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []strfmt.DateTime{}
	lv, d = DateTimeSliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	// Create a list containing our expected datetime strings.
	dtStrings := []string{"2024-01-01T23:59:59.999Z", "2025-01-01T12:01:02.001Z"}
	lv, d = StringSliceToListValue(dtStrings)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)

	// Test list -> datetime slice.
	actualDateTimes, d := ListValueToDateTimeSlice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualDateTimes)
	assert.Equal(t, len(dtStrings), len(actualDateTimes))
	for i, _ := range dtStrings {
		assert.Equal(t, dtStrings[i], actualDateTimes[i].String())
	}

	// Test datetime slice -> list.
	lv, d = DateTimeSliceToListValue(actualDateTimes)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)
	assert.Equal(t, len(actualDateTimes), len(lv.Elements()))
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	actualStrings, d := ListValueToStringSlice(lv)
	assert.False(t, d.HasError())
	for i, _ := range actualStrings {
		assert.Equal(t, dtStrings[i], actualStrings[i])
	}
}

func TestDataConversionsUUID(t *testing.T) {
	// null string value
	nullStringValue := types.StringNull()
	uuid, d := StringValueToUUID(nullStringValue)
	assert.False(t, d.HasError())
	assert.Nil(t, uuid)

	// nil uuid value
	tfValue := UUIDToStringValue(nil)
	assert.NotNil(t, tfValue)
	assert.True(t, tfValue.IsNull())

	// valid uuid value.
	uuidString := "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
	uuid, d = StringValueToUUID(types.StringValue(uuidString))
	assert.False(t, d.HasError())
	assert.NotNil(t, uuid)
	assert.Equal(t, uuidString, uuid.String())

	tfValue = UUIDToStringValue(uuid)
	assert.NotNil(t, tfValue)
	assert.Equal(t, uuidString, tfValue.ValueString())

	// nil slice
	lv, d := UUIDSliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	slice, d := ListValueToUUIDSlice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []strfmt.UUID{}
	lv, d = UUIDSliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	// Create a list containing our expected uuid strings.
	uuidStrings := []string{"9fab83da-98cb-4f18-a7ba-b6f0435c9673", "aaffca34-de6d-11ea-87d0-0242ac130003"}
	lv, d = StringSliceToListValue(uuidStrings)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)

	// Test list -> date slice.
	actualUuids, d := ListValueToUUIDSlice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualUuids)
	assert.Equal(t, len(uuidStrings), len(actualUuids))
	for i, _ := range uuidStrings {
		assert.Equal(t, uuidStrings[i], actualUuids[i].String())
	}

	// Test uuid slice -> list.
	lv, d = UUIDSliceToListValue(actualUuids)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)
	assert.Equal(t, len(actualUuids), len(lv.Elements()))
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	actualStrings, d := ListValueToStringSlice(lv)
	assert.False(t, d.HasError())
	for i, _ := range actualStrings {
		assert.Equal(t, uuidStrings[i], actualStrings[i])
	}
}

func TestDataConversionsBool(t *testing.T) {
	// nil slice
	lv, d := BoolSliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToBoolSlice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []bool{}
	lv, d = BoolSliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	expectedSlice = []bool{false, true, true}
	lv, d = BoolSliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	actualSlice, d := ListValueToBoolSlice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)

	lvNull := basetypes.NewListNull(types.BoolType)
	assert.True(t, lvNull.IsNull())
	actualSlice, d = ListValueToBoolSlice(lvNull)
	assert.False(t, d.HasError())
	assert.Nil(t, actualSlice)
}

func TestDataConversionsInt32(t *testing.T) {
	// nil slice
	lv, d := Int32SliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToInt32Slice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []int32{}
	lv, d = Int32SliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	expectedSlice = []int32{33, 44, 74}
	lv, d = Int32SliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	actualSlice, d := ListValueToInt32Slice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)

	lvNull := basetypes.NewListNull(types.Int32Type)
	assert.True(t, lvNull.IsNull())
	actualSlice, d = ListValueToInt32Slice(lvNull)
	assert.False(t, d.HasError())
	assert.Nil(t, actualSlice)
}

func TestDataConversionsInt64(t *testing.T) {
	// nil slice
	lv, d := Int64SliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToInt64Slice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []int64{}
	lv, d = Int64SliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	expectedSlice = []int64{33, 44, 74}
	lv, d = Int64SliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	actualSlice, d := ListValueToInt64Slice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)

	lvNull := basetypes.NewListNull(types.Int64Type)
	assert.True(t, lvNull.IsNull())
	actualSlice, d = ListValueToInt64Slice(lvNull)
	assert.False(t, d.HasError())
	assert.Nil(t, actualSlice)
}

func TestDataConversionsFloat32(t *testing.T) {
	// nil slice
	lv, d := Float32SliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToFloat32Slice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []float32{}
	lv, d = Float32SliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	expectedSlice = []float32{33.1, 44.99, 74.1234}
	lv, d = Float32SliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	actualSlice, d := ListValueToFloat32Slice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)

	lvNull := basetypes.NewListNull(types.Float32Type)
	assert.True(t, lvNull.IsNull())
	actualSlice, d = ListValueToFloat32Slice(lvNull)
	assert.False(t, d.HasError())
	assert.Nil(t, actualSlice)
}

func TestDataConversionsFloat64(t *testing.T) {
	// nil slice
	lv, d := Float64SliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToFloat64Slice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	expectedSlice := []float64{}
	lv, d = Float64SliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	// non-empty slice
	expectedSlice = []float64{33.1, 44.99, 74.1234}
	lv, d = Float64SliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	actualSlice, d := ListValueToFloat64Slice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)

	lvNull := basetypes.NewListNull(types.Float64Type)
	assert.True(t, lvNull.IsNull())
	actualSlice, d = ListValueToFloat64Slice(lvNull)
	assert.False(t, d.HasError())
	assert.Nil(t, actualSlice)
}

func TestDataConversionsByteArray(t *testing.T) {
	expectedBytes := []byte("This is a test of the emergency broadcast system!")
	expectedString := base64.StdEncoding.EncodeToString(expectedBytes)
	assert.NotNil(t, expectedString)

	stringValue := types.StringValue(expectedString)
	actualBytes, d := StringValueToByteArray(stringValue)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualBytes)
	assert.Equal(t, expectedBytes, *actualBytes)

	actualString := ByteArrayToStringValue(&expectedBytes)
	assert.Equal(t, expectedString, actualString.ValueString())

	// Test with null values.
	ba, d := StringValueToByteArray(types.StringNull())
	assert.False(t, d.HasError())
	assert.Nil(t, ba)

	s := ByteArrayToStringValue(nil)
	assert.NotNil(t, s)
	assert.True(t, s.IsNull())

	// Test for errors.
	badEncodedString := "This is an invalid base64-encoded string."
	ba, d = StringValueToByteArray(types.StringValue(badEncodedString))
	assert.True(t, d.HasError())
	assert.Nil(t, ba)

	// nil slice
	lv, d := ByteArraySliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice, d := ListValueToByteArraySlice(lv)
	assert.False(t, d.HasError())
	assert.Nil(t, slice)

	// empty slice
	emptySlice := [][]byte{}
	lv, d = ByteArraySliceToListValue(emptySlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	actualSlice, d := ListValueToByteArraySlice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, 0, len(actualSlice))

	// non-empty slice
	expectedSlice := [][]byte{[]byte("foo"), []byte("bar"), []byte("this is a test")}
	expectedStringSlice := make([]string, len(expectedSlice))
	for i, ba := range expectedSlice {
		expectedStringSlice[i] = base64.StdEncoding.EncodeToString(ba)
	}

	lv, d = ByteArraySliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, len(expectedSlice), len(lv.Elements()))
	for i, elem := range lv.Elements() {
		assert.Equal(t, types.StringValue(expectedStringSlice[i]), elem)
	}

	actualSlice, d = ListValueToByteArraySlice(lv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)
}

func TestDataConversionsAnyObject(t *testing.T) {
	// nil any object
	// Make sure we get back a "null" map value.
	mv, d := AnyObjectToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())

	// Make sure we get back nil for a "null" map value.
	actualAnyObj := MapValueToAnyObject(mv)
	assert.Nil(t, actualAnyObj)

	// empty any object
	emptyAnyObj := map[string]any{}
	mv, d = AnyObjectToMapValue(emptyAnyObj)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualAnyObj = MapValueToAnyObject(mv)
	assert.NotNil(t, actualAnyObj)
	assert.Equal(t, 0, len(actualAnyObj))

	// non-empty any object
	anyObj := map[string]any{
		"p1": "hello",
		"p2": "world",
		"p3": false,
		"p4": 123.0001,
	}
	mv, d = AnyObjectToMapValue(anyObj)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	actualAnyObj = MapValueToAnyObject(mv)
	assert.NotNil(t, actualAnyObj)

	// We expect to get back a map with all string values.
	// This is because we convert each any value to a string
	// when converting to the terraform map value.
	expectedAnyObj := map[string]any{
		"p1": "hello",
		"p2": "world",
		"p3": "false",
		"p4": "123.0001",
	}
	assert.Equal(t, expectedAnyObj, actualAnyObj)

	// nil slice.
	lv, d := AnyObjectSliceToListValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, lv)
	assert.True(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())

	slice := ListValueToAnyObjectSlice(lv)
	assert.Nil(t, slice)

	// empty slice.
	emptySlice := []map[string]any{}
	lv, d = AnyObjectSliceToListValue(emptySlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, 0, len(lv.Elements()))

	actualSlice := ListValueToAnyObjectSlice(lv)
	assert.NotNil(t, actualSlice)
	assert.Equal(t, 0, len(actualSlice))

	// non-empty slice.
	expectedSlice := []map[string]any{
		{
			"p1": "hello",
			"p2": "world",
		},
		{
			"p3": "foo",
			"p4": "bar",
		},
	}

	lv, d = AnyObjectSliceToListValue(expectedSlice)
	assert.False(t, d.HasError())
	assert.False(t, lv.IsNull())
	assert.False(t, lv.IsUnknown())
	assert.Equal(t, len(expectedSlice), len(lv.Elements()))

	actualSlice = ListValueToAnyObjectSlice(lv)
	assert.NotNil(t, actualSlice)
	assert.Equal(t, expectedSlice, actualSlice)
}

func TestDataConversionsStringMap(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := StringMapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToStringMap(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string]string{}
	mv, d = StringMapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToStringMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	expectedMap := map[string]string{
		"p1": "hello",
		"p2": "world",
	}
	mv, d = StringMapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	actualMap = map[string]string{}
	d = mv.ElementsAs(context.Background(), &actualMap, false)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMap, actualMap)

	actualMap, d = MapValueToStringMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

func TestDataConversionsBoolMap(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := BoolMapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToBoolMap(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string]bool{}
	mv, d = BoolMapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToBoolMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	expectedMap := map[string]bool{
		"p1": false,
		"p2": true,
	}
	mv, d = BoolMapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	actualMap = map[string]bool{}
	d = mv.ElementsAs(context.Background(), &actualMap, false)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMap, actualMap)

	actualMap, d = MapValueToBoolMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

func TestDataConversionsInt64Map(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := Int64MapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToInt64Map(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string]int64{}
	mv, d = Int64MapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToInt64Map(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	expectedMap := map[string]int64{
		"p1": 44,
		"p2": 74,
	}
	mv, d = Int64MapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	actualMap = map[string]int64{}
	d = mv.ElementsAs(context.Background(), &actualMap, false)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMap, actualMap)

	actualMap, d = MapValueToInt64Map(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

func TestDataConversionsFloat64Map(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := Float64MapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToFloat64Map(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string]float64{}
	mv, d = Float64MapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToFloat64Map(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	expectedMap := map[string]float64{
		"p1": float64(44.123456789),
		"p2": float64(74.9999999),
	}
	mv, d = Float64MapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	actualMap = map[string]float64{}
	d = mv.ElementsAs(context.Background(), &actualMap, false)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMap, actualMap)

	actualMap, d = MapValueToFloat64Map(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

func TestDataConversionsFloat32Map(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := Float32MapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToFloat32Map(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string]float32{}
	mv, d = Float32MapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToFloat32Map(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	expectedMap := map[string]float32{
		"p1": 44.123456789,
		"p2": 74.9999999,
	}
	mv, d = Float32MapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	actualMap = map[string]float32{}
	d = mv.ElementsAs(context.Background(), &actualMap, false)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMap, actualMap)

	actualMap, d = MapValueToFloat32Map(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

func TestDataConversionsDateMap(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := DateMapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToDateMap(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string]strfmt.Date{}
	mv, d = DateMapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToDateMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	d1, d := StringValueToDate(types.StringValue("2004-10-24"))
	assert.NotNil(t, d1)
	assert.False(t, d.HasError())

	d2, d := StringValueToDate(types.StringValue("2025-06-30"))
	assert.NotNil(t, d2)
	assert.False(t, d.HasError())

	expectedMap := map[string]strfmt.Date{
		"p1": *d1,
		"p2": *d2,
	}

	expectedMapStr := map[string]string{
		"p1": "2004-10-24",
		"p2": "2025-06-30",
	}

	mv, d = DateMapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	strMap, d := MapValueToStringMap(mv)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMapStr, strMap)

	actualMap, d = MapValueToDateMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

func TestDataConversionsDateTimeMap(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := DateTimeMapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToDateTimeMap(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string]strfmt.DateTime{}
	mv, d = DateTimeMapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToDateTimeMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	// "2024-01-01T23:59:59.999Z", "2025-01-01T12:01:02.001Z"
	d1, d := StringValueToDateTime(types.StringValue("2024-01-01T23:59:59.999Z"))
	assert.NotNil(t, d1)
	assert.False(t, d.HasError())

	d2, d := StringValueToDateTime(types.StringValue("2025-01-01T12:01:02.001Z"))
	assert.NotNil(t, d2)
	assert.False(t, d.HasError())

	expectedMap := map[string]strfmt.DateTime{
		"p1": *d1,
		"p2": *d2,
	}

	expectedMapStr := map[string]string{
		"p1": "2024-01-01T23:59:59.999Z",
		"p2": "2025-01-01T12:01:02.001Z",
	}

	mv, d = DateTimeMapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	strMap, d := MapValueToStringMap(mv)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMapStr, strMap)

	actualMap, d = MapValueToDateTimeMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

func TestDataConversionsUUIDMap(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := UUIDMapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToUUIDMap(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string]strfmt.UUID{}
	mv, d = UUIDMapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToUUIDMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	u1, d := StringValueToUUID(types.StringValue("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
	assert.NotNil(t, u1)
	assert.False(t, d.HasError())

	u2, d := StringValueToUUID(types.StringValue("aaffca34-de6d-11ea-87d0-0242ac130003"))
	assert.NotNil(t, u2)
	assert.False(t, d.HasError())

	expectedMap := map[string]strfmt.UUID{
		"p1": *u1,
		"p2": *u2,
	}

	expectedMapStr := map[string]string{
		"p1": "9fab83da-98cb-4f18-a7ba-b6f0435c9673",
		"p2": "aaffca34-de6d-11ea-87d0-0242ac130003",
	}

	mv, d = UUIDMapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	strMap, d := MapValueToStringMap(mv)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMapStr, strMap)

	actualMap, d = MapValueToUUIDMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

func TestDataConversionsByteArrayMap(t *testing.T) {
	// nil map
	// Make sure we get back a "null" map value.
	mv, d := ByteArrayMapToMapValue(nil)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.True(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())

	// Make sure we get back nil for a "null" map value.
	m, d := MapValueToByteArrayMap(mv)
	assert.False(t, d.HasError())
	assert.Nil(t, m)

	// empty map
	emptyMap := map[string][]byte{}
	mv, d = ByteArrayMapToMapValue(emptyMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)
	assert.False(t, mv.IsNull())
	assert.False(t, mv.IsUnknown())
	assert.Equal(t, 0, len(mv.Elements()))

	actualMap, d := MapValueToByteArrayMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, 0, len(actualMap))

	// non-empty map
	ba1 := []byte("This is a test of the ")
	ba2 := []byte("emergency broadcast system!")
	expectedMap := map[string][]byte{
		"p1": ba1,
		"p2": ba2,
	}

	expectedMapStr := map[string]string{
		"p1": "VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIA==",
		"p2": "ZW1lcmdlbmN5IGJyb2FkY2FzdCBzeXN0ZW0h",
	}

	mv, d = ByteArrayMapToMapValue(expectedMap)
	assert.False(t, d.HasError())
	assert.NotNil(t, mv)

	strMap, d := MapValueToStringMap(mv)
	assert.False(t, d.HasError())
	assert.Equal(t, expectedMapStr, strMap)

	actualMap, d = MapValueToByteArrayMap(mv)
	assert.False(t, d.HasError())
	assert.NotNil(t, actualMap)
	assert.Equal(t, expectedMap, actualMap)
}

//
// When diagnosing test failures, it might be useful to uncomment
// this function, then invoke it from within one of the test functions above.
//
// func displayDiags(t *testing.T, d diag.Diagnostics) {
// 	t.Logf("Diagnostics entries:\n")
// 	for i, entry := range d {
// 		t.Logf("[%d]: [%s] %s: %s\n", i, entry.Severity().String(), entry.Summary(), entry.Detail())
// 	}
// }
