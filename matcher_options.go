package elemental

type matchConfig struct{}

// MatcherOption represents the type for the options that can be passed to the helper `MatchesFilter` which can be used
// to alter the matching behaviour
type MatcherOption func(*matchConfig)
