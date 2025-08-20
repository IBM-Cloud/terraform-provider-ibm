// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

import (
	"context"
	"encoding/base64"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Converts string slice "s" into a Terraform types.List containing the string values.
func StringSliceToListValue(s []string) (types.List, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewListNull(types.StringType), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = basetypes.NewStringValue(elem)
	}
	return basetypes.NewListValue(types.StringType, values)
}

// Converts a Terraform types.List containing strings to a string slice.
func ListValueToStringSlice(l types.List) ([]string, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	var s []string
	diags := l.ElementsAs(context.Background(), &s, false)
	return s, diags
}

// Converts string slice "s" into a Terraform types.Set containing the string values.
func StringSliceToSetValue(s []string) (types.Set, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewSetNull(types.StringType), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = basetypes.NewStringValue(elem)
	}
	return basetypes.NewSetValue(types.StringType, values)
}

// Converts a Terraform types.Set containing strings to a string slice.
func SetValueToStringSlice(l types.Set) ([]string, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	var s []string
	diags := l.ElementsAs(context.Background(), &s, false)
	return s, diags
}

// Converts map "m" from its Go SDK representation (map[string]string) into a Terraform Map value.
func StringMapToMapValue(m map[string]string) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return types.MapNull(types.StringType), nil
	}

	return types.MapValueFrom(context.Background(), types.StringType, m)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string]string).
func MapValueToStringMap(m types.Map) (map[string]string, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	result := map[string]string{}
	diags := m.ElementsAs(context.Background(), &result, false)
	return result, diags
}

// Converts "a" (an any) into a Terraform string value
// by using the Stringify() function.
func AnyToStringValue(a any) types.String {
	if core.IsNil(a) {
		return types.StringNull()
	}

	return types.StringValue(Stringify(a))
}

// Converts "s" (Terraform string value) into an any.
func StringValueToAny(s types.String) any {
	if s.IsNull() {
		return nil
	}

	return s.ValueString()
}

// Converts slice "s" into a Terraform types.List containing
// the any values as strings.  The flex.Stringify() function is
// used to convert each "any" value into a string.
func AnySliceToListValue(s []any) (types.List, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewListNull(types.StringType), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = basetypes.NewStringValue(Stringify(elem))
	}
	return basetypes.NewListValue(types.StringType, values)
}

// Converts list value "l" (containing strings) into a slice of any.
func ListValueToAnySlice(l types.List) ([]any, diag.Diagnostics) {
	// Get the string slice from the list value.
	s, diags := ListValueToStringSlice(l)
	if diags.HasError() {
		return nil, diags
	}

	if s == nil {
		return nil, nil
	}

	// Now create an any slice containing the strings.
	// We do this so that we end up with the proper typing.
	a := make([]any, len(s))
	for i, elem := range s {
		a[i] = elem
	}

	return a, nil
}

// Converts "anyObj" (an "any object") from its Go SDK representation (map[string]any)
// into a Terraform map whose values are strings (i.e. any is mapped to string in Terraform).
func AnyObjectToMapValue(anyObj map[string]any) (types.Map, diag.Diagnostics) {
	if core.IsNil(anyObj) {
		return types.MapNull(types.StringType), nil
	}

	// Convert anyObj's values to string
	a := make(map[string]string, len(anyObj))
	for key, value := range anyObj {
		sv := AnyToStringValue(value)
		a[key] = sv.ValueString()
	}

	return types.MapValueFrom(context.Background(), types.StringType, a)
}

// Converts "tfMap" (a Terraform map value) from it's Terraform representation
// (map[string]string) into its Go SDK representation (map[string]any).
func MapValueToAnyObject(tfMap types.Map) map[string]any {
	if tfMap.IsNull() || tfMap.IsUnknown() {
		return nil
	}

	// Grab the elements as a map[string]attr.Value.
	rawMap := tfMap.Elements()

	// Create a map to hold the SDK representation.
	sdkMap := make(map[string]any, len(rawMap))
	for key, elem := range rawMap {
		v := StringValueToAny(elem.(types.String))
		sdkMap[key] = v
	}
	return sdkMap
}

// Converts slice "s" into a Terraform types.List containing "any object" values.
func AnyObjectSliceToListValue(s []map[string]any) (types.List, diag.Diagnostics) {
	// Each list value is a map (any object is a map[string]any).
	listElementType := types.MapType{
		ElemType: types.StringType,
	}

	if core.IsNil(s) {
		return basetypes.NewListNull(listElementType), nil
	}

	values := make([]attr.Value, len(s))
	for i := range s {
		var diags diag.Diagnostics
		values[i], diags = AnyObjectToMapValue(s[i])
		if diags.HasError() {
			return basetypes.NewListNull(listElementType), diags
		}
	}
	return basetypes.NewListValue(listElementType, values)
}

// Converts list value "l" into a slice of "any object" values.
func ListValueToAnyObjectSlice(l types.List) []map[string]any {
	if l.IsNull() || l.IsUnknown() {
		return nil
	}

	elements := l.Elements()
	s := make([]map[string]any, len(elements))
	for i, elem := range elements {
		s[i] = MapValueToAnyObject(elem.(types.Map))
	}
	return s
}

// Decodes "s" (a Terraform string value) into a byte array using
// a standard base64 decoder.  The input string "s" must be a
// valid base64-encoded string value.
func StringValueToByteArray(s types.String) (*[]byte, diag.Diagnostics) {
	if s.IsNull() {
		return nil, nil
	}
	ba, err := base64.StdEncoding.DecodeString(s.ValueString())
	if err != nil {
		var diags diag.Diagnostics
		diags.AddError("error decoding byte-array string", err.Error())
		return nil, diags
	}
	return &ba, nil
}

// Encodes "ba" (a byte array) into a base64-encoded string
// using a standard base64 encoder.
func ByteArrayToStringValue(ba *[]byte) types.String {
	if ba == nil {
		return types.StringNull()
	}

	s := base64.StdEncoding.EncodeToString(*ba)
	return types.StringValue(s)
}

// Encodes a slice of byte arrays into a types.List containing the
// base64-encoded strings using a standard base64 encoder.
func ByteArraySliceToListValue(s [][]byte) (types.List, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewListNull(types.StringType), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = ByteArrayToStringValue(&elem)
	}
	return basetypes.NewListValue(types.StringType, values)
}

// Decodes the base64-encoded strings contained in "l" into
// a slice of byte arrays using a standard base64 decoder.
func ListValueToByteArraySlice(l types.List) ([][]byte, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	s, diags := ListValueToStringSlice(l)
	if diags.HasError() {
		return nil, diags
	}

	baSlice := make([][]byte, len(s))
	for i, elem := range s {
		b, diags := StringValueToByteArray(types.StringValue(elem))
		if diags.HasError() {
			return nil, diags
		}
		baSlice[i] = *b
	}

	return baSlice, nil
}

// Converts map "m" from its Go SDK representation (map[string][]byte) into a Terraform Map value.
func ByteArrayMapToMapValue(m map[string][]byte) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return basetypes.NewMapNull(types.StringType), nil
	}

	tfMap := make(map[string]attr.Value, len(m))
	for k, v := range m {
		tfMap[k] = ByteArrayToStringValue(&v)
	}

	return types.MapValueFrom(context.Background(), types.StringType, tfMap)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string][]byte).
func MapValueToByteArrayMap(m types.Map) (map[string][]byte, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	tempMap, diags := MapValueToStringMap(m)
	if diags.HasError() {
		return nil, diags
	}

	sdkMap := make(map[string][]byte, len(tempMap))
	for k, v := range tempMap {
		ba, diags := StringValueToByteArray(types.StringValue(v))
		if diags.HasError() {
			return nil, diags
		}
		sdkMap[k] = *ba
	}
	return sdkMap, diags
}

// Converts string value "s" into a Date instance,
// and returns an error if "s" cannot be correctly parsed.
func StringValueToDate(s types.String) (*strfmt.Date, diag.Diagnostics) {
	if s.IsNull() {
		return nil, nil
	}

	d, err := core.ParseDate(s.ValueString())
	if err != nil {
		var diags diag.Diagnostics
		diags.AddError("error parsing date string", err.Error())
		return nil, diags
	}
	return &d, nil
}

// Converts "date" into a string value.
func DateToStringValue(date *strfmt.Date) types.String {
	if date == nil {
		return types.StringNull()
	}
	return types.StringValue(date.String())
}

// Converts "dates" (a Date slice) into a string slice by invoking each Date value's String() method.
// The strings are returned via a types.List instance to be used in a Terraform data model.
func DateSliceToListValue(dates []strfmt.Date) (types.List, diag.Diagnostics) {
	if core.IsNil(dates) {
		return basetypes.NewListNull(types.StringType), nil
	}

	values := make([]attr.Value, len(dates))
	for i, elem := range dates {
		values[i] = DateToStringValue(&elem)
	}
	return basetypes.NewListValue(types.StringType, values)
}

// Converts the strings contained in "l" into a Date slice
// by parsing each string into its corresponding Date value.
func ListValueToDateSlice(l types.List) ([]strfmt.Date, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	s, diags := ListValueToStringSlice(l)
	if diags.HasError() {
		return nil, diags
	}

	dateSlice := make([]strfmt.Date, len(s))
	for i, elem := range s {
		date, diags := StringValueToDate(types.StringValue(elem))
		if diags.HasError() {
			return nil, diags
		}
		dateSlice[i] = *date
	}

	return dateSlice, nil
}

// Converts map "m" from its Go SDK representation (map[string]strfmt.Date) into a Terraform Map value.
func DateMapToMapValue(m map[string]strfmt.Date) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return basetypes.NewMapNull(types.StringType), nil
	}

	tfMap := make(map[string]attr.Value, len(m))
	for k, v := range m {
		tfMap[k] = DateToStringValue(&v)
	}

	return types.MapValueFrom(context.Background(), types.StringType, tfMap)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string]strfmt.Date).
func MapValueToDateMap(m types.Map) (map[string]strfmt.Date, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	tempMap, diags := MapValueToStringMap(m)
	if diags.HasError() {
		return nil, diags
	}

	sdkMap := make(map[string]strfmt.Date, len(tempMap))
	for k, v := range tempMap {
		date, diags := StringValueToDate(types.StringValue(v))
		if diags.HasError() {
			return nil, diags
		}
		sdkMap[k] = *date
	}
	return sdkMap, diags
}

// Converts string value "s" into a DateTime instance,
// and returns an error if "s" cannot be correctly parsed.
func StringValueToDateTime(s types.String) (*strfmt.DateTime, diag.Diagnostics) {
	if s.IsNull() {
		return nil, nil
	}

	dt, err := core.ParseDateTime(s.ValueString())
	if err != nil {
		var diags diag.Diagnostics
		diags.AddError("error parsing date-time string", err.Error())
		return nil, diags
	}
	return &dt, nil
}

// Converts "datetime" into a string value.
func DateTimeToStringValue(datetime *strfmt.DateTime) types.String {
	if datetime == nil {
		return types.StringNull()
	}
	return types.StringValue(datetime.String())
}

// Converts "datetimes" (a DateTime slice) into a string slice by invoking each DateTime value's String() method.
// The strings are returned via a types.List instance to be used in a Terraform data model.
func DateTimeSliceToListValue(datetimes []strfmt.DateTime) (types.List, diag.Diagnostics) {
	if core.IsNil(datetimes) {
		return basetypes.NewListNull(types.StringType), nil
	}

	values := make([]attr.Value, len(datetimes))
	for i, elem := range datetimes {
		values[i] = DateTimeToStringValue(&elem)
	}
	return basetypes.NewListValue(types.StringType, values)
}

// Converts the strings contained in "l" into a DateTime slice
// by parsing each string into its corresponding DateTime value.
func ListValueToDateTimeSlice(l types.List) ([]strfmt.DateTime, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	s, diags := ListValueToStringSlice(l)
	if diags.HasError() {
		return nil, diags
	}

	dtSlice := make([]strfmt.DateTime, len(s))
	for i, elem := range s {
		dt, diags := StringValueToDateTime(types.StringValue(elem))
		if diags.HasError() {
			return nil, diags
		}
		dtSlice[i] = *dt
	}

	return dtSlice, nil
}

// Converts map "m" from its Go SDK representation (map[string]strfmt.DateTime) into a Terraform Map value.
func DateTimeMapToMapValue(m map[string]strfmt.DateTime) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return basetypes.NewMapNull(types.StringType), nil
	}

	tfMap := make(map[string]attr.Value, len(m))
	for k, v := range m {
		tfMap[k] = DateTimeToStringValue(&v)
	}

	return types.MapValueFrom(context.Background(), types.StringType, tfMap)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string]strfmt.DateTime).
func MapValueToDateTimeMap(m types.Map) (map[string]strfmt.DateTime, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	tempMap, diags := MapValueToStringMap(m)
	if diags.HasError() {
		return nil, diags
	}

	sdkMap := make(map[string]strfmt.DateTime, len(tempMap))
	for k, v := range tempMap {
		dateTime, diags := StringValueToDateTime(types.StringValue(v))
		if diags != nil {
			return nil, diags
		}
		sdkMap[k] = *dateTime
	}
	return sdkMap, diags
}

// Converts string value "s" into a UUID instance,
// and returns an error if "s" cannot be correctly parsed.
func StringValueToUUID(s types.String) (*strfmt.UUID, diag.Diagnostics) {
	if s.IsNull() {
		return nil, nil
	}

	uuid := strfmt.UUID(s.ValueString())
	return &uuid, nil
}

// Converts "uuid" into a string value.
func UUIDToStringValue(uuid *strfmt.UUID) types.String {
	if uuid == nil {
		return types.StringNull()
	}
	return types.StringValue(uuid.String())
}

// Converts "uuids" (a UUID slice) into a string slice by invoking each UUID value's String() method.
// The strings are returned via a types.List instance to be used in a Terraform data model.
func UUIDSliceToListValue(uuids []strfmt.UUID) (types.List, diag.Diagnostics) {
	if core.IsNil(uuids) {
		return basetypes.NewListNull(types.StringType), nil
	}

	values := make([]attr.Value, len(uuids))
	for i, elem := range uuids {
		values[i] = UUIDToStringValue(&elem)
	}
	return basetypes.NewListValue(types.StringType, values)
}

// Converts the strings contained in "l" into a UUID slice
// by parsing each string into its corresponding UUID value.
func ListValueToUUIDSlice(l types.List) ([]strfmt.UUID, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	s, diags := ListValueToStringSlice(l)
	if diags.HasError() {
		return nil, diags
	}

	uuidSlice := make([]strfmt.UUID, len(s))
	for i, elem := range s {
		uuid, diags := StringValueToUUID(types.StringValue(elem))
		if diags.HasError() {
			return nil, diags
		}
		uuidSlice[i] = *uuid
	}

	return uuidSlice, nil
}

// Converts map "m" from its Go SDK representation (map[string]strfmt.UUID) into a Terraform Map value.
func UUIDMapToMapValue(m map[string]strfmt.UUID) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return basetypes.NewMapNull(types.StringType), nil
	}

	tfMap := make(map[string]attr.Value, len(m))
	for k, v := range m {
		tfMap[k] = UUIDToStringValue(&v)
	}

	return types.MapValueFrom(context.Background(), types.StringType, tfMap)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string]strfmt.UUID).
func MapValueToUUIDMap(m types.Map) (map[string]strfmt.UUID, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	tempMap, diags := MapValueToStringMap(m)
	if diags.HasError() {
		return nil, diags
	}

	sdkMap := make(map[string]strfmt.UUID, len(tempMap))
	for k, v := range tempMap {
		uuid, diags := StringValueToUUID(types.StringValue(v))
		if diags.HasError() {
			return nil, diags
		}
		sdkMap[k] = *uuid
	}
	return sdkMap, diags
}

// Converts a bool slice "s" into a types.List containing the boolean values.
func BoolSliceToListValue(s []bool) (types.List, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewListNull(types.BoolType), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = basetypes.NewBoolValue(elem)
	}
	return basetypes.NewListValue(types.BoolType, values)
}

// Converts the bool values contained in "l" into a slice of bools.
func ListValueToBoolSlice(l types.List) ([]bool, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	var s []bool
	diags := l.ElementsAs(context.Background(), &s, false)
	return s, diags
}

// Converts map "m" from its Go SDK representation (map[string]bool) into a Terraform Map value.
func BoolMapToMapValue(m map[string]bool) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return types.MapNull(types.BoolType), nil
	}

	return types.MapValueFrom(context.Background(), types.BoolType, m)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string]bool).
func MapValueToBoolMap(m types.Map) (map[string]bool, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	result := map[string]bool{}
	diags := m.ElementsAs(context.Background(), &result, false)
	return result, diags
}

// Converts the int32 slice "s" into a types.List containing the int32 values.
func Int32SliceToListValue(s []int32) (types.List, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewListNull(types.Int32Type), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = basetypes.NewInt32Value(elem)
	}
	return basetypes.NewListValue(types.Int32Type, values)
}

// Converts the int32 values in "l" into a slice of int32.
func ListValueToInt32Slice(l types.List) ([]int32, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	var s []int32
	diags := l.ElementsAs(context.Background(), &s, false)
	return s, diags
}

// Converts the int64 slice "s" into a types.List containing the int64 values.
func Int64SliceToListValue(s []int64) (types.List, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewListNull(types.Int64Type), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = basetypes.NewInt64Value(elem)
	}
	return basetypes.NewListValue(types.Int64Type, values)
}

// Converts the int64 values in "l" into a slice of int64.
func ListValueToInt64Slice(l types.List) ([]int64, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	var s []int64
	diags := l.ElementsAs(context.Background(), &s, false)
	return s, diags
}

// Converts map "m" from its Go SDK representation (map[string]int64) into a Terraform Map value.
func Int64MapToMapValue(m map[string]int64) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return types.MapNull(types.Int64Type), nil
	}

	return types.MapValueFrom(context.Background(), types.Int64Type, m)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string]int64).
func MapValueToInt64Map(m types.Map) (map[string]int64, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	result := map[string]int64{}
	d := m.ElementsAs(context.Background(), &result, false)
	return result, d
}

// Converts the float32 slice "s" into a types.List containing the float32 values.
func Float32SliceToListValue(s []float32) (types.List, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewListNull(types.Float32Type), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = basetypes.NewFloat32Value(elem)
	}
	return basetypes.NewListValue(types.Float32Type, values)
}

// Converts the float32 values in "l" into a slice of float32.
func ListValueToFloat32Slice(l types.List) ([]float32, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	var s []float32
	diags := l.ElementsAs(context.Background(), &s, false)
	return s, diags
}

// Converts map "m" from its Go SDK representation (map[string]float32) into a Terraform Map value.
func Float32MapToMapValue(m map[string]float32) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return types.MapNull(types.Float32Type), nil
	}

	return types.MapValueFrom(context.Background(), types.Float32Type, m)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string]float32)
func MapValueToFloat32Map(m types.Map) (map[string]float32, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	result := map[string]float32{}
	diags := m.ElementsAs(context.Background(), &result, false)
	return result, diags
}

// Converts the float64 slice "s" into a types.List containing the float64 values.
func Float64SliceToListValue(s []float64) (types.List, diag.Diagnostics) {
	if core.IsNil(s) {
		return basetypes.NewListNull(types.Float64Type), nil
	}

	values := make([]attr.Value, len(s))
	for i, elem := range s {
		values[i] = basetypes.NewFloat64Value(elem)
	}
	return basetypes.NewListValue(types.Float64Type, values)
}

// Converts the float64 values in "l" into a slice of float64.
func ListValueToFloat64Slice(l types.List) ([]float64, diag.Diagnostics) {
	if l.IsNull() || l.IsUnknown() {
		return nil, nil
	}

	var s []float64
	diags := l.ElementsAs(context.Background(), &s, false)
	return s, diags
}

// Converts map "m" from its Go SDK representation (map[string]float64) into a Terraform Map value.
func Float64MapToMapValue(m map[string]float64) (types.Map, diag.Diagnostics) {
	if core.IsNil(m) {
		return types.MapNull(types.Float64Type), nil
	}

	return types.MapValueFrom(context.Background(), types.Float64Type, m)
}

// Converts Terraform Map value "m" into its Go SDK representation (map[string]float64)
func MapValueToFloat64Map(m types.Map) (map[string]float64, diag.Diagnostics) {
	if m.IsNull() || m.IsUnknown() {
		return nil, nil
	}

	result := map[string]float64{}
	diags := m.ElementsAs(context.Background(), &result, false)
	return result, diags
}
