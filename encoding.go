// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elemental

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/ugorji/go/codec"
)

var (
	externalSupportedContentType = map[string]struct{}{}
	externalSupportedAcceptType  = map[string]struct{}{}
)

// RegisterSupportedContentType registers a new media type
// that elemental should support for Content-Type.
// Note that this needs external intervention to handle encoding.
func RegisterSupportedContentType(mimetype string) {
	externalSupportedContentType[mimetype] = struct{}{}
}

// RegisterSupportedAcceptType registers a new media type
// that elemental should support for Accept.
// Note that this needs external intervention to handle decoding.
func RegisterSupportedAcceptType(mimetype string) {
	externalSupportedAcceptType[mimetype] = struct{}{}
}

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
	// If you need to understand all of this, go there http://ugorji.net/blog/go-codec-primer
	// But you should not need to touch that.
	jsonHandle.Canonical = true
	jsonHandle.MapType = reflect.TypeOf(map[string]interface{}(nil))

	msgpackHandle.Canonical = true
	msgpackHandle.WriteExt = true
	msgpackHandle.MapType = reflect.TypeOf(map[string]interface{}(nil))
	msgpackHandle.TypeInfos = codec.NewTypeInfos([]string{"msgpack"})
}

// Decode decodes the given data using an appropriate decoder chosen
// from the given encoding.
func Decode(encoding EncodingType, data []byte, dest interface{}) error {

	var pool *sync.Pool

	switch encoding {
	case EncodingTypeMSGPACK:
		pool = &msgpackDecodersPool
	default:
		pool = &jsonDecodersPool
		encoding = EncodingTypeJSON
	}

	dec := pool.Get().(*codec.Decoder)
	defer pool.Put(dec)

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

	var pool *sync.Pool

	switch encoding {
	case EncodingTypeMSGPACK:
		pool = &msgpackEncodersPool
	default:
		pool = &jsonEncodersPool
		encoding = EncodingTypeJSON
	}

	enc := pool.Get().(*codec.Encoder)
	defer pool.Put(enc)

	buf := bytes.NewBuffer(nil)
	enc.Reset(buf)

	if err := enc.Encode(obj); err != nil {
		return nil, fmt.Errorf("unable to encode %s: %s", encoding, err.Error())
	}

	return buf.Bytes(), nil
}

// MakeStreamDecoder returns a function that can be used to decode a stream from the
// given reader using the given encoding.
//
// This function returns the decoder function that can be called until it returns an
// io.EOF error, indicating the stream is over, and a dispose function that will
// put back the decoder in the memory pool.
// The dispose function will be called automatically when the decoding is over,
// but not on a single decoding error.
// In any case, the dispose function should be always called, in a defer for example.
func MakeStreamDecoder(encoding EncodingType, reader io.Reader) (func(dest interface{}) error, func()) {

	var pool *sync.Pool

	switch encoding {
	case EncodingTypeMSGPACK:
		pool = &msgpackDecodersPool
	default:
		pool = &jsonDecodersPool
	}

	dec := pool.Get().(*codec.Decoder)
	dec.Reset(reader)

	clean := func() {
		if pool != nil {
			pool.Put(dec)
			pool = nil
		}
	}

	return func(dest interface{}) error {

			if err := dec.Decode(dest); err != nil {

				if err == io.EOF {
					clean()
					return err
				}

				return fmt.Errorf("unable to decode %s: %s", encoding, err.Error())
			}

			return nil
		}, func() {
			clean()
		}
}

// MakeStreamEncoder returns a function that can be user en encode given data
// into the given io.Writer using the given encoding.
//
// It also returns a function must be called once the encoding procedure
// is complete, so the internal encoders can be put back into the shared
// memory pools.
func MakeStreamEncoder(encoding EncodingType, writer io.Writer) (func(obj interface{}) error, func()) {

	var pool *sync.Pool

	switch encoding {
	case EncodingTypeMSGPACK:
		pool = &msgpackEncodersPool
	default:
		pool = &jsonEncodersPool
	}

	enc := pool.Get().(*codec.Encoder)
	enc.Reset(writer)

	clean := func() {
		if pool != nil {
			pool.Put(enc)
			pool = nil
		}
	}

	return func(dest interface{}) error {

			if err := enc.Encode(dest); err != nil {
				return fmt.Errorf("unable to encode %s: %s", encoding, err.Error())
			}

			return nil
		}, func() {
			clean()
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

		switch ct {

		case "application/msgpack":
			read = EncodingTypeMSGPACK

		case "application/*", "*/*", "application/json":
			read = EncodingTypeJSON

		default:
			var supported bool
			for t := range externalSupportedContentType {
				if ct == t {
					supported = true
					break
				}
			}
			if !supported {
				return "", "", NewError("Unsupported Media Type", fmt.Sprintf("Cannot find any acceptable Content-Type media type in provided header: %s", v), "elemental", http.StatusUnsupportedMediaType)
			}

			read = EncodingType(ct)
		}
	}

	if v := header.Get("Accept"); v != "" {
		var agreed bool
	L:
		for _, item := range strings.Split(v, ",") {

			at, _, err := mime.ParseMediaType(item)
			if err != nil {
				return "", "", NewError("Bad Request", fmt.Sprintf("Invalid Accept header: %s", err), "elemental", http.StatusBadRequest)
			}

			switch at {

			case "application/msgpack":
				write = EncodingTypeMSGPACK
				agreed = true
				break L

			case "application/*", "*/*", "application/json":
				write = EncodingTypeJSON
				agreed = true
				break L

			default:
				for t := range externalSupportedAcceptType {
					if at == t {
						agreed = true
						write = EncodingType(at)
						break L
					}
				}
			}
		}

		if !agreed {
			return "", "", NewError("Unsupported Media Type", fmt.Sprintf("Cannot find any acceptable Accept media type in provided header: %s", v), "elemental", http.StatusUnsupportedMediaType)
		}
	}

	return read, write, nil
}
