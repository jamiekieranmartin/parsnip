package parsnip

import (
	"fmt"
	"regexp"
	"strings"
)

// Parse an input to key-value pairs given an expression
func Parse(exp, in string) (map[string]string, error) {
	// ensure regexp compiles
	re, err := regexp.Compile(exp)
	if err != nil {
		// extract helper text
		e := strings.TrimSpace(strings.Split(err.Error(), ":")[1])
		return nil, RegExpError{e}
	}

	// ensure input matches
	if ok := re.MatchString(in); !ok {
		return nil, ErrNoMatch
	}

	// extract keys and match
	keys := re.SubexpNames()
	matches := re.FindAllStringSubmatch(in, -1)[0]

	obj := map[string]string{}

	// iterate matched groups
	for i, match := range matches {
		// 0 is generally the full match
		if i == 0 {
			continue
		}

		key := keys[i]

		// use the index if key is empty
		if key == "" {
			key = fmt.Sprintf("%d", i)
		}

		// attach the match if it exists
		if match != "" {
			obj[key] = match
		}
	}

	return obj, nil
}
