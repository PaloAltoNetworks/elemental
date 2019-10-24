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
	"encoding/base64"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAttribute_Interface(t *testing.T) {

	Convey("Given I create a new *List", t, func() {

		l := NewList()

		Convey("Then the list should implement the AttributeSpecifiable interface", func() {
			So(l, ShouldImplement, (*AttributeSpecifiable)(nil))
		})
	})
}

func TestAttribute_SpecificationForAttribute(t *testing.T) {

	Convey("Given I create a new List", t, func() {

		l := NewList()

		Convey("When I get the Attribute specification for the name", func() {

			spec := l.SpecificationForAttribute("Name")

			Convey("then it should be correct", func() {
				So(spec.AllowedChars, ShouldEqual, "")
				So(len(spec.AllowedChoices), ShouldEqual, 0)
				So(spec.Autogenerated, ShouldBeFalse)
				So(spec.Availability, ShouldBeEmpty)
				So(spec.Channel, ShouldBeEmpty)
				So(spec.CreationOnly, ShouldBeFalse)
				So(spec.Deprecated, ShouldBeFalse)
				So(spec.Exposed, ShouldBeTrue)
				So(spec.Filterable, ShouldBeTrue)
				So(spec.ForeignKey, ShouldBeFalse)
				So(spec.Getter, ShouldBeTrue)
				So(spec.Identifier, ShouldBeFalse)
				So(spec.Index, ShouldBeFalse)
				So(spec.MaxLength, ShouldEqual, 0)
				So(spec.MaxValue, ShouldEqual, 0)
				So(spec.MinLength, ShouldEqual, 0)
				So(spec.MinValue, ShouldEqual, 0)
				So(spec.Orderable, ShouldBeTrue)
				So(spec.PrimaryKey, ShouldBeFalse)
				So(spec.ReadOnly, ShouldBeFalse)
				So(spec.Required, ShouldBeTrue)
				So(spec.Setter, ShouldBeTrue)
				So(spec.Stored, ShouldBeTrue)
				So(spec.SubType, ShouldBeEmpty)
				So(spec.Transient, ShouldBeFalse)
			})
		})
	})
}

func Test_ResetSecretAttributesValues(t *testing.T) {

	Convey("Given I have an identifiable with a secret property", t, func() {

		l := NewList()
		l.Secret = "it's a secret to everybody"

		Convey("When I call ResetSecretAttributesValues", func() {

			ResetSecretAttributesValues(l)

			Convey("Then the secret should have been erased", func() {
				So(l.Secret, ShouldEqual, "")
			})
		})
	})

	Convey("Given I have some identifiables with a secret property", t, func() {

		l1 := NewList()
		l1.Secret = "it's a secret to everybody"

		l2 := NewList()
		l2.Secret = "it's a secret to everybody"

		Convey("When I call ResetSecretAttributesValues", func() {

			ResetSecretAttributesValues(ListsList{l1, l2})

			Convey("Then the secret should have been erased", func() {
				So(l1.Secret, ShouldEqual, "")
				So(l2.Secret, ShouldEqual, "")
			})
		})
	})

	Convey("Given I have an sparse identifiable with a secret property", t, func() {

		val := "it's a secret to everybody"

		l := NewSparseList()
		l.Secret = &val

		Convey("When I call ResetSecretAttributesValues", func() {

			ResetSecretAttributesValues(l)

			Convey("Then the secret should have been erased", func() {
				So(l.Secret, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some sparse identifiables with a secret property", t, func() {

		l1 := NewSparseList()
		val1 := "it's a secret to everybody"
		l1.Secret = &val1

		l2 := NewSparseList()
		val2 := "it's a secret to everybody"
		l2.Secret = &val2

		Convey("When I call ResetSecretAttributesValues", func() {

			ResetSecretAttributesValues(SparseListsList{l1, l2})

			Convey("Then the secret should have been erased", func() {
				So(l1.Secret, ShouldBeNil)
				So(l2.Secret, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some random non pointer struct", t, func() {

		s := struct{}{}

		Convey("When I call ResetSecretAttributesValues", func() {

			ResetSecretAttributesValues(s)

			Convey("Then it should not panic", func() {
				So(func() { ResetSecretAttributesValues(s) }, ShouldNotPanic)
			})
		})
	})

	Convey("Given I have some random pointer struct", t, func() {

		s := &struct{}{}

		Convey("When I call ResetSecretAttributesValues", func() {

			ResetSecretAttributesValues(s)

			Convey("Then it should not panic", func() {
				So(func() { ResetSecretAttributesValues(s) }, ShouldNotPanic)
			})
		})
	})

	Convey("Given I have an non pointer identifiable with a secret property", t, func() {

		l := NewList()
		l.Secret = "it's a secret to everybody"

		Convey("When I call ResetSecretAttributesValues", func() {

			ResetSecretAttributesValues(*l)

			Convey("Then the secret should have been erased", func() {
				So(l.Secret, ShouldEqual, "it's a secret to everybody")
			})
		})
	})
}

func TestNewAESEncrypter(t *testing.T) {

	Convey("Given I call AESAttributeEncrypter with valid passphrase", t, func() {

		enc, err := NewAESAttributeEncrypter("0123456789ABCDEF")

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then enc should be correct", func() {
			So(enc.(*aesAttributeEncrypter).passphrase, ShouldResemble, []byte("0123456789ABCDEF"))
		})
	})

	Convey("Given I call AESAttributeEncrypter with passphrase that is too small", t, func() {

		enc, err := NewAESAttributeEncrypter("0123456789ABCDE")

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "invalid passphrase: size must be exactly 16 bytes")
		})

		Convey("Then enc should be nil", func() {
			So(enc, ShouldBeNil)
		})
	})

	Convey("Given I call AESAttributeEncrypter with passphrase that is too long", t, func() {

		enc, err := NewAESAttributeEncrypter("0123456789ABCDE WEEEE")

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "invalid passphrase: size must be exactly 16 bytes")
		})

		Convey("Then enc should be nil", func() {
			So(enc, ShouldBeNil)
		})
	})
}

func TestAESEncrypterEncryption(t *testing.T) {

	Convey("Given I have an AESAttributeEncrypter ", t, func() {

		value := "hello world"
		enc, _ := NewAESAttributeEncrypter("0123456789ABCDEF")

		Convey("When I encrypt some data", func() {

			encstring, err := enc.EncryptString(value)
			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			b64decodeddata, err1 := base64.StdEncoding.DecodeString(encstring)
			Convey("Then err1 should be nil", func() {
				So(err1, ShouldBeNil)
			})

			Convey("Then encstring should be encrypted", func() {
				So(encstring, ShouldNotEqual, value)
				So(b64decodeddata, ShouldNotEqual, value)
			})

			Convey("When I decrypt the data", func() {

				decstring, err := enc.DecryptString(encstring)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then decstring should be decrypted", func() {
					So(decstring, ShouldEqual, value)
				})
			})
		})

		Convey("When I encrypt empty string", func() {

			encstring, err := enc.EncryptString("")

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then encstring should be empty", func() {
				So(encstring, ShouldEqual, "")
			})
		})

		Convey("When I decrypt empty string", func() {

			decstring, err := enc.DecryptString("")

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then decstring should be empty", func() {
				So(decstring, ShouldEqual, "")
			})
		})

		Convey("When I decrypt non base64", func() {

			decstring, err := enc.DecryptString("1")

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "illegal base64 data at input byte 0")
			})

			Convey("Then decstring should be empty", func() {
				So(decstring, ShouldEqual, "")
			})
		})

		Convey("When I decrypt too small data", func() {

			decstring, err := enc.DecryptString("abcd")

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "data is too small")
			})

			Convey("Then decstring should be empty", func() {
				So(decstring, ShouldEqual, "")
			})
		})
	})
}
