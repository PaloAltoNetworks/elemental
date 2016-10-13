// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

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
		})
	})

	Convey("Given I call the method ValidateMaximumFloat with a unvalid float and exclusive", t, func() {

		validationError := ValidateMaximumFloat("age", 18, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be less or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
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
		})
	})

	Convey("Given I call the method ValidateMinimumFloat with a unvalid float and exclusive", t, func() {

		validationError := ValidateMinimumFloat("age", 18, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
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
		})
	})

	Convey("Given I call the method ValidateMaximumInt with a unvalid float and exclusive", t, func() {

		validationError := ValidateMaximumInt("age", 18, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be less or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
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
		})
	})

	Convey("Given I call the method ValidateMinimumInt with a unvalid float and exclusive", t, func() {

		validationError := ValidateMinimumInt("age", 18, 18, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater or equal than 18`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
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
		})
	})

	Convey("Given I call the method ValidateMaximumLength with a unvalid length and exclusive", t, func() {

		validationError := ValidateMaximumLength("name", "Alexandre", 9, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be less or equal than 9 chars long`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
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
		})
	})

	Convey("Given I call the method ValidateMinimumLength with a unvalid length and exclusive", t, func() {

		validationError := ValidateMinimumLength("name", "Alexandre", 9, true).(Error)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be greater or equal than 9 chars long`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
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
		})
	})
}

func TestValidator_ValidatePattern(t *testing.T) {

	Convey("Given I call the method ValidatePattern with a valid string", t, func() {

		validationError := ValidatePattern("name", "Alexandre", "Alexandre")

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})

	Convey("Given I call the method ValidatePattern with a valid string", t, func() {

		validationError := ValidatePattern("name", "Alexandre", "Antoine").(Error)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should match 'Antoine'`)
			So(validationError.Code, ShouldEqual, http.StatusUnprocessableEntity)
		})
	})
}
