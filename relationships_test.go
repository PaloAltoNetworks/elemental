// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRelationship_AddChild(t *testing.T) {

	Convey("Given I have two Relationship", t, func() {
		r1 := &Relationship{}
		r2 := &Relationship{}

		Convey("When I add r2 as child of r1", func() {

			r1.AddChild(ListIdentity, r2)

			Convey("Then r2 should be added to the children of r1", func() {
				So(r1.Children[ListIdentity], ShouldEqual, r2)
			})
		})
	})
}

func TestRelationship_IsRetrieveAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows retrieve", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				AllowsRetrieve: true,
			},
		}

		Convey("When I call IsRetrieveAllowed", func() {

			ok := IsRetrieveAllowed(registry, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have some relationships that don't allows retrieve", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsRetrieveAllowed", func() {

			ok := IsRetrieveAllowed(registry, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsRetrieveAllowed", func() {

			ok := IsRetrieveAllowed(registry, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestRelationship_IsUpdateAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows update", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				AllowsUpdate: true,
			},
		}

		Convey("When I call IsUpdateAllowed", func() {

			ok := IsUpdateAllowed(registry, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have some relationships that don't allows update", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsUpdateAllowed", func() {

			ok := IsUpdateAllowed(registry, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsUpdateAllowed", func() {

			ok := IsUpdateAllowed(registry, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestRelationship_IsDeleteAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows delete", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				AllowsDelete: true,
			},
		}

		Convey("When I call IsDeleteAllowed", func() {

			ok := IsDeleteAllowed(registry, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have some relationships that don't allows delete", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsDeleteAllowed", func() {

			ok := IsDeleteAllowed(registry, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsDeleteAllowed", func() {

			ok := IsDeleteAllowed(registry, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestRelationship_IsRetrieveManyAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows retrieveMany", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{
					TaskIdentity: &Relationship{
						AllowsRetrieveMany: true,
					},
				},
			},
		}

		Convey("When I call IsRetrieveManyAllowed", func() {

			ok := IsRetrieveManyAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have some relationships that don't allows retrieveMany", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{
					TaskIdentity: &Relationship{},
				},
			},
		}

		Convey("When I call IsRetrieveManyAllowed", func() {

			ok := IsRetrieveManyAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsRetrieveManyAllowed", func() {

			ok := IsRetrieveManyAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some partial relationships", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{},
			},
		}

		Convey("When I call IsRetrieveManyAllowed", func() {

			ok := IsRetrieveManyAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestRelationship_IsInfoAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows info", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{
					TaskIdentity: &Relationship{
						AllowsInfo: true,
					},
				},
			},
		}

		Convey("When I call IsInfoAllowed", func() {

			ok := IsInfoAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have some relationships that don't allows info", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{
					TaskIdentity: &Relationship{},
				},
			},
		}

		Convey("When I call IsInfoAllowed", func() {

			ok := IsInfoAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsInfoAllowed", func() {

			ok := IsInfoAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some partial relationships", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{},
			},
		}

		Convey("When I call IsInfoAllowed", func() {

			ok := IsInfoAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestRelationship_IsPatchAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows patch", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{
					TaskIdentity: &Relationship{
						AllowsPatch: true,
					},
				},
			},
		}

		Convey("When I call IsPatchAllowed", func() {

			ok := IsPatchAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have some relationships that don't allows patch", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{
					TaskIdentity: &Relationship{},
				},
			},
		}

		Convey("When I call IsPatchAllowed", func() {

			ok := IsPatchAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsPatchAllowed", func() {

			ok := IsPatchAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some partial relationships", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{},
			},
		}

		Convey("When I call IsPatchAllowed", func() {

			ok := IsPatchAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestRelationship_IsCreateAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows create", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{
					TaskIdentity: &Relationship{
						AllowsCreate: true,
					},
				},
			},
		}

		Convey("When I call IsCreateAllowed", func() {

			ok := IsCreateAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have some relationships that don't allows create", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{
					TaskIdentity: &Relationship{},
				},
			},
		}

		Convey("When I call IsCreateAllowed", func() {

			ok := IsCreateAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsCreateAllowed", func() {

			ok := IsCreateAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have some partial relationships", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Children: RelationshipsRegistry{},
			},
		}

		Convey("When I call IsCreateAllowed", func() {

			ok := IsCreateAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}
