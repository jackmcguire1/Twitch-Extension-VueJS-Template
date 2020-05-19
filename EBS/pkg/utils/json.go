package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// Make a generic json encoded couch key
func ToJSON(s interface{}) string {
	buf, _ := json.Marshal(s)
	return string(buf)
}

// Convert to RawMessage
func ToRawMessage(s interface{}) json.RawMessage {
	buf, _ := json.Marshal(s)
	return buf
}

// ... and make it even easier to use
func ToJSONArray(s ...interface{}) string {
	return ToJSON(s)
}

// JSONRemap will apply the remap argument to each top level key string in doc
func JSONRemap(doc []byte, remap func(string) string) (buf []byte, err error) {

	src := map[string]*json.RawMessage{}
	err = json.Unmarshal(doc, &src)
	if err != nil {
		return
	}

	dst := map[string]*json.RawMessage{}
	for k, v := range src {
		nk := remap(k)
		if _, collision := dst[nk]; collision {
			err = fmt.Errorf("Key collision %s", nk)
			return
		}
		dst[nk] = v
	}

	return json.Marshal(&dst)
}

// JSONBytesEqual compares the JSON in two byte slices
func JSONBytesEqual(a, b []byte) (bool, error) {
	if bytes.Equal(a, b) {
		return true, nil
	}

	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// RawMessageMapsEqual compares two map[string]json.RawMessage
func RawMessageMapsEqual(
	a, b map[string]json.RawMessage,
) (
	equal bool,
	err error,
) {
	if len(a) == len(b) {
		for k, v := range a {
			val, ok := b[k]
			if !ok {
				break
			}

			equal, err = JSONBytesEqual(v, val)
			if err != nil {
				log.Print(err)
				return
			}

			if !equal {
				break
			}
		}
	}
	return
}
