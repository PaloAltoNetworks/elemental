package testmodel

import "github.com/aporeto-inc/elemental"

const nodocString = "[nodoc]" // nolint: varcheck,deadcode

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[ListIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"root": true,
		},
		AllowsInfo: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TaskIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"list": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"list": true,
			"root": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"list": true,
			"root": true,
		},
		AllowsInfo: map[string]bool{
			"list": true,
			"root": true,
		},
	}

	relationshipsRegistry[UserIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"list": true,
			"root": true,
		},
		AllowsPatch: map[string]bool{
			"list": true,
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"list": true,
			"root": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"list": true,
			"root": true,
		},
		AllowsInfo: map[string]bool{
			"list": true,
			"root": true,
		},
	}

}
