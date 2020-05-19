package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

var expectHave = "Expect %q, have %q"

type ThriftGenerated struct {
	Id        string `json:"id"`
	CamelCase string `json:"camelCase"`
}

func TestJSONRemap(t *testing.T) {

	thriftStruct := ThriftGenerated{}

	id := "someid"
	camelCase := "mongoosed"

	doc := fmt.Sprintf(`{"_id":%q,"camel_case":%q}`, id, camelCase)

	buf, err := JSONRemap([]byte(doc), CamelCase)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(buf, &thriftStruct)
	if err != nil {
		t.Error(err)
	}

	if thriftStruct.Id != id {
		t.Errorf(expectHave, id, thriftStruct.Id)
	}

	if thriftStruct.CamelCase != camelCase {
		t.Errorf(expectHave, camelCase, thriftStruct.CamelCase)
	}
}
