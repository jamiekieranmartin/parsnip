package parsnip

import (
	"fmt"
	"regexp"
	"strings"
)

// Parse an input to key-value pairs given an expression
func Parse(exp, in string) (map[string]string, error) {
	re, err := regexp.Compile(exp)
	if err != nil {
		e := strings.TrimSpace(strings.Split(err.Error(), ":")[1])
		return nil, RegExpError{e}
	}

	if ok := re.MatchString(in); !ok {
		return nil, ErrNoMatch
	}

	keys := re.SubexpNames()
	matches := re.FindAllStringSubmatch(in, -1)[0]

	obj := map[string]string{}

	for i, match := range matches {
		if i == 0 {
			continue
		}

		key := keys[i]

		if key == "" {
			key = fmt.Sprintf("%d", i)
		}

		if match != "" {
			obj[key] = match
		}
	}

	return obj, nil
}
