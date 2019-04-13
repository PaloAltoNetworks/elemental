package elemental

import (
	"bytes"
	"fmt"
	"mime"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/ugorji/go/codec"
)

// An Encodable is the interface of objects
// that can hold encoding information.
type Encodable interface {
	GetEncoding() EncodingType
}

// A Encoder is an Encodable that can be encoded.
type Encoder interface {
	Encode(obj interface{}) (err error)
	Encodable
}

// A Decoder is an Encodable that can be decoded.
type Decoder interface {
	Decode(dst interface{}) error
	Encodable
}

// An EncodingType represents one type of data encoding
type EncodingType string

// Various values for EncodingType.
const (
	EncodingTypeJSON    EncodingType = "application/json"
	EncodingTypeMSGPACK EncodingType = "application/msgpack"
)

var (
	jsonHandle       = &codec.JsonHandle{}
	jsonEncodersPool = sync.Pool{
		New: func() interface{} {
			return codec.NewEncoder(nil, jsonHandle)
		},
	}
	jsonDecodersPool = sync.Pool{
		New: func() interface{} {
			return codec.NewDecoder(nil, jsonHandle)
		},
	}

	msgpackHandle       = &codec.MsgpackHandle{}
	msgpackEncodersPool = sync.Pool{
		New: func() interface{} {
			return codec.NewEncoder(nil, msgpackHandle)
		},
	}
	msgpackDecodersPool = sync.Pool{
		New: func() interface{} {
			return codec.NewDecoder(nil, msgpackHandle)
		},
	}
)

func init() {
	time.Local = time.UTC

	// If you need to understand all of this, go there http://ugorji.net/blog/go-codec-primer
	// But you should not need to touch that.
	jsonHandle.Canonical = true
	jsonHandle.MapType = reflect.ValueOf(map[string]interface{}{}).Type()
	jsonHandle.MapValueReset = true

	msgpackHandle.WriteExt = true
	msgpackHandle.Canonical = true
	msgpackHandle.TypeInfos = codec.NewTypeInfos([]string{"msgpack"})
	msgpackHandle.MapType = reflect.ValueOf(map[string]interface{}{}).Type()
	msgpackHandle.MapValueReset = true
}

// Decode decodes the given data using an appropriate decoder chosen
// from the given contentType.
func Decode(encoding EncodingType, data []byte, dest interface{}) error {

	var dec *codec.Decoder

	switch encoding {
	case EncodingTypeMSGPACK:
		dec = msgpackDecodersPool.Get().(*codec.Decoder)
		defer msgpackDecodersPool.Put(dec)
	default:
		dec = jsonDecodersPool.Get().(*codec.Decoder)
		defer jsonDecodersPool.Put(dec)
		encoding = EncodingTypeJSON
	}

	dec.Reset(bytes.NewBuffer(data))

	if err := dec.Decode(dest); err != nil {
		return fmt.Errorf("unable to decode %s: %s", encoding, err.Error())
	}

	return nil
}

// Encode encodes the given object using an appropriate encoder chosen
// from the given acceptType.
func Encode(encoding EncodingType, obj interface{}) ([]byte, error) {

	if obj == nil {
		return nil, fmt.Errorf("encode received a nil object")
	}

	var enc *codec.Encoder

	switch encoding {
	case EncodingTypeMSGPACK:
		enc = msgpackEncodersPool.Get().(*codec.Encoder)
		defer msgpackEncodersPool.Put(enc)
	default:
		enc = jsonEncodersPool.Get().(*codec.Encoder)
		defer jsonEncodersPool.Put(enc)
		encoding = EncodingTypeJSON
	}

	buf := bytes.NewBuffer(nil)
	enc.Reset(buf)

	if err := enc.Encode(obj); err != nil {
		return nil, fmt.Errorf("unable to encode %s: %s", encoding, err.Error())
	}

	return buf.Bytes(), nil
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
