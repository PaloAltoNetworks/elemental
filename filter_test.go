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

func TestFilterComparator_NewFilterComparators(t *testing.T) {

	Convey("Given I create a empty FilterComparators with 2 comparators", t, func() {

		fc := FilterComparators{EqualComparator, InComparator}

		Convey("Then it should should not be empty", func() {
			So(len(fc), ShouldEqual, 2)
		})

		Convey("Then it should be correctly populated", func() {
			So(fc, ShouldResemble, FilterComparators{EqualComparator, InComparator})
		})

		Convey("When I use Then to add 2 other comparators", func() {

			fc = fc.add(ContainComparator, GreaterComparator)

			Convey("Then it should be correctly populated", func() {
				So(fc, ShouldResemble, FilterComparators{
					EqualComparator,
					InComparator,
					ContainComparator,
					GreaterComparator,
				})
			})
		})
	})
}

func TestFilter_NewFilter(t *testing.T) {

	Convey("Given I create a new filter", t, func() {

		f := NewFilter()

		Convey("Then the filter should implement the correct interfaces", func() {
			So(f, ShouldImplement, (*FilterKeyComposer)(nil))
			So(f, ShouldImplement, (*FilterValueComposer)(nil))
		})

		Convey("Then all values should be initialized", func() {
			So(f.keys, ShouldResemble, FilterKeys{})
			So(f.values, ShouldResemble, FilterValues{})
			So(f.operators, ShouldResemble, FilterOperators{})
			So(f.comparators, ShouldResemble, FilterComparators{})
			So(f.ands, ShouldResemble, SubFilters{})
			So(f.ors, ShouldResemble, SubFilters{})
		})
	})
}

func TestFilter_NewComposer(t *testing.T) {

	Convey("Given I create a new FilterComposer", t, func() {

		f := NewFilterComposer().Done()

		Convey("When I add the initial Equals statement", func() {

			f.WithKey("hello").Equals(1)

			Convey("Then the filter should be correctly populated", func() {
				So(f.Keys(), ShouldResemble, FilterKeys{"hello"})
				So(f.Values(), ShouldResemble, FilterValues{FilterValue{1}})
				So(f.Operators(), ShouldResemble, FilterOperators{AndOperator})
				So(f.Comparators(), ShouldResemble, FilterComparators{EqualComparator})
			})

			Convey("When I add a new GreaterThan statement", func() {

				f.WithKey("gt").GreaterThan(12)

				Convey("Then the filter should be correctly populated", func() {

					So(f.Keys(), ShouldResemble, FilterKeys{
						"hello",
						"gt",
					})
					So(f.Values(), ShouldResemble, FilterValues{
						FilterValue{1},
						FilterValue{12},
					})
					So(f.Operators(), ShouldResemble, FilterOperators{
						AndOperator,
						AndOperator,
					})
					So(f.Comparators(), ShouldResemble, FilterComparators{
						EqualComparator,
						GreaterComparator,
					})

					Convey("When I add a new LesserThan statement", func() {

						f.WithKey("lt").LesserThan(13)
						f.WithKey("gte").GreaterOrEqualThan(22)
						f.WithKey("lte").LesserOrEqualThan(23)

						Convey("Then the filter should be correctly populated", func() {
							So(f.Keys(), ShouldResemble, FilterKeys{
								"hello",
								"gt",
								"lt",
								"gte",
								"lte",
							})
							So(f.Values(), ShouldResemble, FilterValues{
								FilterValue{1},
								FilterValue{12},
								FilterValue{13},
								FilterValue{22},
								FilterValue{23},
							})
							So(f.Operators(), ShouldResemble, FilterOperators{
								AndOperator,
								AndOperator,
								AndOperator,
								AndOperator,
								AndOperator,
							})
							So(f.Comparators(), ShouldResemble, FilterComparators{
								EqualComparator,
								GreaterComparator,
								LesserComparator,
								GreaterOrEqualComparator,
								LesserOrEqualComparator,
							})
						})

						Convey("Then the string representation should be correct", func() {
							So(f.String(), ShouldEqual, `hello == 1 and gt > 12 and lt < 13 and gte >= 22 and lte <= 23`)
						})
					})
				})
			})

			Convey("When I add a new In statement", func() {

				f.WithKey("in").In("a", "b", "c")

				Convey("Then the filter should be correctly populated", func() {
					So(f.keys, ShouldResemble, FilterKeys{
						"hello",
						"in",
					})
					So(f.Values(), ShouldResemble, FilterValues{
						FilterValue{1},
						FilterValue{"a", "b", "c"},
					})
					So(f.Operators(), ShouldResemble, FilterOperators{
						AndOperator,
						AndOperator,
					})
					So(f.Comparators(), ShouldResemble, FilterComparators{
						EqualComparator,
						InComparator,
					})

					Convey("When I add a new Contains statement", func() {

						f.WithKey("ctn").Contains(false)

						Convey("Then the filter should be correctly populated", func() {
							So(f.Keys(), ShouldResemble, FilterKeys{
								"hello",
								"in",
								"ctn",
							})
							So(f.Values(), ShouldResemble, FilterValues{
								FilterValue{1},
								FilterValue{"a", "b", "c"},
								FilterValue{false},
							})
							So(f.Operators(), ShouldResemble, FilterOperators{
								AndOperator,
								AndOperator,
								AndOperator,
							})
							So(f.Comparators(), ShouldResemble, FilterComparators{
								EqualComparator,
								InComparator,
								ContainComparator,
							})

							Convey("Then the string representation should be correct", func() {
								So(f.String(), ShouldEqual, `hello == 1 and in in ["a", "b", "c"] and ctn contains [false]`)
							})
						})
					})
				})
			})

			Convey("When I add a new difference comparator", func() {
				f.WithKey("x").NotEquals(true)

				Convey("Then the filter should be correctly populated", func() {
					So(f.keys, ShouldResemble, FilterKeys{
						"hello",
						"x",
					})
					So(f.Values(), ShouldResemble, FilterValues{
						FilterValue{1},
						FilterValue{true},
					})
					So(f.Operators(), ShouldResemble, FilterOperators{
						AndOperator,
						AndOperator,
					})
					So(f.Comparators(), ShouldResemble, FilterComparators{
						EqualComparator,
						NotEqualComparator,
					})
					So(f.String(), ShouldEqual, `hello == 1 and x != true`)
				})
			})
		})
	})
}

func TestFilter_AppendToExisting(t *testing.T) {

	Convey("Given I have a simple filter", t, func() {

		f := NewFilterComposer().WithKey("a").Equals("b").Done()

		Convey("When I append mode", func() {

			f = f.WithKey("b").Equals("c").Done()

			Convey("Then f should be correct", func() {
				So(f.String(), ShouldEqual, `a == "b" and b == "c"`)
			})
		})
	})

	Convey("Given I have a composed filter", t, func() {

		f := NewFilterComposer().And(
			NewFilterComposer().WithKey("a").Equals("b").Done(),
		).Done()

		Convey("When I append mode", func() {

			f = f.WithKey("b").Equals("c").Done()

			Convey("Then f should be correct", func() {
				So(f.String(), ShouldEqual, `((a == "b")) and b == "c"`)
			})
		})
	})
}

func TestFilter_Date(t *testing.T) {

	Convey("Given I have a filter with date", t, func() {

		f := NewFilterComposer().WithKey("date").Equals(time.Time{}).Done()

		Convey("When I call String", func() {

			s := f.String()

			Convey("Then the string should be correct", func() {
				So(s, ShouldEqual, `date == date("0001-01-01T00:00:00Z")`)
			})
		})
	})

	Convey("Given I have a filter with duration", t, func() {

		f := NewFilterComposer().WithKey("duration").Equals(-2 * time.Second).Done()

		Convey("When I call String", func() {

			s := f.String()

			Convey("Then the string should be correct", func() {
				So(s, ShouldEqual, `duration == now("-2s")`)
			})
		})
	})
}

func TestFilter_SubFilters(t *testing.T) {

	Convey("Given I have a composed filters", t, func() {

		f := NewFilterComposer().
			WithKey("namespace").Equals("coucou").
			WithKey("number").Equals(32.9).
			And(
				NewFilterComposer().
					WithKey("name").Equals("toto").
					WithKey("value").Equals(1).
					Done(),
				NewFilterComposer().
					WithKey("color").Contains("red", "green", "blue", 43).
					WithKey("something").NotContains("stuff").
					Or(
						NewFilterComposer().
							WithKey("size").Matches(".*").
							Done(),
						NewFilterComposer().
							WithKey("size").Equals("medium").
							WithKey("fat").Equals(false).
							Done(),
						NewFilterComposer().
							WithKey("size").In(true, false).
							WithKey("size").NotIn(1).
							Done(),
					).
					Done(),
			).
			Done()

		Convey("When I call string it should be correct ", func() {
			So(f.String(), ShouldEqual, `namespace == "coucou" and number == 32.900000 and ((name == "toto" and value == 1) and (color contains ["red", "green", "blue", 43] and something not contains "stuff" or ((size matches [".*"]) or (size == "medium" and fat == false) or (size in [true, false] and size not in 1))))`)
		})
	})
}

func TestFilter_Exists(t *testing.T) {

	Convey("Given I create a filter", t, func() {

		f := NewFilterComposer().WithKey("key1").Exists().WithKey("key2").NotExists().Done()

		Convey("When I call string it should be correct ", func() {
			So(f.String(), ShouldEqual, "key1 exists and key2 not exists")
		})
	})
}
