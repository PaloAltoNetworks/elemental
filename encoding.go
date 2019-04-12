package elemental

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"time"

	"github.com/vmihailenco/msgpack"
)

// An EncodingType represents one type of data encoding
type EncodingType string

// Various values for EncodingType.
const (
	EncodingTypeJSON    EncodingType = "application/json"
	EncodingTypeMSGPACK EncodingType = "application/msgpack"
)

func init() {
	time.Local = time.UTC
}

// Decode decodes the given data using an appropriate decoder chosen
// from the given contentType.
func Decode(encoding EncodingType, data []byte, dest interface{}) error {

	switch encoding {

	case EncodingTypeMSGPACK:
		// fmt.Print("M")
		dec := msgpack.NewDecoder(bytes.NewBuffer(data))
		dec.UseJSONTag(true)
		if err := dec.Decode(dest); err != nil {
			return fmt.Errorf("unable to decode msgpack: %s", err.Error())
		}

	default:
		// fmt.Print("J")
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

	case EncodingTypeMSGPACK:
		// fmt.Print("M")
		buf := bytes.NewBuffer(nil)
		enc := msgpack.NewEncoder(buf)
		enc.UseJSONTag(true)
		enc.SortMapKeys(true)
		if err := enc.Encode(obj); err != nil {
			return nil, fmt.Errorf("unable to encode msgpack: %s", err.Error())
		}
		return buf.Bytes(), nil

	default:
		// fmt.Print("J")
		data, err := json.Marshal(obj)
		if err != nil {
			return nil, fmt.Errorf("unable to encode json: %s", err.Error())
		}
		return data, nil
	}
}

// Convert converts from one EncodingType to another
func Convert(from EncodingType, to EncodingType, data []byte) ([]byte, error) {

	if from == to {
		return data, nil
	}

	m := map[string]interface{}{}
	if err := Decode(from, data, &m); err != nil {
		return nil, err
	}

	return Encode(to, m)
}

// EncodingFromHeaders returns the read (Content-Type) and write (Accept) encoding
// from the given http.Header.
func EncodingFromHeaders(header http.Header) (read EncodingType, write EncodingType, err error) {

	read = EncodingTypeJSON
	write = EncodingTypeJSON

	if header == nil {
		return read, write, nil
	}

	if v := header.Get("Content-Type"); v != "" {
		ct, _, err := mime.ParseMediaType(v)
		if err != nil {
			return "", "", NewError("Bad Request", fmt.Sprintf("Invalid Content-Type header: %s", err), "elemental", http.StatusBadRequest)
		}
		read = EncodingType(ct)
	}

	if v := header.Get("Accept"); v != "" {
		at, _, err := mime.ParseMediaType(v)
		if err != nil {
			return "", "", NewError("Bad Request", fmt.Sprintf("Invalid Accept header: %s", err), "elemental", http.StatusBadRequest)
		}
		write = EncodingType(at)
	}

	// TODO: handle unsupported types.

	return read, write, nil
}
