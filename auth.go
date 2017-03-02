package elemental

// A ClaimsHolder is the interface of a structure that can hold
// Identity Claims (as in a JWT).
type ClaimsHolder interface {
	SetClaims([]string)
	GetClaims() []string
}

// A TokenHolder is the interface of a structure that can hold a token.
type TokenHolder interface {
	GetToken() string
}

// An SessionHolder is both a ClaimsHolder and a TokenHolder.
type SessionHolder interface {
	ClaimsHolder
	TokenHolder
}
