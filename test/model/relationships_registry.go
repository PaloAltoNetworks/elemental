package testmodel

import "go.aporeto.io/elemental"

const nodocString = "[nodoc]" // nolint: varcheck,deadcode

var relationshipsRegistry elemental.RelationshipsRegistry

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[ListIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TaskIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[UserIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{}).Build(),
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{}).Build(),
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

}
