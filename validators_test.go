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
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidator_ValidateRequiredInt(t *testing.T) {

	Convey("Given I call the method ValidateRequiredInt with a valid int", t, func() {

		validationError := ValidateRequiredInt("age", 15)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateRequiredInt with a nonvalid int", t, func() {

		validationError := ValidateRequiredInt("age", 0).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, "Attribute 'age' is required")
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidateRequiredExternal(t *testing.T) {

	Convey("Given I call the method ValidateRequiredExternal with a valid value", t, func() {

		validationError := ValidateRequiredExternal("age", 15)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateRequiredExternal with a valid array", t, func() {

		validationError := ValidateRequiredExternal("ages", []string{"coucou"})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateRequiredExternal with an empty array", t, func() {

		validationError := ValidateRequiredExternal("ages", []string{}).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, "Attribute 'ages' is required")
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "ages"})
		})
	})

	Convey("Given I call the method ValidateRequiredExternal with a nonvalid int", t, func() {

		validationError := ValidateRequiredExternal("age", nil).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, "Attribute 'age' is required")
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidateRequiredFloat(t *testing.T) {

	Convey("Given I call the method ValidateRequiredFloat with a valid float", t, func() {

		validationError := ValidateRequiredFloat("age", 15.3)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateRequiredFloat with a nonvalid float", t, func() {

		validationError := ValidateRequiredFloat("age", 0).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, "Attribute 'age' is required")
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidateRequiredString(t *testing.T) {

	Convey("Given I call the method ValidateRequiredString with a valid string", t, func() {

		validationError := ValidateRequiredString("name", "Alexandre")

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateRequiredString with a nonvalid string", t, func() {

		validationError := ValidateRequiredString("name", "").(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, "Attribute 'name' is required")
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})
}

func TestValidator_ValidateRequiredTime(t *testing.T) {

	Convey("Given I call the method ValidateRequiredTime with a valid time", t, func() {

		validationError := ValidateRequiredTime("date", time.Now())

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateRequiredTime with a nonvalid time", t, func() {

		var t time.Time
		validationError := ValidateRequiredTime("date", t).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, "Attribute 'date' is required")
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "date"})
		})
	})
}

func TestValidator_ValidateMaximumFloat(t *testing.T) {

	Convey("Given I call the method ValidateMaximumFloat with a valid float and none exclusive", t, func() {

		validationError := ValidateMaximumFloat("age", 12.4, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumFloat with a unvalid float and none exclusive", t, func() {

		validationError := ValidateMaximumFloat("age", 18.1, 18, false).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18.1' of attribute 'age' should be less than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})

	Convey("Given I call the method ValidateMaximumFloat with a valid float and none exclusive", t, func() {
		validationError := ValidateMaximumFloat("age", 18, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumFloat with a valid float and exclusive", t, func() {
		validationError := ValidateMaximumFloat("age", 12.4, 18, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumFloat with a unvalid float and exclusive", t, func() {

		validationError := ValidateMaximumFloat("age", 18.1, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18.1' of attribute 'age' should be less or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})

	Convey("Given I call the method ValidateMaximumFloat with a unvalid float and exclusive", t, func() {

		validationError := ValidateMaximumFloat("age", 18, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be less or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidateMinimumFloat(t *testing.T) {

	Convey("Given I call the method ValidateMinimumFloat with a valid float and none exclusive", t, func() {

		validationError := ValidateMinimumFloat("age", 12.4, 6.1, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumFloat with a unvalid float and none exclusive", t, func() {

		validationError := ValidateMinimumFloat("age", 18.1, 19, false).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18.1' of attribute 'age' should be greater than 19`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})

	Convey("Given I call the method ValidateMinimumFloat with a valid float and none exclusive", t, func() {

		validationError := ValidateMinimumFloat("age", 18, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumFloat with a valid float and exclusive", t, func() {

		validationError := ValidateMinimumFloat("age", 12.4, 6, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumFloat with a unvalid float and exclusive", t, func() {

		validationError := ValidateMinimumFloat("age", 18.1, 19, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18.1' of attribute 'age' should be greater or equal than 19`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})

	Convey("Given I call the method ValidateMinimumFloat with a unvalid float and exclusive", t, func() {

		validationError := ValidateMinimumFloat("age", 18, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidateMaximumInt(t *testing.T) {

	Convey("Given I call the method ValidateMaximumInt with a valid int and none exclusive", t, func() {

		validationError := ValidateMaximumInt("age", 12, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumInt with a unvalid int and none exclusive", t, func() {

		validationError := ValidateMaximumInt("age", 19, 18, false).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '19' of attribute 'age' should be less than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})

	Convey("Given I call the method ValidateMaximumInt with a valid float and none exclusive", t, func() {

		validationError := ValidateMaximumInt("age", 18, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumInt with a valid float and exclusive", t, func() {

		validationError := ValidateMaximumInt("age", 12, 18, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumInt with a unvalid float and exclusive", t, func() {

		validationError := ValidateMaximumInt("age", 19, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '19' of attribute 'age' should be less or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})

	Convey("Given I call the method ValidateMaximumInt with a unvalid float and exclusive", t, func() {

		validationError := ValidateMaximumInt("age", 18, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be less or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidateMinimumInt(t *testing.T) {

	Convey("Given I call the method ValidateMinimumInt with a valid float and none exclusive", t, func() {

		validationError := ValidateMinimumInt("age", 12, 6, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumInt with a unvalid float and none exclusive", t, func() {

		validationError := ValidateMinimumInt("age", 18, 19, false).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater than 19`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})

	Convey("Given I call the method ValidateMinimumInt with a valid float and none exclusive", t, func() {

		validationError := ValidateMinimumInt("age", 18, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumInt with a valid float and exclusive", t, func() {

		validationError := ValidateMinimumInt("age", 12, 6, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumInt with a unvalid float and exclusive", t, func() {

		validationError := ValidateMinimumInt("age", 18, 19, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater or equal than 19`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})

	Convey("Given I call the method ValidateMinimumInt with a unvalid float and exclusive", t, func() {

		validationError := ValidateMinimumInt("age", 18, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidateMaximumLength(t *testing.T) {

	Convey("Given I call the method ValidateMaximumLength with a valid length and none exclusive", t, func() {

		validationError := ValidateMaximumLength("name", "Alexandre", 20, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumLength with a unvalid length and none exclusive", t, func() {

		validationError := ValidateMaximumLength("name", "Alexandre", 1, false).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be less than 1 chars long`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})

	Convey("Given I call the method ValidateMaximumLength with a valid length and none exclusive", t, func() {

		validationError := ValidateMaximumLength("name", "Alexandre", 9, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumLength with a valid length and exclusive", t, func() {

		validationError := ValidateMaximumLength("name", "Alexandre", 18, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMaximumLength with a unvalid length and exclusive", t, func() {

		validationError := ValidateMaximumLength("name", "Alexandre", 1, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be less or equal than 1 chars long`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})

	Convey("Given I call the method ValidateMaximumLength with a unvalid length and exclusive", t, func() {

		validationError := ValidateMaximumLength("name", "Alexandre", 9, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be less or equal than 9 chars long`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})
}

func TestValidator_ValidateMinimumLength(t *testing.T) {

	Convey("Given I call the method ValidateMinimumLength with a valid length and none exclusive", t, func() {

		validationError := ValidateMinimumLength("name", "Alexandre", 6, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumLength with a unvalid length and none exclusive", t, func() {

		validationError := ValidateMinimumLength("name", "Alexandre", 19, false).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be greater than 19 chars long`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})

	Convey("Given I call the method ValidateMinimumLength with a valid length and none exclusive", t, func() {

		validationError := ValidateMinimumLength("name", "Alexandre", 9, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumLength with a valid length and exclusive", t, func() {

		validationError := ValidateMinimumLength("name", "Alexandre", 6, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateMinimumLength with a unvalid length and exclusive", t, func() {

		validationError := ValidateMinimumLength("name", "Alexandre", 19, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be greater or equal than 19 chars long`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})

	Convey("Given I call the method ValidateMinimumLength with a unvalid length and exclusive", t, func() {

		validationError := ValidateMinimumLength("name", "Alexandre", 9, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be greater or equal than 9 chars long`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})
}

func TestValidator_ValidateStringInList(t *testing.T) {

	Convey("Given I call the method ValidateStringInList with a valid string", t, func() {

		validationError := ValidateStringInList("name", "Alexandre", []string{"Dimitri", "Alexandre", "Antoine"}, false)

		Convey("Then I should get nil in return", func() {
			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateStringInList with an empty string and autogenerated", t, func() {

		validationError := ValidateStringInList("name", "", []string{"Dimitri", "Alexandre", "Antoine"}, true)

		Convey("Then I should get nil in return", func() {
			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateStringInList with a unvalid string", t, func() {

		validationError := ValidateStringInList("name", "Alexandre", []string{"Dimitri", "Antoine"}, false).(Error)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' is not in list '[Dimitri Antoine]'`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})
}

func TestValidator_ValidateIntInList(t *testing.T) {

	Convey("Given I call the method ValidateIntInList with a valid int", t, func() {

		validationError := ValidateIntInList("age", 18, []int{31, 12, 18})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateIntInList with a unvalid int", t, func() {

		validationError := ValidateIntInList("age", 18, []int{31, 12}).(Error)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' is not in list '[31 12]'`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidateFloatInList(t *testing.T) {

	Convey("Given I call the method ValidateFloatInList with a valid float", t, func() {

		validationError := ValidateFloatInList("age", 18.1, []float64{31, 12, 18.1})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateFloatInList with a unvalid float", t, func() {

		validationError := ValidateFloatInList("age", 18.3, []float64{31, 12}).(Error)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18.3' of attribute 'age' is not in list '[31 12]'`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "age"})
		})
	})
}

func TestValidator_ValidatePattern(t *testing.T) {

	Convey("Given I call the method ValidatePattern with a valid string", t, func() {

		validationError := ValidatePattern("name", "Alexandre", "Alexandre", "should be 'Alexandre'", true)

		Convey("Then I should get nil in return", func() {
			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidatePattern with a valid string", t, func() {

		validationError := ValidatePattern("name", "Alexandre", "Antoine", "should be 'Alexandre'", true).(Error)

		Convey("Then I should get nil in return", func() {
			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be 'Alexandre'`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
			So(validationError.Data, ShouldResemble, map[string]string{"attribute": "name"})
		})
	})

	Convey("Given I call the method ValidatePattern with a valid string", t, func() {

		validationError := ValidatePattern("name", "", "Antoine", "should be 'Alexandre'", false)

		Convey("Then I should get nil in return", func() {
			So(validationError, ShouldBeNil)
		})
	})
}

func TestValidator_ValidateFloatInMap(t *testing.T) {

	Convey("Given I call the method ValidateFloatInMap with a valid float", t, func() {

		validationError := ValidateFloatInMap("age", 18.1, map[float64]any{float64(18.1): true})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateFloatInMap with a unvalid float", t, func() {

		validationError := ValidateFloatInMap("age", 18.3, map[float64]any{float64(32.1): true}).(Error)

		Convey("Then I should not get nil in return", func() {
			So(validationError, ShouldNotBeNil)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
		})
	})
}

func TestValidator_ValidateIntInMap(t *testing.T) {

	Convey("Given I call the method ValidateIntInMap with a valid float", t, func() {

		validationError := ValidateIntInMap("age", 666, map[int]any{666: true})

		Convey("Then I should get nil in return", func() {
			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateIntInMap with a unvalid float", t, func() {

		validationError := ValidateIntInMap("age", 666, map[int]any{}).(Error)

		Convey("Then I should not get nil in return", func() {
			So(validationError, ShouldNotBeNil)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
		})
	})
}

func TestValidator_ValidateStringInMap(t *testing.T) {

	Convey("Given I call the method ValidateStringInMap with a valid float", t, func() {

		validationError := ValidateStringInMap("age", "666", map[string]any{"666": true}, false)

		Convey("Then I should get nil in return", func() {
			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidateStringInMap with a unvalid float", t, func() {

		validationError := ValidateStringInMap("age", "666", map[string]any{}, false).(Error)

		Convey("Then I should not get nil in return", func() {
			So(validationError, ShouldNotBeNil)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
		})
	})

	Convey("Given I call the method ValidateStringInMap with an empty value and mark it as autogen", t, func() {

		validationError := ValidateStringInMap("age", "", map[string]any{}, true)

		Convey("Then I should get nil in return", func() {
			So(validationError, ShouldBeNil)
		})
	})
}
