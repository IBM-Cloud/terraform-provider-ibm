package errors

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.ibm.com/riaas/rias-api/riaas/models"
)

// ErrorTarget ...
type ErrorTarget struct {
	Name string
	Type string
}

// SingleError ...
type SingleError struct {
	Code     string
	Message  string
	MoreInfo string
	Target   ErrorTarget
}

// RiaasError ...
type RiaasError struct {
	Payload *models.Riaaserror
}

func (e RiaasError) Error() string {
	b, _ := json.Marshal(e.Payload)
	return string(b)
}

// ToError ...
func ToError(err error) error {
	if err == nil {
		return nil
	}

	/*
		fmt.Println(reflect.TypeOf(err))
		v, isAPIError := err.(*runtime.APIError)
		if isAPIError {
			response := reflect.ValueOf(v.Response)
			fun := response.MethodByName("Body")
			retvals := fun.Call([]reflect.Value{})
			body, _ := retvals[0].Interface().(io.ReadCloser)
			var bodyBytes []byte
			if body != nil {
				bodyBytes, _ = ioutil.ReadAll(body)
			}
			var payload models.Riaaserror

			berr := json.Unmarshal(bodyBytes, &payload)
			if berr != nil {
				return errors.New(string(bodyBytes))
			}
			return RiaasError{
				Payload: &payload,
			}
		}
	*/

	// check if its ours
	kind := reflect.TypeOf(err).Kind()
	if kind != reflect.Ptr {
		return err
	}

	// next follow pointer
	errstruct := reflect.TypeOf(err).Elem()
	if errstruct.Kind() != reflect.Struct {
		return err
	}

	n := errstruct.NumField()
	found := false
	for i := 0; i < n; i++ {
		if errstruct.Field(i).Name == "Payload" {
			found = true
			break
		}
	}

	if !found {
		return err
	}

	// check if a payload field exists
	payloadValue := reflect.ValueOf(err).Elem().FieldByName("Payload")
	if payloadValue.Interface() == nil {
		return err
	}

	payloadIntf := payloadValue.Elem().Interface()
	payload, parsed := payloadIntf.(models.Riaaserror)
	if !parsed {
		return err
	}

	if len(payload.Errors) == 0 {
		return nil
	}

	if len(payload.Errors) == 1 && payload.Errors[0].Code == "unexpected_return_value" {
		statuscode := reflect.ValueOf(err).Elem().FieldByName("_statusCode")
		payload.Errors[0].Target.Name = strconv.Itoa(int(statuscode.Int()))
		payload.Errors[0].Target.Type = "http_code"
	}
	var reterr = RiaasError{
		Payload: &payload,
	}

	return reterr
}
