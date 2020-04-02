package testmodel

import "go.aporeto.io/elemental"

var relationshipsRegistry elemental.RelationshipsRegistry

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[ListIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "rlcp1",
						Type: "string",
					},
					{
						Name: "rlcp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "lup1",
						Type: "string",
					},
					{
						Name: "lup2",
						Type: "boolean",
					},
				},
			},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "lup1",
						Type: "string",
					},
					{
						Name: "lup2",
						Type: "boolean",
					},
				},
			},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "ldp1",
						Type: "string",
					},
					{
						Name: "ldp2",
						Type: "boolean",
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "lgp1",
						Type: "string",
					},
					{
						Name: "lgp2",
						Type: "boolean",
					},
					{
						Name: "sAp1",
						Type: "string",
					},
					{
						Name: "sAp2",
						Type: "boolean",
					},
					{
						Name: "sBp1",
						Type: "string",
					},
					{
						Name: "sBp2",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "rlgmp1",
						Type: "string",
					},
					{
						Name: "rlgmp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "rlgmp1",
						Type: "string",
					},
					{
						Name: "rlgmp2",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TaskIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"list": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "ltcp1",
						Type: "string",
					},
					{
						Name: "ltcp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"list": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "ltgp1",
						Type: "string",
					},
					{
						Name: "ltgp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"list": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "ltgp1",
						Type: "string",
					},
					{
						Name: "ltgp2",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[UserIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "rucp1",
						Type: "string",
					},
					{
						Name: "rucp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"confirm",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "confirm",
						Type: "boolean",
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"list": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "rugmp1",
						Type: "string",
					},
					{
						Name: "rugmp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"list": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "rugmp1",
						Type: "string",
					},
					{
						Name: "rugmp2",
						Type: "boolean",
					},
				},
			},
		},
	}

}
