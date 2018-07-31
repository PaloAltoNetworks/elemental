package testmodel

import "go.aporeto.io/elemental"

const nodocString = "[nodoc]" // nolint: varcheck,deadcode

var relationshipsRegistry elemental.RelationshipsRegistry

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[ListIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "rlcp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "rlcp2",
						Type: "boolean",
					},
				},
			}),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "lup1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "lup2",
						Type: "boolean",
					},
				},
			}),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "lup1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "lup2",
						Type: "boolean",
					},
				},
			}),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "ldp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "ldp2",
						Type: "boolean",
					},
				},
			}),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "lgp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "lgp2",
						Type: "boolean",
					},
				},
			}),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "rlgmp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "rlgmp2",
						Type: "boolean",
					},
				},
			}),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "rlgmp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "rlgmp2",
						Type: "boolean",
					},
				},
			}),
		},
	}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TaskIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "ltcp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "ltcp2",
						Type: "boolean",
					},
				},
			}),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "ltgp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "ltgp2",
						Type: "boolean",
					},
				},
			}),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "ltgp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "ltgp2",
						Type: "boolean",
					},
				},
			}),
		},
	}

	relationshipsRegistry[UserIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "rucp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "rucp2",
						Type: "boolean",
					},
				},
			}),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{}),
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "rugmp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "rugmp2",
						Type: "boolean",
					},
				},
			}),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"list": (&elemental.RelationshipInfo{}),
			"root": (&elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "rugmp1",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "rugmp2",
						Type: "boolean",
					},
				},
			}),
		},
	}

}
