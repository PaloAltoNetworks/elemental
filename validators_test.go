// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMethodValidateRequiredIntWithValidInt(t *testing.T) {

	Convey("Given I call the method ValidateRequiredInt with a valid int", t, func() {

		validationError := ValidateRequiredInt("age", 15)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateRequiredIntWithUnvalidInt(t *testing.T) {

	Convey("Given I call the method ValidateRequiredInt with a nonvalid int", t, func() {

		validationError := ValidateRequiredInt("age", 0)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, "Attribute 'age' is required")
		})
	})
}

func TestMethodValidateRequiredFloatWithValidFloat(t *testing.T) {

	Convey("Given I call the method ValidateRequiredFloat with a valid float", t, func() {

		validationError := ValidateRequiredFloat("age", 15.3)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateRequiredFloatWithUnvalidFloat(t *testing.T) {

	Convey("Given I call the method ValidateRequiredFloat with a nonvalid float", t, func() {

		validationError := ValidateRequiredFloat("age", 0)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, "Attribute 'age' is required")
		})
	})
}

func TestMethodValidateRequiredStringWithValidString(t *testing.T) {

	Convey("Given I call the method ValidateRequiredString with a valid string", t, func() {

		validationError := ValidateRequiredString("name", "Alexandre")

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateRequiredStringWithUnvalidString(t *testing.T) {

	Convey("Given I call the method ValidateRequiredString with a nonvalid string", t, func() {

		validationError := ValidateRequiredString("name", "")

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, "Attribute 'name' is required")
		})
	})
}

func TestMethodValidateMaximumFloatNonExclusiveWithValidFloat(t *testing.T) {

	Convey("Given I call the method ValidateMaximumFloat with a valid float and none exclusive", t, func() {
		validationError := ValidateMaximumFloat("age", 12.4, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumFloatNonExclusiveWithUnvalidFloat(t *testing.T) {

	Convey("Given I call the method ValidateMaximumFloat with a unvalid float and none exclusive", t, func() {
		validationError := ValidateMaximumFloat("age", 18.1, 18, false)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18.1' of attribute 'age' should be less than 18`)
		})
	})
}

func TestMethodValidateMaximumFloatNonExclusiveWithValidFloatEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMaximumFloat with a valid float and none exclusive", t, func() {
		validationError := ValidateMaximumFloat("age", 18, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumFloatExclusiveWithValidFloat(t *testing.T) {

	Convey("Given I call the method ValidateMaximumFloat with a valid float and exclusive", t, func() {
		validationError := ValidateMaximumFloat("age", 12.4, 18, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumFloatExclusiveWithUnvalidFloat(t *testing.T) {

	Convey("Given I call the method ValidateMaximumFloat with a unvalid float and exclusive", t, func() {
		validationError := ValidateMaximumFloat("age", 18.1, 18, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18.1' of attribute 'age' should be less or equal than 18`)
		})
	})
}

func TestMethodValidateMaximumFloatExclusiveWithValidFloatEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMaximumFloat with a unvalid float and exclusive", t, func() {
		validationError := ValidateMaximumFloat("age", 18, 18, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be less or equal than 18`)
		})
	})
}

func TestMethodValidateMinimumFloatNonExclusiveWithValidFloat(t *testing.T) {

	Convey("Given I call the method ValidateMinimumFloat with a valid float and none exclusive", t, func() {
		validationError := ValidateMinimumFloat("age", 12.4, 6.1, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumFloatNonExclusiveWithUnvalidFloat(t *testing.T) {

	Convey("Given I call the method ValidateMinimumFloat with a unvalid float and none exclusive", t, func() {
		validationError := ValidateMinimumFloat("age", 18.1, 19, false)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18.1' of attribute 'age' should be greater than 19`)
		})
	})
}

func TestMethodValidateMinimumFloatNonExclusiveWithValidFloatEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMinimumFloat with a valid float and none exclusive", t, func() {
		validationError := ValidateMinimumFloat("age", 18, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumFloatExclusiveWithValidFloat(t *testing.T) {

	Convey("Given I call the method ValidateMinimumFloat with a valid float and exclusive", t, func() {
		validationError := ValidateMinimumFloat("age", 12.4, 6, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumFloatExclusiveWithUnvalidFloat(t *testing.T) {

	Convey("Given I call the method ValidateMinimumFloat with a unvalid float and exclusive", t, func() {
		validationError := ValidateMinimumFloat("age", 18.1, 19, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18.1' of attribute 'age' should be greater or equal than 19`)
		})
	})
}

func TestMethodValidateMinimumFloatExclusiveWithValidFloatEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMinimumFloat with a unvalid float and exclusive", t, func() {
		validationError := ValidateMinimumFloat("age", 18, 18, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater or equal than 18`)
		})
	})
}

func TestMethodValidateMaximumIntNonExclusiveWithValidInt(t *testing.T) {

	Convey("Given I call the method ValidateMaximumInt with a valid int and none exclusive", t, func() {
		validationError := ValidateMaximumInt("age", 12, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumIntNonExclusiveWithUnvalidInt(t *testing.T) {

	Convey("Given I call the method ValidateMaximumInt with a unvalid int and none exclusive", t, func() {
		validationError := ValidateMaximumInt("age", 19, 18, false)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '19' of attribute 'age' should be less than 18`)
		})
	})
}

func TestMethodValidateMaximumIntNonExclusiveWithValidIntEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMaximumInt with a valid float and none exclusive", t, func() {
		validationError := ValidateMaximumInt("age", 18, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumIntExclusiveWithValidInt(t *testing.T) {

	Convey("Given I call the method ValidateMaximumInt with a valid float and exclusive", t, func() {
		validationError := ValidateMaximumInt("age", 12, 18, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumIntExclusiveWithUnvalidInt(t *testing.T) {

	Convey("Given I call the method ValidateMaximumInt with a unvalid float and exclusive", t, func() {
		validationError := ValidateMaximumInt("age", 19, 18, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '19' of attribute 'age' should be less or equal than 18`)
		})
	})
}

func TestMethodValidateMaximumIntExclusiveWithValidIntEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMaximumInt with a unvalid float and exclusive", t, func() {
		validationError := ValidateMaximumInt("age", 18, 18, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be less or equal than 18`)
		})
	})
}

func TestMethodValidateMinimumIntNonExclusiveWithValidInt(t *testing.T) {

	Convey("Given I call the method ValidateMinimumInt with a valid float and none exclusive", t, func() {
		validationError := ValidateMinimumInt("age", 12, 6, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumIntNonExclusiveWithUnvalidInt(t *testing.T) {

	Convey("Given I call the method ValidateMinimumInt with a unvalid float and none exclusive", t, func() {
		validationError := ValidateMinimumInt("age", 18, 19, false)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater than 19`)
		})
	})
}

func TestMethodValidateMinimumIntNonExclusiveWithValidIntEqualToTheMin(t *testing.T) {

	Convey("Given I call the method ValidateMinimumInt with a valid float and none exclusive", t, func() {
		validationError := ValidateMinimumInt("age", 18, 18, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumIntExclusiveWithValidInt(t *testing.T) {

	Convey("Given I call the method ValidateMinimumInt with a valid float and exclusive", t, func() {
		validationError := ValidateMinimumInt("age", 12, 6, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumIntExclusiveWithUnvalidInt(t *testing.T) {

	Convey("Given I call the method ValidateMinimumInt with a unvalid float and exclusive", t, func() {
		validationError := ValidateMinimumInt("age", 18, 19, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater or equal than 19`)
		})
	})
}

func TestMethodValidateMinimumIntExclusiveWithValidIntEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMinimumInt with a unvalid float and exclusive", t, func() {
		validationError := ValidateMinimumInt("age", 18, 18, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' should be greater or equal than 18`)
		})
	})
}

func TestMethodValidateMaximumLengthNonExclusiveWithValidLength(t *testing.T) {

	Convey("Given I call the method ValidateMaximumLength with a valid length and none exclusive", t, func() {
		validationError := ValidateMaximumLength("name", "Alexandre", 20, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumLengthNonExclusiveWithUnvalidLength(t *testing.T) {

	Convey("Given I call the method ValidateMaximumLength with a unvalid length and none exclusive", t, func() {
		validationError := ValidateMaximumLength("name", "Alexandre", 1, false)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be less than 1 chars long`)
		})
	})
}

func TestMethodValidateMaximumLengthNonExclusiveWithValidLengthEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMaximumLength with a valid length and none exclusive", t, func() {
		validationError := ValidateMaximumLength("name", "Alexandre", 9, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumLengthExclusiveWithValidLength(t *testing.T) {

	Convey("Given I call the method ValidateMaximumLength with a valid length and exclusive", t, func() {
		validationError := ValidateMaximumLength("name", "Alexandre", 18, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMaximumLengthExclusiveWithUnvalidLength(t *testing.T) {

	Convey("Given I call the method ValidateMaximumLength with a unvalid length and exclusive", t, func() {
		validationError := ValidateMaximumLength("name", "Alexandre", 1, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be less or equal than 1 chars long`)
		})
	})
}

func TestMethodValidateMaximumLengthExclusiveWithValidLengthEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMaximumLength with a unvalid length and exclusive", t, func() {
		validationError := ValidateMaximumLength("name", "Alexandre", 9, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be less or equal than 9 chars long`)
		})
	})
}

func TestMethodValidateMinimumLengthNonExclusiveWithValidLength(t *testing.T) {

	Convey("Given I call the method ValidateMinimumLength with a valid length and none exclusive", t, func() {
		validationError := ValidateMinimumLength("name", "Alexandre", 6, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumLengthNonExclusiveWithUnvalidLength(t *testing.T) {

	Convey("Given I call the method ValidateMinimumLength with a unvalid length and none exclusive", t, func() {
		validationError := ValidateMinimumLength("name", "Alexandre", 19, false)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be greater than 19 chars long`)
		})
	})
}

func TestMethodValidateMinimumLengthNonExclusiveWithValidLengthEqualToTheMin(t *testing.T) {

	Convey("Given I call the method ValidateMinimumLength with a valid length and none exclusive", t, func() {
		validationError := ValidateMinimumLength("name", "Alexandre", 9, false)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumLengthExclusiveWithValidLength(t *testing.T) {

	Convey("Given I call the method ValidateMinimumLength with a valid length and exclusive", t, func() {
		validationError := ValidateMinimumLength("name", "Alexandre", 6, true)

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateMinimumLengthExclusiveWithUnvalidLength(t *testing.T) {

	Convey("Given I call the method ValidateMinimumLength with a unvalid length and exclusive", t, func() {
		validationError := ValidateMinimumLength("name", "Alexandre", 19, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be greater or equal than 19 chars long`)
		})
	})
}

func TestMethodValidateMinimumLengthExclusiveWithValidLengthEqualToTheMax(t *testing.T) {

	Convey("Given I call the method ValidateMinimumLength with a unvalid length and exclusive", t, func() {
		validationError := ValidateMinimumLength("name", "Alexandre", 9, true)

		Convey("Then I should not get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should be greater or equal than 9 chars long`)
		})
	})
}

func TestMethodValidateStringInListWithValidString(t *testing.T) {

	Convey("Given I call the method ValidateStringInList with a valid string", t, func() {
		validationError := ValidateStringInList("name", "Alexandre", []string{"Dimitri", "Alexandre", "Antoine"})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateStringInListWithUnvalidString(t *testing.T) {

	Convey("Given I call the method ValidateStringInList with a unvalid string", t, func() {
		validationError := ValidateStringInList("name", "Alexandre", []string{"Dimitri", "Antoine"})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' is not in list '[Dimitri Antoine]'`)

		})
	})
}

func TestMethodValidateIntInListWithValidInt(t *testing.T) {

	Convey("Given I call the method ValidateIntInList with a valid int", t, func() {
		validationError := ValidateIntInList("age", 18, []int{31, 12, 18})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateIntInListWithUnvalidInt(t *testing.T) {

	Convey("Given I call the method ValidateIntInList with a unvalid int", t, func() {
		validationError := ValidateIntInList("age", 18, []int{31, 12})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18' of attribute 'age' is not in list '[31 12]'`)

		})
	})
}

func TestMethodValidateFloatInListWithValidFloat(t *testing.T) {

	Convey("Given I call the method ValidateFloatInList with a valid float", t, func() {
		validationError := ValidateFloatInList("age", 18.1, []float64{31, 12, 18.1})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidateFloatInListWithUnvalidFloat(t *testing.T) {

	Convey("Given I call the method ValidateFloatInList with a unvalid float", t, func() {
		validationError := ValidateFloatInList("age", 18.3, []float64{31, 12})

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "age")
			So(validationError.Description, ShouldEqual, `Data '18.3' of attribute 'age' is not in list '[31 12]'`)

		})
	})
}

func TestMethodValidatePatternWithValidString(t *testing.T) {

	Convey("Given I call the method ValidatePattern with a valid string", t, func() {
		validationError := ValidatePattern("name", "Alexandre", "Alexandre")

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldBeNil)
		})
	})
}

func TestMethodValidatePatternWithUnvalidString(t *testing.T) {

	Convey("Given I call the method ValidatePattern with a valid string", t, func() {
		validationError := ValidatePattern("name", "Alexandre", "Antoine")

		Convey("Then I should get nil in return", func() {

			So(validationError, ShouldNotBeNil)
			So(validationError.Subject, ShouldEqual, "name")
			So(validationError.Description, ShouldEqual, `Data 'Alexandre' of attribute 'name' should match 'Antoine'`)
		})
	})
}
