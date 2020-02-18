package elemental

import "strings"

type filterParserConfig struct {
	// map to support an O(1) lookup during parsing
	unsupportedComparators map[parserToken]struct{}
}

// FilterParserOption represents the type for the options that can be passed to `NewFilterParser` which can be used to
// alter parsing behaviour
type FilterParserOption func(*filterParserConfig)

// OptUnsupportedComparators accepts a slice of comparators that will limit the set of comparators that the parser will accept.
// If supplied, the parser will return an error if the filter being parsed contains a comparator provided in the blacklist.
func OptUnsupportedComparators(blacklist []FilterComparator) FilterParserOption {
	return func(config *filterParserConfig) {
		config.unsupportedComparators = map[parserToken]struct{}{}
		for _, c := range blacklist {
			if token, ok := operatorsToToken[strings.ToUpper(translateComparator(c))]; ok {
				config.unsupportedComparators[token] = struct{}{}
			}
		}
	}
}
