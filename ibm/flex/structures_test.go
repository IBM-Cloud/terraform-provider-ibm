package flex

import (
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStringifyString(t *testing.T) {
	var foo interface{}
	foo = nil
	assert.Equal(t, "", Stringify(foo))

	foo = "a string"
	assert.Equal(t, foo, Stringify(foo))

	foo = ""
	assert.Equal(t, foo, Stringify(foo))
}

func TestStringifyDate(t *testing.T) {
	var foo interface{}
	foo = nil
	assert.Equal(t, "", Stringify(foo))

	d := "2025-06-03"
	foo, err := core.ParseDate(d)
	assert.Nil(t, err)
	assert.Equal(t, d, Stringify(foo))
}

func TestStringifyDateTime(t *testing.T) {
	var foo interface{}
	foo = nil
	assert.Equal(t, "", Stringify(foo))

	dt := "2025-06-03T11:59:59.999Z"
	foo, err := core.ParseDateTime(dt)
	assert.Nil(t, err)
	assert.Equal(t, dt, Stringify(foo))
}

func TestStringifyUUID(t *testing.T) {
	var foo interface{}
	foo = nil
	assert.Equal(t, "", Stringify(foo))

	u := uuid.New().String()
	foo = strfmt.UUID(u)
	assert.NotNil(t, foo)
	assert.Equal(t, u, Stringify(foo))
}
func TestStringifyBoolean(t *testing.T) {
	var foo interface{}
	foo = true
	assert.Equal(t, "true", Stringify(foo))

	foo = false
	assert.Equal(t, "false", Stringify(foo))
}

func TestStringifyNumber(t *testing.T) {
	var foo interface{}
	foo = 38
	assert.Equal(t, "38", Stringify(foo))

	foo = 38.12345
	assert.Equal(t, "38.12345", Stringify(foo))
}

func TestStringifyList(t *testing.T) {
	var foo interface{}
	foo = []string{"foo", "bar"}
	assert.Equal(t, `["foo","bar"]`, Stringify(foo))

	foo = []bool{true, false, true}
	assert.Equal(t, `[true,false,true]`, Stringify(foo))
}

func TestStringifyMap(t *testing.T) {
	var foo interface{} = map[string]interface{}{"foo": "bar"}
	assert.Equal(t, `{"foo":"bar"}`, Stringify(foo))
}
