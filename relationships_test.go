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

	. "github.com/smartystreets/goconvey/convey"
)

func TestRelationship_IsOperationAllowed_Retrieve(t *testing.T) {

	Convey("Given I have some relationships that allows retrieve", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Retrieve: map[string]*RelationshipInfo{
					RootIdentity.Name: {
						Parameters: []ParameterDefinition{
							{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []ParameterDefinition{{Name: "toto"}})
			})
		})
	})

	Convey("Given I have some relationships that don't allows retrieve", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationRetrieve)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsOperationAllowed_Update(t *testing.T) {

	Convey("Given I have some relationships that allows update", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Update: map[string]*RelationshipInfo{
					RootIdentity.Name: {
						Parameters: []ParameterDefinition{
							{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []ParameterDefinition{{Name: "toto"}})
			})
		})
	})

	Convey("Given I have some relationships that don't allows update", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationUpdate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsOperationAllowed_Delete(t *testing.T) {

	Convey("Given I have some relationships that allows delete", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{
				Delete: map[string]*RelationshipInfo{
					RootIdentity.Name: {
						Parameters: []ParameterDefinition{
							{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []ParameterDefinition{{Name: "toto"}})
			})
		})
	})

	Convey("Given I have some relationships that don't allows delete", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, OperationDelete)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsOperationAllowed_RetrieveMany(t *testing.T) {

	Convey("Given I have some relationships that allows retrieveMany", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				RetrieveMany: map[string]*RelationshipInfo{
					ListIdentity.Name: {
						Parameters: []ParameterDefinition{
							{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []ParameterDefinition{{Name: "toto"}})
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

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

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

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationRetrieveMany)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsOperationAllowed_Info(t *testing.T) {

	Convey("Given I have some relationships that allows info", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Info: map[string]*RelationshipInfo{
					ListIdentity.Name: {
						Parameters: []ParameterDefinition{
							{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []ParameterDefinition{{Name: "toto"}})
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

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

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

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationInfo)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsOperationAllowed_Patch(t *testing.T) {

	Convey("Given I have some relationships that allows patch", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Patch: map[string]*RelationshipInfo{
					ListIdentity.Name: {
						Parameters: []ParameterDefinition{
							{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []ParameterDefinition{{Name: "toto"}})
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

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

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

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationPatch)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsOperationAllowed_Create(t *testing.T) {

	Convey("Given I have some relationships that allows create", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Create: map[string]*RelationshipInfo{
					ListIdentity.Name: {
						Parameters: []ParameterDefinition{
							{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then retrieve should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldResemble, []ParameterDefinition{{Name: "toto"}})
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

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})

	Convey("Given I have some empty relationships", t, func() {

		registry := RelationshipsRegistry{}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

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

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then retrieve should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, TaskIdentity, ListIdentity, OperationCreate)

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestRelationship_IsOperationAllowed_UnknownOperation(t *testing.T) {

	Convey("Given I have some relationships that allows create", t, func() {

		registry := RelationshipsRegistry{
			ListIdentity: &Relationship{},
			TaskIdentity: &Relationship{
				Create: map[string]*RelationshipInfo{
					ListIdentity.Name: {
						Parameters: []ParameterDefinition{
							{Name: "toto"},
						},
					},
				},
			},
		}

		Convey("When I call IsOperationAllowed", func() {

			ok := IsOperationAllowed(registry, TaskIdentity, ListIdentity, Operation("unknown"))

			Convey("Then retrieve should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call ParametersForOperation", func() {

			p := ParametersForOperation(registry, ListIdentity, RootIdentity, Operation("unknown"))

			Convey("Then parameters should be correct", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}
