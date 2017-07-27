package elemental

// MockIdentifiable mocks a Identifiable
type MockIdentifiable struct {
	DefinedIdentity         string
	DefinedIdentifier       string
	ExpectedValidationError error
}

// Identity returns the Identity of the of the receiver.
func (p *MockIdentifiable) Identity() Identity {
	return MakeIdentity(p.DefinedIdentity, "FakeCategory")
}

// Identifier returns the unique identifier of the of the receiver.
func (p *MockIdentifiable) Identifier() string {
	return p.DefinedIdentifier
}

// SetIdentifier sets the unique identifier of the of the receiver.
func (p *MockIdentifiable) SetIdentifier(DefinedIdentifier string) {
	p.DefinedIdentifier = DefinedIdentifier
}

// Version returns the version number
func (p *MockIdentifiable) Version() int {
	return 1
}

// Validate is the method that verifies the object is valid
func (p *MockIdentifiable) Validate() error {
	return p.ExpectedValidationError
}
