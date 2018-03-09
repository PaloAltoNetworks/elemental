package elemental

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils_extractFieldNames(t *testing.T) {

	Convey("Given I have a list", t, func() {

		l1 := NewList()

		Convey("When I extract the fields", func() {

			fields := extractFieldNames(l1)

			fmt.Println(fields)

			Convey("Then all fields should be present", func() {
				So(len(fields), ShouldEqual, 12)
				So(fields, ShouldContain, "ID")
				So(fields, ShouldContain, "Description")
				So(fields, ShouldContain, "Name")
				So(fields, ShouldContain, "ParentID")
				So(fields, ShouldContain, "ParentType")
				So(fields, ShouldContain, "CreationOnly")
				So(fields, ShouldContain, "ReadOnly")
				So(fields, ShouldContain, "Unexposed")
				So(fields, ShouldContain, "Date")
				So(fields, ShouldContain, "Slice")
				So(fields, ShouldContain, "ModelVersion")
				So(fields, ShouldContain, "Mutex")
			})
		})
	})
}

func TestUtils_areFieldValuesEqual(t *testing.T) {

	Convey("Given I have 2 list", t, func() {

		l1 := NewList()
		l2 := NewList()

		Convey("When I set the same name", func() {

			l1.Name = "list1"
			l2.Name = "list1"

			Convey("Then the values should be equal", func() {
				So(areFieldValuesEqual("Name", l1, l2), ShouldBeTrue)
			})
		})

		Convey("When I set a different name", func() {

			l1.Name = "list1"
			l2.Name = "list2"

			Convey("Then the values should not be equal", func() {
				So(areFieldValuesEqual("Name", l1, l2), ShouldBeFalse)
			})
		})

		Convey("When I set a same time", func() {

			l1.Date = time.Date(2009, time.November, 10, 10, 0, 0, 0, time.UTC)
			l2.Date = time.Date(2009, time.November, 10, 5, 0, 0, 0, time.FixedZone("Eastern", -5*3600))

			Convey("Then the values should not be equal", func() {
				So(areFieldValuesEqual("Date", l1, l2), ShouldBeTrue)
			})
		})

		Convey("When I set a different time", func() {

			l1.Date = time.Now()
			l2.Date = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

			Convey("Then the values should not be equal", func() {
				So(areFieldValuesEqual("Date", l1, l2), ShouldBeFalse)
			})
		})

		Convey("When I set a same slice", func() {

			l1.Slice = []string{"a", "b", "c"}
			l2.Slice = []string{"a", "b", "c"}

			Convey("Then the values should not be equal", func() {
				So(areFieldValuesEqual("Slice", l1, l2), ShouldBeTrue)
			})
		})

		Convey("When I set a different slice", func() {

			l1.Slice = []string{"a", "b", "c"}
			l2.Slice = []string{"a", "b"}

			Convey("Then the values should not be equal", func() {
				So(areFieldValuesEqual("Slice", l1, l2), ShouldBeFalse)
			})
		})
	})
}

func TestUtils_areFieldsValueEqualValue(t *testing.T) {
	Convey("Given I have a struct", t, func() {

		type S struct {
			S   string
			B   bool
			I   int
			F   float32
			A   []string
			M   map[string]string
			T   time.Time
			Sub *S
		}

		Convey("When I set all zero values", func() {

			var t time.Time
			s := &S{"", false, 0, 0.0, nil, nil, t, &S{}}

			Convey("Then areFieldsValueEqualValue on S should return true", func() {
				So(areFieldsValueEqualValue("S", s, ""), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on B should return true", func() {
				So(areFieldsValueEqualValue("B", s, false), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on I should return true", func() {
				So(areFieldsValueEqualValue("I", s, 0), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on F should return true", func() {
				So(areFieldsValueEqualValue("F", s, float32(0.0)), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on A should return true", func() {
				So(areFieldsValueEqualValue("A", s, nil), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on M should return true", func() {
				So(areFieldsValueEqualValue("M", s, nil), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on T should return true", func() {
				So(areFieldsValueEqualValue("T", s, t), ShouldBeTrue)
			})
		})

		Convey("When I set all non zero values with equal values", func() {

			t := time.Now()
			s := &S{"hello", true, 1, 1.0, []string{"a"}, map[string]string{"a": "b"}, t, &S{S: "nope"}}

			Convey("Then areFieldsValueEqualValue on S should return true", func() {
				So(areFieldsValueEqualValue("S", s, "hello"), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on B should return true", func() {
				So(areFieldsValueEqualValue("B", s, true), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on I should return true", func() {
				So(areFieldsValueEqualValue("I", s, 1), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on F should return true", func() {
				So(areFieldsValueEqualValue("F", s, float32(1)), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on A should return true", func() {
				So(areFieldsValueEqualValue("A", s, []string{"a"}), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on M should return true", func() {
				So(areFieldsValueEqualValue("M", s, map[string]string{"a": "b"}), ShouldBeTrue)
			})

			Convey("Then areFieldsValueEqualValue on T should return true", func() {
				So(areFieldsValueEqualValue("T", s, t), ShouldBeTrue)
			})
		})

		Convey("When I set all non zero values with not equal values", func() {
			s := &S{"hello", true, 1, 1.0, []string{"a"}, map[string]string{"a": "b"}, time.Now(), &S{S: "nope"}}

			Convey("Then areFieldsValueEqualValue on S should return false", func() {
				So(areFieldsValueEqualValue("S", s, "hello1"), ShouldBeFalse)
			})

			Convey("Then areFieldsValueEqualValue on B should return false", func() {
				So(areFieldsValueEqualValue("B", s, false), ShouldBeFalse)
			})

			Convey("Then areFieldsValueEqualValue on I should return false", func() {
				So(areFieldsValueEqualValue("I", s, 2), ShouldBeFalse)
			})

			Convey("Then areFieldsValueEqualValue on F should return false", func() {
				So(areFieldsValueEqualValue("F", s, float32(2)), ShouldBeFalse)
			})

			Convey("Then areFieldsValueEqualValue on A should return false", func() {
				So(areFieldsValueEqualValue("A", s, []string{"b"}), ShouldBeFalse)
				So(areFieldsValueEqualValue("A", s, []string{"b", "a"}), ShouldBeFalse)
			})

			Convey("Then areFieldsValueEqualValue on M should return false", func() {
				So(areFieldsValueEqualValue("M", s, map[string]string{"b": "a"}), ShouldBeFalse)
			})
		})
	})
}

func TestUtils_isFieldValueZero(t *testing.T) {

	Convey("Given I have a struct", t, func() {

		type S struct {
			S   string
			B   bool
			I   int
			F   float32
			A   []string
			M   map[string]string
			T   time.Time
			Sub *S
		}

		Convey("When I set all zero values", func() {

			var t time.Time
			s := &S{"", false, 0, 0.0, nil, nil, t, &S{}}

			Convey("Then isFieldValueZero on S should return true", func() {
				So(isFieldValueZero("S", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on B should return true", func() {
				So(isFieldValueZero("B", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on I should return true", func() {
				So(isFieldValueZero("I", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on F should return true", func() {
				So(isFieldValueZero("F", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on A should return true", func() {
				So(isFieldValueZero("A", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on M should return true", func() {
				So(isFieldValueZero("M", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on T should return true", func() {
				So(isFieldValueZero("T", s), ShouldBeTrue)
			})
		})

		Convey("When I set all non zero values", func() {

			t := time.Now()
			s := &S{"hello", true, 1, 1.0, []string{"a"}, map[string]string{"a": "b"}, t, &S{S: "nope"}}

			Convey("Then isFieldValueZero on S should return false", func() {
				So(isFieldValueZero("S", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on B should return false", func() {
				So(isFieldValueZero("B", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on I should return false", func() {
				So(isFieldValueZero("I", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on F should return false", func() {
				So(isFieldValueZero("F", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on A should return false", func() {
				So(isFieldValueZero("A", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on M should return false", func() {
				So(isFieldValueZero("M", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on T should return false", func() {
				So(isFieldValueZero("T", s), ShouldBeFalse)
			})
		})
	})
}

type timeList struct {
	Time    time.Time
	Times   []time.Time
	String  string
	Strings []string
	Int     int
	Ints    []int
	Bool    bool
	Bools   []bool
}

func TestVerify_areFieldValuesEqualWithEncoding(t *testing.T) {

	Convey("Given I have 2 structs with list of time", t, func() {

		now := time.Now()
		s1 := &timeList{
			Time:    now,
			Times:   []time.Time{now, now},
			String:  "A",
			Strings: []string{"a", "b"},
			Int:     42,
			Ints:    []int{1, 2},
			Bool:    true,
			Bools:   []bool{true, false},
		}
		d1, _ := json.Marshal(s1)
		_ = json.Unmarshal(d1, s1)

		time.Local = time.FixedZone("PST", 0)
		s2 := &timeList{
			Time:    now,
			Times:   []time.Time{now, now},
			String:  "A",
			Strings: []string{"a", "b"},
			Int:     42,
			Ints:    []int{1, 2},
			Bool:    true,
			Bools:   []bool{true, false},
		}
		d2, _ := json.Marshal(s2)
		_ = json.Unmarshal(d2, s2)

		Convey("When I call areFieldValuesEqual on Time", func() {

			ok := areFieldValuesEqual("Time", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call areFieldValuesEqual again", func() {

				// delta < 1s precision is lost during encoding.
				s1.Time = s1.Time.Add(1 * time.Second)
				ok := areFieldValuesEqual("Time", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call areFieldValuesEqual on Times", func() {

			ok := areFieldValuesEqual("Times", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call areFieldValuesEqual again", func() {

				// delta < 1s precision is lost during encoding.
				s1.Times[1] = s1.Times[1].Add(1 * time.Second)
				ok := areFieldValuesEqual("Times", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call areFieldValuesEqual on String", func() {

			ok := areFieldValuesEqual("String", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call areFieldValuesEqual again", func() {

				s1.String = "B"
				ok := areFieldValuesEqual("String", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call areFieldValuesEqual on Strings", func() {

			ok := areFieldValuesEqual("Strings", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call areFieldValuesEqual again", func() {

				s1.Strings[0] = "B"
				ok := areFieldValuesEqual("Strings", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call areFieldValuesEqual on Int", func() {

			ok := areFieldValuesEqual("Int", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call areFieldValuesEqual again", func() {

				s1.Int++
				ok := areFieldValuesEqual("Int", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call areFieldValuesEqual on Ints", func() {

			ok := areFieldValuesEqual("Ints", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call areFieldValuesEqual again", func() {

				s1.Ints[0]++
				ok := areFieldValuesEqual("Ints", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call areFieldValuesEqual on Bool", func() {

			ok := areFieldValuesEqual("Bool", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call areFieldValuesEqual again", func() {

				s1.Bool = !s1.Bool
				ok := areFieldValuesEqual("Bool", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call areFieldValuesEqual on Bools", func() {

			ok := areFieldValuesEqual("Bools", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call areFieldValuesEqual again", func() {

				s1.Bools[0] = !s1.Bools[0]
				ok := areFieldValuesEqual("Bools", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})
	})
}
