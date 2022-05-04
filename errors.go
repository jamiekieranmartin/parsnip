package parsnip

// ParseError stores information about an internal parsing error.
type ParseError struct {
	Key string // the parsing error.
}

func (e ParseError) Error() string {
	return e.Key
}

var ErrNoMatch = ParseError{"no match found"}

// RegExpError stores information about a regexp compilation error.
type RegExpError struct {
	Key string // the regexp error.
}

func (e RegExpError) Error() string {
	return e.Key
}
