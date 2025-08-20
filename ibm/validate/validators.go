// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package validate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func validateAllowedStringValue(validValues []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		input := v.(string)
		existed := false
		for _, s := range validValues {
			if s == input {
				existed = true
				break
			}
		}
		if !existed {
			errors = append(errors, fmt.Errorf(
				"%q must contain a value from %#v, got %q",
				k, validValues, input))
		}
		return

	}
}

func validateRegexpLen(min, max int, regex string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)

		acceptedcharacters, _ := regexp.MatchString(regex, value)

		if acceptedcharacters {
			if (len(value) < min) || (len(value) > max) && (min > 0 && max > 0) {
				errors = append(errors, fmt.Errorf(
					"%q (%q) must contain from %d to %d characters ", k, value, min, max))
			}
		} else {
			errors = append(errors, fmt.Errorf(
				"%q (%q) should match regexp %s ", k, v, regex))
		}

		return

	}
}

func ValidateAllowedIntValue(is []int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		existed := false
		for _, i := range is {
			if i == value {
				existed = true
				break
			}
		}
		if !existed {
			errors = append(errors, fmt.Errorf(
				"%q must contain a valid int value should in array %#v, got %q",
				k, is, value))
		}
		return

	}
}

func validateRegexp(regex string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)

		acceptedcharacters, _ := regexp.MatchString(regex, value)

		if !acceptedcharacters {
			errors = append(errors, fmt.Errorf(
				"%q (%q) should match regexp %s ", k, v, regex))
		}

		return

	}
}

// ValidateFunc is honored only when the schema's Type is set to TypeInt,
// TypeFloat, TypeString, TypeBool, or TypeMap. It is ignored for all other types.
// enum to list all the validator functions supported by this tool.
type FunctionIdentifier int

const (
	IntBetween FunctionIdentifier = iota
	ValidateAllowedStringValue
	StringLenBetween
	ValidateRegexpLen
	ValidateRegexp
)

// ValueType -- Copied from Terraform for now. You can refer to Terraform ValueType directly.
// ValueType is an enum of the type that can be represented by a schema.
type ValueType int

const (
	TypeInvalid ValueType = iota
	TypeBool
	TypeInt
	TypeFloat
	TypeString
)

// Type of constraints required for validation
type ValueConstraintType int

const (
	MinValue ValueConstraintType = iota
	MaxValue
	MinValueLength
	MaxValueLength
	AllowedValues
	MatchesValue
)

// Schema is used to describe the validation schema.
type ValidateSchema struct {

	//This is the parameter name.
	//Ex: private_subnet in ibm_compute_bare_metal resource
	Identifier string

	// this is similar to schema.ValueType
	Type ValueType

	// The actual validation function that needs to be invoked.
	// Ex: IntBetween, validateAllowedIntValue, validateAllowedStringValue
	ValidateFunctionIdentifier FunctionIdentifier

	MinValue       string
	MaxValue       string
	AllowedValues  string //Comma separated list of strings.
	Matches        string
	Regexp         string
	MinValueLength int
	MaxValueLength int

	// Is this nullable
	Nullable bool

	Optional bool
	Required bool
	Default  interface{}
	ForceNew bool
}

type ResourceValidator struct {
	// This is the resource name - Found in provider.go of IBM Terraform provider.
	// Ex: ibm_compute_monitor, ibm_compute_bare_metal, ibm_compute_dedicated_host, ibm_cis_global_load_balancer etc.,
	ResourceName string

	// Array of validator objects. Each object refers to one parameter in the resource provider.
	Schema []ValidateSchema
}

type ValidatorDict struct {
	ResourceValidatorDictionary   map[string]*ResourceValidator
	DataSourceValidatorDictionary map[string]*ResourceValidator
}

// Resource Validator Dictionary -- For all terraform IBM Resource Providers.
// This is of type - Array of ResourceValidators.
// Each object in this array is a type of map, where key == ResourceName and value == array of ValidateSchema objects. Each of these
// ValidateSchema corresponds to a parameter in the resourceProvider.

var validatorDict ValidatorDict

func SetValidatorDict(v ValidatorDict) {
	validatorDict = v
}

// This is the main validation function. This function will be used in all the provider code.
func InvokeValidator(resourceName, identifier string) schema.SchemaValidateFunc {
	// Loop through dictionary and identify the resource and then the parameter configuration.
	var schemaToInvoke ValidateSchema
	found := false
	resourceItem := validatorDict.ResourceValidatorDictionary[resourceName]
	if resourceItem.ResourceName == resourceName {
		parameterValidateSchema := resourceItem.Schema
		for _, validateSchema := range parameterValidateSchema {
			if validateSchema.Identifier == identifier {
				schemaToInvoke = validateSchema
				found = true
				break
			}
		}
	}

	if found {
		return invokeValidatorInternal(schemaToInvoke)
	} else {
		// Add error code later. TODO
		return nil
	}
}

func InvokeDataSourceValidator(resourceName, identifier string) schema.SchemaValidateFunc {
	// Loop through dictionary and identify the resource and then the parameter configuration.
	var schemaToInvoke ValidateSchema
	found := false

	dataSourceItem := validatorDict.DataSourceValidatorDictionary[resourceName]
	if dataSourceItem.ResourceName == resourceName {
		parameterValidateSchema := dataSourceItem.Schema
		for _, validateSchema := range parameterValidateSchema {
			if validateSchema.Identifier == identifier {
				schemaToInvoke = validateSchema
				found = true
				break
			}
		}
	}

	if found {
		return invokeValidatorInternal(schemaToInvoke)
	} else {
		// Add error code later. TODO
		return nil
	}
}

// the function is currently modified to invoke SchemaValidateFunc directly.
// But in terraform, we will just return SchemaValidateFunc as shown below.. So terraform will invoke this func
func invokeValidatorInternal(schema ValidateSchema) schema.SchemaValidateFunc {

	funcIdentifier := schema.ValidateFunctionIdentifier
	switch funcIdentifier {
	case IntBetween:
		minValue := schema.GetValue(MinValue)
		maxValue := schema.GetValue(MaxValue)
		return validation.IntBetween(minValue.(int), maxValue.(int))
	case ValidateAllowedStringValue:
		allowedValues := schema.GetValue(AllowedValues)
		return validateAllowedStringValue(allowedValues.([]string))
	case StringLenBetween:
		return validation.StringLenBetween(schema.MinValueLength, schema.MaxValueLength)
	case ValidateRegexpLen:
		return validateRegexpLen(schema.MinValueLength, schema.MaxValueLength, schema.Regexp)
	case ValidateRegexp:
		return validateRegexp(schema.Regexp)

	default:
		return nil
	}
}

// utility functions - Move to different package
func (vs ValidateSchema) GetValue(valueConstraint ValueConstraintType) interface{} {

	var valueToConvert string
	switch valueConstraint {
	case MinValue:
		valueToConvert = vs.MinValue
	case MaxValue:
		valueToConvert = vs.MaxValue
	case AllowedValues:
		valueToConvert = vs.AllowedValues
	case MatchesValue:
		valueToConvert = vs.Matches
	}

	switch vs.Type {
	case TypeInvalid:
		return nil
	case TypeBool:
		b, err := strconv.ParseBool(valueToConvert)
		if err != nil {
			return vs.Zero()
		}
		return b
	case TypeInt:
		// Convert comma separated string to array
		if strings.Contains(valueToConvert, ",") {
			var arr2 []int
			arr1 := strings.Split(valueToConvert, ",")
			for _, ele := range arr1 {
				e, err := strconv.Atoi(strings.TrimSpace(ele))
				if err != nil {
					return vs.Zero()
				}
				arr2 = append(arr2, e)
			}
			return arr2
		} else {
			num, err := strconv.Atoi(valueToConvert)
			if err != nil {
				return vs.Zero()
			}
			return num
		}

	case TypeFloat:
		f, err := strconv.ParseFloat(valueToConvert, 32)
		if err != nil {
			return vs.Zero()
		}
		return f
	case TypeString:
		//return valueToConvert
		// Convert comma separated string to array
		arr := strings.Split(valueToConvert, ",")
		for i, ele := range arr {
			arr[i] = strings.TrimSpace(ele)
		}
		return arr
	default:
		panic(fmt.Sprintf("unknown type %s", vs.Type))
	}
}

// Use stringer tool to generate this later.
func (i FunctionIdentifier) String() string {
	return [...]string{"IntBetween", "IntAtLeast", "IntAtMost"}[i]
}

// Use Stringer tool to generate this later.
func (i ValueType) String() string {
	return [...]string{"TypeInvalid", "TypeBool", "TypeInt", "TypeFloat", "TypeString"}[i]
}

// Use Stringer tool to generate this later.
func (i ValueConstraintType) String() string {
	return [...]string{"MinValue", "MaxValue", "MinValueLength", "MaxValueLength", "AllowedValues", "MatchesValue"}[i]
}

// Zero returns the zero value for a type.
func (vs ValidateSchema) Zero() interface{} {
	switch vs.Type {
	case TypeInvalid:
		return nil
	case TypeBool:
		return false
	case TypeInt:
		return make([]string, 0)
	case TypeFloat:
		return 0.0
	case TypeString:
		return make([]int, 0)
	default:
		panic(fmt.Sprintf("unknown type %s", vs.Type))
	}
}
