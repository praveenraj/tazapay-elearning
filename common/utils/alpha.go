package utils

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"

	"tazapay.com/elearning/common/constants"
	"tazapay.com/elearning/common/responses"
)

// MakeResponseData build Data object for the common response
func MakeResponseData(data interface{}) *responses.Data {
	return &responses.Data{
		Value: data,
	}
}

// MakeResponseError build custom Err object with Env error type & code.
// @param envErrType: env error type.
// @param code: error code w.r.t error type.
// args...
// @type string: remarks for the Err obj.
func MakeResponseError(envErrType string, code int, args ...interface{}) *responses.Err {
	var remarks string
	for _, arg := range args {
		// more than one arg, use switch type assertions
		if reflect.TypeOf(arg).Kind() == reflect.String {
			remarks = cast.ToString(arg)
		}
	}
	return &responses.Err{
		Code:    code,
		Type:    envErrType,
		Remarks: remarks,
	}
}

// JSONMarshalAndUnmarshal unmarshal the source to destination.
// Provide the destination object as a pointer reference.
func JSONMarshalAndUnmarshal(src, dest interface{}) error {
	var b []byte
	var err error

	// get the source as bytes
	switch v := src.(type) {
	case string:
		b = []byte(v)

	case []byte:
		b = v

	default: // interface{}
		b, err = json.Marshal(src)
		// check for err if any
		if err != nil {
			return err
		}
	}

	// convert bytes into destination type
	err = json.Unmarshal(b, dest)
	if err != nil {
		return err
	}
	return nil
}

// IsValidStruct validates a struct based on beego tags.
// @param data: struct to be validated
func IsValidStruct(data interface{}) (isValid bool, tag string) {
	isValid = constants.BoolTrue

	valid := validation.Validation{}
	validated, err := valid.Valid(data)
	if err != nil {
		log.Error().Msgf("error validating struct: %v", err)
		isValid = constants.BoolFalse
		return isValid, tag
	}

	if !validated {
		isValid = constants.BoolFalse
		// get the reflect value of the struct to find the json tag
		value := reflect.ValueOf(data)
		// work around for a *Ptr struct
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		// iterate through errors and find out the field/tag
		for _, err := range valid.Errors {
			if err.Key != constants.Empty {
				field := err.Key[:strings.Index(err.Key, constants.Dot)] // field.Required.
				log.Error().Msgf("invalid field: %v | message: %v", field, err.Message)
				if structField, ok := value.Type().FieldByName(field); ok {
					tag = structField.Tag.Get("json")
				}
				break
			}
		}
	}
	return isValid, tag
}
