package utils_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"tazapay.com/elearning/common/responses"
	"tazapay.com/elearning/common/utils"
)

func TestMakeResponseData(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Data
	}{
		{"default", args{map[string]interface{}{"company": "tryllo"}}, &responses.Data{Value: map[string]interface{}{"company": "tryllo"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.MakeResponseData(tt.args.data)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMakeResponseError(t *testing.T) {
	type args struct {
		envErrType string
		code       int
		args       []interface{}
	}
	tests := []struct {
		name string
		args args
		want *responses.Err
	}{
		{"default", args{"errors.auth", 1100, []interface{}{}}, &responses.Err{Code: 1100, Message: "", Type: "errors.auth", Remarks: ""}},
		{"remarks", args{"errors.auth", 1100, []interface{}{"test"}}, &responses.Err{Code: 1100, Message: "", Type: "errors.auth", Remarks: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.MakeResponseError(tt.args.envErrType, tt.args.code, tt.args.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestJSONMarshalAndUnmarshal(t *testing.T) {
	type source struct {
		Key   string      `json:"key,omitempty"`
		Value interface{} `json:"value,omitempty"`
	}
	type destination struct {
		Key string `json:"key,omitempty"`
	}
	type args struct {
		src  interface{}
		dest interface{}
	}

	srcByte, _ := json.Marshal(source{Key: "one", Value: 1})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success", args{src: source{Key: "one", Value: 1}, dest: &destination{}}, false},
		{"success[]byte", args{src: srcByte, dest: &destination{}}, false},
		{"marshalError", args{src: make(chan int), dest: &destination{}}, true},
		{"unmarshalError", args{src: `{"key":what?}`, dest: &destination{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := utils.JSONMarshalAndUnmarshal(tt.args.src, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("JSONMarshalAndUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidStruct(t *testing.T) {
	type d struct {
		Name   string `valid:"Required;Match(/^.*?$/)" json:"name"`
		Age    int    `valid:"Required;Range(1, 140)" json:"age"`
		Mobile string `valid:"Required;Match(/^[0-9]{10}$/)" json:"mobile"`
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantIsValid bool
		wantField   string
	}{
		{"valid", args{d{Name: "test", Age: 40, Mobile: "1234567895"}}, true, ""},
		{"invalidFieldName", args{d{Name: "", Age: 40, Mobile: "1234567895"}}, false, "name"},
		{"invalidFieldAge", args{d{Name: "abc", Age: 240, Mobile: "1234567895"}}, false, "age"},
		{"invalidFieldMobile", args{d{Name: "abc", Age: 40, Mobile: "12345895"}}, false, "mobile"},
		{"invalidFieldNamePtrStruct", args{&d{Name: "", Age: 40, Mobile: "1234567895"}}, false, "name"},
		{"validateErr", args{100}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsValid, gotField := utils.IsValidStruct(tt.args.data)
			assert.Equal(t, tt.wantIsValid, gotIsValid)
			assert.Equal(t, tt.wantField, gotField)
		})
	}
}
