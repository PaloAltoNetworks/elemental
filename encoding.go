package elemental

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"

	"github.com/vmihailenco/msgpack"
)

// An EncodingType represents one type of data encoding
type EncodingType string

// Various values for EncodingType.
const (
	EncodingTypeJSON    EncodingType = "application/json"
	EncodingTypeMSGPACK EncodingType = "application/msgpack"
	EncodingTypeGOB     EncodingType = "application/gob"
)

func init() {
	time.Local = time.UTC
}

// Decode decodes the given data using an appropriate decoder chosen
// from the given contentType.
func Decode(encoding EncodingType, data []byte, dest interface{}) error {

	switch encoding {

	case EncodingTypeGOB:
		dec := gob.NewDecoder(bytes.NewBuffer(data))
		if err := dec.Decode(dest); err != nil {
			return fmt.Errorf("unable to decode gob: %s", err.Error())
		}

	case EncodingTypeMSGPACK:
		dec := msgpack.NewDecoder(bytes.NewBuffer(data))
		dec.UseJSONTag(true)
		if err := dec.Decode(dest); err != nil {
			return fmt.Errorf("unable to decode msgpack: %s", err.Error())
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
func Encode(encoding EncodingType, obj interface{}) ([]byte, error) {

	switch encoding {

	case EncodingTypeGOB:
		buf := bytes.NewBuffer(nil)
		enc := gob.NewEncoder(buf)
		if err := enc.Encode(obj); err != nil {
			return nil, fmt.Errorf("unable to encode gob: %s", err.Error())
		}
		return buf.Bytes(), nil

	case EncodingTypeMSGPACK:
		buf := bytes.NewBuffer(nil)
		enc := msgpack.NewEncoder(buf)
		enc.UseJSONTag(true)
		if err := enc.Encode(obj); err != nil {
			return nil, fmt.Errorf("unable to encode msgpack: %s", err.Error())
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
