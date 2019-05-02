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
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParser_Spaces(t *testing.T) {

	Convey("Given the operator '==' is not separated by spaces but quoted", t, func() {

		parser := NewFilterParser(`"tag"=="@sys:image=nginx"`)
		expectedFilter := NewFilterComposer().WithKey("tag").Equals("@sys:image=nginx").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator '==' is not separated by spaces", t, func() {

		parser := NewFilterParser(`a==b`)
		expectedFilter := NewFilterComposer().WithKey("a").Equals("b").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator '<=' is not separated by spaces", t, func() {

		parser := NewFilterParser(`a<=b`)
		expectedFilter := NewFilterComposer().WithKey("a").LesserOrEqualThan("b").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator '<' is not separated by spaces", t, func() {

		parser := NewFilterParser(`a<b`)
		expectedFilter := NewFilterComposer().WithKey("a").LesserThan("b").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator is separated by spaces", t, func() {

		parser := NewFilterParser(`a == b`)
		expectedFilter := NewFilterComposer().WithKey("a").Equals("b").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator complex value", t, func() {

		parser := NewFilterParser(`value == "age>=3"`)
		expectedFilter := NewFilterComposer().WithKey("value").Equals("age>=3").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator complex key", t, func() {

		parser := NewFilterParser(`"tag==toto"=="3==5"`)
		expectedFilter := NewFilterComposer().WithKey("tag==toto").Equals("3==5").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the weird case '"=="=="=="'`, t, func() {

		parser := NewFilterParser(`"=="=="=="`)
		expectedFilter := NewFilterComposer().WithKey("==").Equals("==").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the weird case '"=="    ==    "=="'`, t, func() {

		parser := NewFilterParser(`"=="    ==    "=="`)
		expectedFilter := NewFilterComposer().WithKey("==").Equals("==").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})
}

func TestParser_Keys(t *testing.T) {

	Convey("Given the quoted expression", t, func() {

		parser := NewFilterParser("\"key\" == value")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("value").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the unquoted expression", t, func() {

		parser := NewFilterParser("key == value")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("value").Done()

		Convey("When I run Parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the expression contains digits", t, func() {

		parser := NewFilterParser("key1234 == value")
		expectedFilter := NewFilterComposer().WithKey("key1234").Equals("value").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the expression is a tag like '@sys:usr:key'", t, func() {

		parser := NewFilterParser("@sys:usr:key == value")
		expectedFilter := NewFilterComposer().WithKey("@sys:usr:key").Equals("value").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the expression is a tag like '$key'", t, func() {

		parser := NewFilterParser("$key == value")
		// expectedFilter := NewFilterComposer().WithKey("$key").Equals("value").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should have an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "could not start a parameter with $")
				So(filter, ShouldEqual, nil)
			})
		})
	})

}
func TestParser_Keys_Errors(t *testing.T) {

	Convey(`Given the expression: "key`, t, func() {

		parser := NewFilterParser(`"key == chris`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `missing quote after key == chris`)
			})
		})
	})
}

func TestParser_Operators(t *testing.T) {

	Convey("Given the operator: '=='", t, func() {

		parser := NewFilterParser("key == value")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("value").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator '!='", t, func() {

		parser := NewFilterParser("key != value")
		expectedFilter := NewFilterComposer().WithKey("key").NotEquals("value").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'contains'", t, func() {

		parser := NewFilterParser(`key contains ["value1", "value2"]`)
		expectedFilter := NewFilterComposer().WithKey("key").Contains("value1", "value2").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'not contains'", t, func() {

		parser := NewFilterParser(`key not contains ["value1", "value2"]`)
		expectedFilter := NewFilterComposer().WithKey("key").NotContains("value1", "value2").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'in'", t, func() {

		parser := NewFilterParser(`key in ["value1", "value2"]`)
		expectedFilter := NewFilterComposer().WithKey("key").In("value1", "value2").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'not in'", t, func() {

		parser := NewFilterParser(`key not in ["value1", "value2"]`)
		expectedFilter := NewFilterComposer().WithKey("key").NotIn("value1", "value2").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'matches'", t, func() {

		parser := NewFilterParser(`key matches ["value1", "value2"]`)
		expectedFilter := NewFilterComposer().WithKey("key").Matches("value1", "value2").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'exists'", t, func() {

		parser := NewFilterParser(`key exists`)
		expectedFilter := NewFilterComposer().WithKey("key").Exists().Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'not exists'", t, func() {

		parser := NewFilterParser(`key not exists`)
		expectedFilter := NewFilterComposer().WithKey("key").NotExists().Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})
}

func TestParser_Operators_Errors(t *testing.T) {

	Convey(`Given the wrong operator '"'`, t, func() {

		parser := NewFilterParser(`key" == chris`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `invalid operator. found " instead of (==, !=, <, <=, >, >=, contains, in, matches, exists)`)
			})
		})
	})

	Convey(`Given the wrong operator 'and'`, t, func() {

		parser := NewFilterParser(`key and chris`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `invalid operator. found and instead of (==, !=, <, <=, >, >=, contains, in, matches, exists)`)
			})
		})
	})

	Convey(`Given the wrong operator: name == 0 and toto contains "1" an contains "@hello=2"`, t, func() {

		parser := NewFilterParser(`name == 0 and toto contains "1" an contains "@hello=2"`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, `invalid keyword after toto contains ["1"]. found an`)
			})
		})
	})

}

func TestParser_Values_StringType(t *testing.T) {

	Convey("Given the string value: '\"hello world\"'", t, func() {

		parser := NewFilterParser("key == \"hello world\"")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("hello world").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the string value with single quote: 'hello world'", t, func() {

		parser := NewFilterParser("key == 'hello world'")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("hello world").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the string value: '\"hello\"word\"'", t, func() {

		parser := NewFilterParser("key == \"hello\\\"world\"")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("hello\"world").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the string value: 'hello\\'word'", t, func() {

		parser := NewFilterParser("key == 'hello\\'world'")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("hello'world").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the string value: 'hello\\ word'", t, func() {

		parser := NewFilterParser("key == hello\\ world")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("hello world").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the string value: 'nowadays'", t, func() {

		parser := NewFilterParser("key == nowadays")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("nowadays").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the string value: 'datetime'", t, func() {

		parser := NewFilterParser("key == datetime")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("datetime").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

}

func TestParser_Values_Errors(t *testing.T) {

	Convey("Given the string value: 'hello word'", t, func() {

		parser := NewFilterParser("key == hello world")

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "missing parenthese to protect value: hello world")
			})
		})
	})

	Convey(`Given the string value: 'key == "hello'`, t, func() {

		parser := NewFilterParser(`key == "hello`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `missing quote after hello`)
			})
		})
	})

	Convey(`Given the single quoted string value: key == 'hello`, t, func() {

		parser := NewFilterParser(`key == 'hello`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `missing quote after hello`)
			})
		})
	})

	Convey(`Given the string value: key == hello"`, t, func() {

		parser := NewFilterParser(`key == hello"`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `missing quote before the value: hello`)
			})
		})
	})
	Convey(`Given the string value: key == hello'`, t, func() {

		parser := NewFilterParser(`key == hello'`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `missing quote before the value: hello`)
			})
		})
	})

	Convey(`Given the string value: key == 'hello"`, t, func() {

		parser := NewFilterParser(`key == 'hello"`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `missing quote after hello"`)
			})
		})
	})

	Convey(`Given the string value: key == "hello'`, t, func() {

		parser := NewFilterParser(`key == "hello'`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `missing quote after hello'`)
			})
		})
	})

	Convey(`Given the wrong value and: key == and"`, t, func() {

		parser := NewFilterParser(`key == and"`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `invalid value. found and`)
			})
		})
	})

	Convey(`Given the wrong value and: key exists value"`, t, func() {

		parser := NewFilterParser(`key exists value"`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `invalid keyword after key exists. found value`)
			})
		})
	})

}

func TestParser_Values_BoolType(t *testing.T) {
	Convey("Given the boolean value: 'true'", t, func() {

		parser := NewFilterParser("key == true")
		expectedFilter := NewFilterComposer().WithKey("key").Equals(true).Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the boolean value: 'false'", t, func() {

		parser := NewFilterParser("key == false")
		expectedFilter := NewFilterComposer().WithKey("key").Equals(false).Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the boolean value: '\"true\"'", t, func() {

		parser := NewFilterParser("key == \"true\"")
		expectedFilter := NewFilterComposer().WithKey("key").Equals("true").Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

}

func TestParser_Values_DateType(t *testing.T) {
	Convey("Given the date value: '2018-04-26'", t, func() {
		parser := NewFilterParser(`key == date("2018-04-26")`)
		expectedValue := time.Date(2018, time.April, 26, 0, 0, 0, 0, time.UTC)
		expectedFilter := NewFilterComposer().WithKey("key").Equals(expectedValue).Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the date value: '2018-04-26 23:50'", t, func() {
		parser := NewFilterParser(`key == date("2018-04-26 23:50")`)
		expectedValue := time.Date(2018, time.April, 26, 23, 50, 0, 0, time.UTC)
		expectedFilter := NewFilterComposer().WithKey("key").Equals(expectedValue).Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the date value: '2018-04-26T23:50:30.0Z'", t, func() {
		parser := NewFilterParser(`key > date("2018-04-26T23:50:30.0Z")`)
		expectedValue := time.Date(2018, time.April, 26, 23, 50, 30, 0, time.UTC)
		expectedFilter := NewFilterComposer().WithKey("key").GreaterThan(expectedValue).Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})
}

func TestParser_Values_DateType_Errors(t *testing.T) {

	Convey("Given the invalid date: 'invalid-date'", t, func() {
		parser := NewFilterParser(`key == date("invalid-date")`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `unable to parse date format invalid-date`)
			})
		})
	})

	Convey("Given the invalid date: '2012-24-2'", t, func() {
		parser := NewFilterParser(`key == date("2012-24-2")`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `unable to parse date format 2012-24-2`)
			})
		})
	})

	Convey("Given the invalid date: '2012-24-2' (missing quote)", t, func() {
		parser := NewFilterParser(`key == date(2012-24-2")`)

		Convey("When I run Parse", func() {

			_, err := parser.Parse()

			Convey("Then there should be an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, `unable to parse date format 2012-24-2`)
			})
		})
	})
}

func TestParser_Values_DurationType(t *testing.T) {
	Convey("Given the duration value: 'now()'", t, func() {
		parser := NewFilterParser(`key == now()`)
		expectedValue := time.Duration(0)
		expectedFilter := NewFilterComposer().WithKey("key").Equals(expectedValue).Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the duration value: '-1h'", t, func() {
		parser := NewFilterParser(`key == now("-1h")`)
		expectedValue := -1 * time.Hour
		expectedFilter := NewFilterComposer().WithKey("key").Equals(expectedValue).Done()

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})
}

func Test_Parser(t *testing.T) {

	Convey("Given the filter: namespace == chris and test == true", t, func() {

		parser := NewFilterParser("namespace == chris and test == true")

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()

			expectedFilter := NewFilterComposer().And(
				NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
				NewFilterComposer().WithKey("test").Equals(true).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the filter: namespace == chris and test == true or value > 30", t, func() {

		parser := NewFilterParser("namespace == chris and test == true or value > 30")

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().Or(
				NewFilterComposer().And(
					NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
					NewFilterComposer().WithKey("test").Equals(true).Done(),
				).Done(),
				NewFilterComposer().WithKey("value").GreaterThan(30).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the filter: namespace == chris or test == true and value > 30", t, func() {

		parser := NewFilterParser("namespace == chris or test == true and value > 30")

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().And(
				NewFilterComposer().Or(
					NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
					NewFilterComposer().WithKey("test").Equals(true).Done(),
				).Done(),
				NewFilterComposer().WithKey("value").GreaterThan(30).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: "namespace"=="chris" and "test"== true`, t, func() {

		parser := NewFilterParser(`"namespace"=="chris" and "test"== true`)

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().And(
				NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
				NewFilterComposer().WithKey("test").Equals(true).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: "namespace" == "chris" and "test" == true and date > date("2016-03-12")`, t, func() {

		parser := NewFilterParser(`"namespace" == "chris" and "test" == true and date > date("2016-03-12")`)

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().And(
				NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
				NewFilterComposer().WithKey("test").Equals(true).Done(),
				NewFilterComposer().WithKey("date").GreaterThan(time.Date(2016, time.March, 12, 0, 0, 0, 0, time.UTC)).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: "age" <= 32 or "age" > 50`, t, func() {

		parser := NewFilterParser(`"age" <= 32 or "age" > 50`)

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().Or(
				NewFilterComposer().WithKey("age").LesserOrEqualThan(32).Done(),
				NewFilterComposer().WithKey("age").GreaterThan(50).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: "age" < 32 or "age" >= 50`, t, func() {

		parser := NewFilterParser(`"age" < 32 or "age" >= 50`)

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().Or(
				NewFilterComposer().WithKey("age").LesserThan(32).Done(),
				NewFilterComposer().WithKey("age").GreaterOrEqualThan(50).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: ("file" matches "*.txt" and "file" contains "search")`, t, func() {

		parser := NewFilterParser(`("file" matches "*.txt" and "file" contains "search")`)

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().And(
				NewFilterComposer().WithKey("file").Matches("*.txt").Done(),
				NewFilterComposer().WithKey("file").Contains("search").Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: "namespace" == "/chris"`, t, func() {

		parser := NewFilterParser(`"namespace" == "/chris"`)

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().WithKey("namespace").Equals("/chris").Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})

	})

	Convey(`Given the filter: "namespace" == "/chris" and test == true and ("name" == toto or "name" == tata)`, t, func() {

		parser := NewFilterParser(`"namespace" == "/chris" and test == true and ("name" == toto or "name" == tata)`)

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().And(
				NewFilterComposer().WithKey("namespace").Equals("/chris").Done(),
				NewFilterComposer().WithKey("test").Equals(true).Done(),
				NewFilterComposer().Or(
					NewFilterComposer().WithKey("name").Equals("toto").Done(),
					NewFilterComposer().WithKey("name").Equals("tata").Done(),
				).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: (name == toto or name == tata) and age exists and age == 31`, t, func() {

		parser := NewFilterParser("(name == toto or name == tata) and age exists and age == 31")

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().And(
				NewFilterComposer().Or(
					NewFilterComposer().WithKey("name").Equals("toto").Done(),
					NewFilterComposer().WithKey("name").Equals("tata").Done(),
				).Done(),
				NewFilterComposer().WithKey("age").Exists().Done(),
				NewFilterComposer().WithKey("age").Equals(31).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: (name == toto and name == tata) or (age == 31 and age == 32)`, t, func() {

		parser := NewFilterParser("(name == toto and name == tata) or (age == 31 and age == 32)")

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().Or(
				NewFilterComposer().And(
					NewFilterComposer().WithKey("name").Equals("toto").Done(),
					NewFilterComposer().WithKey("name").Equals("tata").Done(),
				).Done(),
				NewFilterComposer().And(
					NewFilterComposer().WithKey("age").Equals(31).Done(),
					NewFilterComposer().WithKey("age").Equals(32).Done(),
				).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey(`Given the filter: name == toto and value == 38.9000 and ((name == toto and name == tata) or (age == 31 and age == 32) or protected in [true, false])`, t, func() {

		parser := NewFilterParser("name == toto and value == 38.9000 and ((name == toto and name == tata) or (age == 31 and age == 32) or protected in [true, false])")

		Convey("When I run Parse", func() {

			filter, err := parser.Parse()
			expectedFilter := NewFilterComposer().And(
				NewFilterComposer().WithKey("name").Equals("toto").Done(),
				NewFilterComposer().WithKey("value").Equals(38.9).Done(),
				NewFilterComposer().Or(
					NewFilterComposer().And(
						NewFilterComposer().WithKey("name").Equals("toto").Done(),
						NewFilterComposer().WithKey("name").Equals("tata").Done(),
					).Done(),
					NewFilterComposer().And(
						NewFilterComposer().WithKey("age").Equals(31).Done(),
						NewFilterComposer().WithKey("age").Equals(32).Done(),
					).Done(),
					NewFilterComposer().WithKey("protected").In(true, false).Done(),
				).Done(),
			).Done()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

}

func TestParser_AdvancedFilter(t *testing.T) {

	advancedFilter := `"namespace" == "coucou" and "number" == 32.900000 and (("name" == "toto" and "value" == 1) and ("color" contains ["red", "green", "blue", 43] and "something" in ["stuff"] or (("size" matches [".*"]) or ("size" == "medium" and "fat" == false) or ("size" in [true, false]))))`

	Convey(`Given I have an advanced complex filter and a parser`, t, func() {

		parser := NewFilterParser(advancedFilter)
		expectedFilter := NewFilterComposer().And(
			NewFilterComposer().WithKey("namespace").Equals("coucou").Done(),
			NewFilterComposer().WithKey("number").Equals(32.9).Done(),
			NewFilterComposer().And(

				NewFilterComposer().And(
					NewFilterComposer().WithKey("name").Equals("toto").Done(),
					NewFilterComposer().WithKey("value").Equals(1).Done(),
				).Done(),

				NewFilterComposer().Or(
					NewFilterComposer().And(
						NewFilterComposer().WithKey("color").Contains("red", "green", "blue", 43).Done(),
						NewFilterComposer().WithKey("something").In("stuff").Done(),
					).Done(),

					NewFilterComposer().Or(
						NewFilterComposer().WithKey("size").Matches(".*").Done(),
						NewFilterComposer().And(
							NewFilterComposer().WithKey("size").Equals("medium").Done(),
							NewFilterComposer().WithKey("fat").Equals(false).Done(),
						).Done(),
						NewFilterComposer().
							WithKey("size").In(true, false).
							Done(),
					).Done(),
				).Done(),
			).Done(),
		).Done()

		Convey("When I run parse", func() {

			filter, err := parser.Parse()

			Convey("Then there should be no error and the filter should as expected", func() {
				So(err, ShouldEqual, nil)
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})

		Convey("When I run multiple parse", func() {

			filter, err := parser.Parse()

			So(err, ShouldEqual, nil)

			p := NewFilterParser(filter.String())
			f, err := p.Parse()

			Convey("Then there should be no error and the filter should be equal to the previous filter", func() {
				So(err, ShouldEqual, nil)
				So(f, ShouldNotEqual, nil)
				So(f.String(), ShouldEqual, filter.String())
			})
		})

	})

}

func Test_isLetter(t *testing.T) {
	Convey("Given I have a FilterParser", t, func() {
		So(isLetter('<'), ShouldEqual, true)
		So(isLetter('b'), ShouldEqual, true)
		So(isLetter(4), ShouldEqual, false)
		So(isLetter('\\'), ShouldEqual, true)
	})
}
