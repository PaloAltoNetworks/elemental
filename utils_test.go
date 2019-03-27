package elemental

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils_ExtractFieldNames(t *testing.T) {

	Convey("Given I have a list", t, func() {

		l1 := NewList()

		Convey("When I extract the fields", func() {

			fields := ExtractFieldNames(l1)

			fmt.Println(fields)

			Convey("Then all fields should be present", func() {
				So(len(fields), ShouldEqual, 13)
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
				So(fields, ShouldContain, "Secret")
			})
		})
	})
}

func TestUtils_AreFieldValuesEqual(t *testing.T) {

	type testStruct struct {
		String  string
		Strings []string
		Time    time.Time
		Map     map[string]interface{}
		Int     int
		Float   float64
	}

	Convey("Given I have 2 list", t, func() {

		s1 := &testStruct{}
		s2 := &testStruct{}

		Convey("When I set the same name", func() {

			s1.String = "v1"
			s2.String = "v1"

			Convey("Then the values should be equal", func() {
				So(AreFieldValuesEqual("String", s1, s2), ShouldBeTrue)
			})
		})

		Convey("When I set a different name", func() {

			s1.String = "v1"
			s2.String = "v2"

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("String", s1, s2), ShouldBeFalse)
			})
		})

		Convey("When I set a same int", func() {

			s1.Int = 42
			s2.Int = 42

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Int", s1, s2), ShouldBeTrue)
			})
		})

		Convey("When I set a different int", func() {

			s1.Int = 42
			s2.Int = 43

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Int", s1, s2), ShouldBeFalse)
			})
		})

		Convey("When I set a same Float", func() {

			s1.Float = 42.42
			s2.Float = 42.42

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Float", s1, s2), ShouldBeTrue)
			})
		})

		Convey("When I set a different Float", func() {

			s1.Float = 42.42
			s2.Float = 42.43

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Float", s1, s2), ShouldBeFalse)
			})
		})

		Convey("When I set a same time", func() {

			s1.Time = time.Date(2009, time.November, 10, 10, 0, 0, 0, time.UTC)
			s2.Time = time.Date(2009, time.November, 10, 5, 0, 0, 0, time.FixedZone("Eastern", -5*3600))

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Time", s1, s2), ShouldBeTrue)
			})
		})

		Convey("When I set a different time", func() {

			s1.Time = time.Now()
			s2.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Time", s1, s2), ShouldBeFalse)
			})
		})

		Convey("When I set a same slice", func() {

			s1.Strings = []string{"a", "b", "c"}
			s2.Strings = []string{"a", "b", "c"}

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Strings", s1, s2), ShouldBeTrue)
			})
		})

		Convey("When I set a different slice", func() {

			s1.Strings = []string{"a", "b", "c"}
			s2.Strings = []string{"a", "b"}

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Strings", s1, s2), ShouldBeFalse)
			})
		})

		Convey("When I set a same map", func() {

			s1.Map = map[string]interface{}{"a": 1}
			s2.Map = map[string]interface{}{"a": 1}

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Map", s1, s2), ShouldBeTrue)
			})
		})

		Convey("When I set a different map with same len", func() {

			s1.Map = map[string]interface{}{"a": 1}
			s2.Map = map[string]interface{}{"a": 2}

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Map", s1, s2), ShouldBeFalse)
			})
		})

		Convey("When I set a different map with different len", func() {

			s1.Map = map[string]interface{}{"a": 1}
			s2.Map = map[string]interface{}{"a": 2, "b": 1}

			Convey("Then the values should not be equal", func() {
				So(AreFieldValuesEqual("Map", s1, s2), ShouldBeFalse)
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

			Convey("Then AreFieldsValueEqualValue on S should return true", func() {
				So(AreFieldsValueEqualValue("S", s, ""), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on B should return true", func() {
				So(AreFieldsValueEqualValue("B", s, false), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on I should return true", func() {
				So(AreFieldsValueEqualValue("I", s, 0), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on F should return true", func() {
				So(AreFieldsValueEqualValue("F", s, float32(0.0)), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on A should return true", func() {
				So(AreFieldsValueEqualValue("A", s, nil), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on M should return true", func() {
				So(AreFieldsValueEqualValue("M", s, nil), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on T should return true", func() {
				So(AreFieldsValueEqualValue("T", s, t), ShouldBeTrue)
			})
		})

		Convey("When I set all non zero values with equal values", func() {

			t := time.Now()
			s := &S{"hello", true, 1, 1.0, []string{"a"}, map[string]string{"a": "b"}, t, &S{S: "nope"}}

			Convey("Then AreFieldsValueEqualValue on S should return true", func() {
				So(AreFieldsValueEqualValue("S", s, "hello"), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on B should return true", func() {
				So(AreFieldsValueEqualValue("B", s, true), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on I should return true", func() {
				So(AreFieldsValueEqualValue("I", s, 1), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on F should return true", func() {
				So(AreFieldsValueEqualValue("F", s, float32(1)), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on A should return true", func() {
				So(AreFieldsValueEqualValue("A", s, []string{"a"}), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on M should return true", func() {
				So(AreFieldsValueEqualValue("M", s, map[string]string{"a": "b"}), ShouldBeTrue)
			})

			Convey("Then AreFieldsValueEqualValue on T should return true", func() {
				So(AreFieldsValueEqualValue("T", s, t), ShouldBeTrue)
			})
		})

		Convey("When I set all non zero values with not equal values", func() {
			s := &S{"hello", true, 1, 1.0, []string{"a"}, map[string]string{"a": "b"}, time.Now(), &S{S: "nope"}}

			Convey("Then AreFieldsValueEqualValue on S should return false", func() {
				So(AreFieldsValueEqualValue("S", s, "hello1"), ShouldBeFalse)
			})

			Convey("Then AreFieldsValueEqualValue on B should return false", func() {
				So(AreFieldsValueEqualValue("B", s, false), ShouldBeFalse)
			})

			Convey("Then AreFieldsValueEqualValue on I should return false", func() {
				So(AreFieldsValueEqualValue("I", s, 2), ShouldBeFalse)
			})

			Convey("Then AreFieldsValueEqualValue on F should return false", func() {
				So(AreFieldsValueEqualValue("F", s, float32(2)), ShouldBeFalse)
			})

			Convey("Then AreFieldsValueEqualValue on A should return false", func() {
				So(AreFieldsValueEqualValue("A", s, []string{"b"}), ShouldBeFalse)
				So(AreFieldsValueEqualValue("A", s, []string{"b", "a"}), ShouldBeFalse)
			})

			Convey("Then AreFieldsValueEqualValue on M should return false", func() {
				So(AreFieldsValueEqualValue("M", s, map[string]string{"b": "a"}), ShouldBeFalse)
			})
		})
	})
}

func TestUtils_IsFieldValueZero(t *testing.T) {

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

			Convey("Then IsFieldValueZero on S should return true", func() {
				So(IsFieldValueZero("S", s), ShouldBeTrue)
			})

			Convey("Then IsFieldValueZero on B should return true", func() {
				So(IsFieldValueZero("B", s), ShouldBeTrue)
			})

			Convey("Then IsFieldValueZero on I should return true", func() {
				So(IsFieldValueZero("I", s), ShouldBeTrue)
			})

			Convey("Then IsFieldValueZero on F should return true", func() {
				So(IsFieldValueZero("F", s), ShouldBeTrue)
			})

			Convey("Then IsFieldValueZero on A should return true", func() {
				So(IsFieldValueZero("A", s), ShouldBeTrue)
			})

			Convey("Then IsFieldValueZero on M should return true", func() {
				So(IsFieldValueZero("M", s), ShouldBeTrue)
			})

			Convey("Then IsFieldValueZero on T should return true", func() {
				So(IsFieldValueZero("T", s), ShouldBeTrue)
			})
		})

		Convey("When I set all non zero values", func() {

			t := time.Now()
			s := &S{"hello", true, 1, 1.0, []string{"a"}, map[string]string{"a": "b"}, t, &S{S: "nope"}}

			Convey("Then IsFieldValueZero on S should return false", func() {
				So(IsFieldValueZero("S", s), ShouldBeFalse)
			})

			Convey("Then IsFieldValueZero on B should return false", func() {
				So(IsFieldValueZero("B", s), ShouldBeFalse)
			})

			Convey("Then IsFieldValueZero on I should return false", func() {
				So(IsFieldValueZero("I", s), ShouldBeFalse)
			})

			Convey("Then IsFieldValueZero on F should return false", func() {
				So(IsFieldValueZero("F", s), ShouldBeFalse)
			})

			Convey("Then IsFieldValueZero on A should return false", func() {
				So(IsFieldValueZero("A", s), ShouldBeFalse)
			})

			Convey("Then IsFieldValueZero on M should return false", func() {
				So(IsFieldValueZero("M", s), ShouldBeFalse)
			})

			Convey("Then IsFieldValueZero on T should return false", func() {
				So(IsFieldValueZero("T", s), ShouldBeFalse)
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

func TestVerify_AreFieldValuesEqualWithEncoding(t *testing.T) {

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

		Convey("When I call AreFieldValuesEqual on Time", func() {

			ok := AreFieldValuesEqual("Time", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call AreFieldValuesEqual again", func() {

				// delta < 1s precision is lost during encoding.
				s1.Time = s1.Time.Add(1 * time.Second)
				ok := AreFieldValuesEqual("Time", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call AreFieldValuesEqual on Times", func() {

			ok := AreFieldValuesEqual("Times", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call AreFieldValuesEqual again", func() {

				// delta < 1s precision is lost during encoding.
				s1.Times[1] = s1.Times[1].Add(1 * time.Second)
				ok := AreFieldValuesEqual("Times", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call AreFieldValuesEqual on String", func() {

			ok := AreFieldValuesEqual("String", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call AreFieldValuesEqual again", func() {

				s1.String = "B"
				ok := AreFieldValuesEqual("String", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call AreFieldValuesEqual on Strings", func() {

			ok := AreFieldValuesEqual("Strings", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call AreFieldValuesEqual again", func() {

				s1.Strings[0] = "B"
				ok := AreFieldValuesEqual("Strings", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call AreFieldValuesEqual on Int", func() {

			ok := AreFieldValuesEqual("Int", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call AreFieldValuesEqual again", func() {

				s1.Int++
				ok := AreFieldValuesEqual("Int", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call AreFieldValuesEqual on Ints", func() {

			ok := AreFieldValuesEqual("Ints", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call AreFieldValuesEqual again", func() {

				s1.Ints[0]++
				ok := AreFieldValuesEqual("Ints", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call AreFieldValuesEqual on Bool", func() {

			ok := AreFieldValuesEqual("Bool", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call AreFieldValuesEqual again", func() {

				s1.Bool = !s1.Bool
				ok := AreFieldValuesEqual("Bool", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})

		Convey("When I call AreFieldValuesEqual on Bools", func() {

			ok := AreFieldValuesEqual("Bools", s1, s2)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I change one value and call AreFieldValuesEqual again", func() {

				s1.Bools[0] = !s1.Bools[0]
				ok := AreFieldValuesEqual("Bools", s1, s2)

				Convey("Then ok should be false", func() {
					So(ok, ShouldBeFalse)
				})
			})
		})
	})
}
