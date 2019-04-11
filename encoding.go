package elemental

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// Decode decodes the given data using an appropriate decoder chosen
// from the given contentType.
// It supports "application/gob". Anything else uses JSON.
func Decode(contentType string, data []byte, dest interface{}) error {

	switch contentType {

	case "application/gob":
		if err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(dest); err != nil {
			return fmt.Errorf("unable to decode gob: %s", err.Error())
		}

	default:
		if err := json.Unmarshal(data, dest); err != nil {
			return fmt.Errorf("unable to decode json: %s", err.Error())
		}
	}

	return nil
}

// Encode encodes the given object using an appropriate encoder chosen
// from the given acceptType.
// It supports "application/gob". Anything else uses JSON.
func Encode(acceptType string, obj interface{}) ([]byte, error) {

	switch acceptType {
	case "application/gob":
		buf := bytes.NewBuffer(nil)
		if err := gob.NewEncoder(buf).Encode(obj); err != nil {
			return nil, fmt.Errorf("unable to encode gob: %s", err.Error())
		}
		return buf.Bytes(), nil

	default:
		data, err := json.Marshal(obj)
		if err != nil {
			return nil, fmt.Errorf("unable to encode json: %s", err.Error())
		}
		return data, nil
	}
}
