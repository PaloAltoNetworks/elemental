// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRelationship_IsRetrieveAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows retrieve", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Retrieve: map[string]*RelationshipInfo{
					RootIdentity.Name: &RelationshipInfo{
						Parameters: []Parameter{
							Parameter{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsRetrieveAllowed", func() {

			ok := IsRetrieveAllowed(registry, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call RetrieveParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []Parameter{Parameter{Name: "toto"}})
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

		Convey("When I call RetrieveParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
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

		Convey("When I call RetrieveParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsUpdateAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows update", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Update: map[string]*RelationshipInfo{
					RootIdentity.Name: &RelationshipInfo{
						Parameters: []Parameter{
							Parameter{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsUpdateAllowed", func() {

			ok := IsUpdateAllowed(registry, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call UpdateParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []Parameter{Parameter{Name: "toto"}})
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

		Convey("When I call UpdateParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
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

		Convey("When I call UpdateParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsDeleteAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows delete", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Delete: map[string]*RelationshipInfo{
					RootIdentity.Name: &RelationshipInfo{
						Parameters: []Parameter{
							Parameter{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsDeleteAllowed", func() {

			ok := IsDeleteAllowed(registry, ListIdentity)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call DeleteParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []Parameter{Parameter{Name: "toto"}})
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

		Convey("When I call DeleteParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
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

		Convey("When I call DeleteParameters", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsRetrieveManyAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows retrieveMany", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				RetrieveMany: map[string]*RelationshipInfo{
					ListIdentity.Name: &RelationshipInfo{
						Parameters: []Parameter{
							Parameter{Name: "toto"},
						},
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

		Convey("When I call RetrieveManyParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []Parameter{Parameter{Name: "toto"}})
			})
		})
	})

	Convey("Given I have some relationships that don't allows retrieveMany", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				RetrieveMany: map[string]*RelationshipInfo{},
			},
		}

		Convey("When I call IsRetrieveManyAllowed", func() {

			ok := IsRetrieveManyAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call RetrieveManyParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
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

		Convey("When I call RetrieveManyParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some partial relationships", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsRetrieveManyAllowed", func() {

			ok := IsRetrieveManyAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call RetrieveManyParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsInfoAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows info", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Info: map[string]*RelationshipInfo{
					ListIdentity.Name: &RelationshipInfo{
						Parameters: []Parameter{
							Parameter{Name: "toto"},
						},
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

		Convey("When I call InfoParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []Parameter{Parameter{Name: "toto"}})
			})
		})
	})

	Convey("Given I have some relationships that don't allows info", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Info: map[string]*RelationshipInfo{},
			},
		}

		Convey("When I call IsInfoAllowed", func() {

			ok := IsInfoAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call InfoParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
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

		Convey("When I call InfoParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some partial relationships", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsInfoAllowed", func() {

			ok := IsInfoAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call InfoParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsPatchAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows patch", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Patch: map[string]*RelationshipInfo{
					ListIdentity.Name: &RelationshipInfo{
						Parameters: []Parameter{
							Parameter{Name: "toto"},
						},
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

		Convey("When I call PatchParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []Parameter{Parameter{Name: "toto"}})
			})
		})
	})

	Convey("Given I have some relationships that don't allows patch", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Patch: map[string]*RelationshipInfo{},
			},
		}

		Convey("When I call IsPatchAllowed", func() {

			ok := IsPatchAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call PatchParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
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

		Convey("When I call PatchParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some partial relationships", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsPatchAllowed", func() {

			ok := IsPatchAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call PatchParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsCreateAllowed(t *testing.T) {

	Convey("Given I have some relationships that allows create", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Create: map[string]*RelationshipInfo{
					ListIdentity.Name: &RelationshipInfo{
						Parameters: []Parameter{
							Parameter{Name: "toto"},
						},
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

		Convey("When I call CreateParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []Parameter{Parameter{Name: "toto"}})
			})
		})
	})

	Convey("Given I have some relationships that don't allows create", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Create: map[string]*RelationshipInfo{},
			},
		}

		Convey("When I call IsCreateAllowed", func() {

			ok := IsCreateAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call CreateParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
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

		Convey("When I call CreateParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some partial relationships", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsCreateAllowed", func() {

			ok := IsCreateAllowed(registry, TaskIdentity, ListIdentity)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call CreateParameters", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}
