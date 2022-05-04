package parsnip_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jamiekieranmartin/parsnip"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		exp, in string
		wantErr error
		want    map[string]string
	}{
		// Empty
		{"", "", nil, map[string]string{}},
		// Successful
		{`(?P<first>\S+) (?P<middle>\S+) (?P<last>\S+)`, "Jamie Kieran Martin", nil, map[string]string{"first": "Jamie", "middle": "Kieran", "last": "Martin"}},
		// Compilation Error
		{`( NAME`, "", parsnip.RegExpError{"missing closing )"}, nil},
		// Parse Error
		{`^(?P<first>\S+) (?P<middle>\S+) (?P<last>\S+)$`, "this should not match", parsnip.ErrNoMatch, nil},
		// Unnamed Groups
		{`(?P<first>\S+) (\S+) (?P<last>\S+)`, "Jamie Kieran Martin", nil, map[string]string{"first": "Jamie", "2": "Kieran", "last": "Martin"}},
		// Optional, not included
		{`(\S+)\s(\S+)(?:\s(\S+))?`, "Jamie Martin", nil, map[string]string{"1": "Jamie", "2": "Martin"}},
		// Optional, included
		{`(\S+)\s(\S+)(?:\s(\S+))?`, "Jamie Kieran Martin", nil, map[string]string{"1": "Jamie", "2": "Kieran", "3": "Martin"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.exp, tt.in)

		t.Run(testname, func(t *testing.T) {
			out, err := parsnip.Parse(tt.exp, tt.in)

			if err != tt.wantErr {
				t.Errorf("got %s, want %s", err, tt.wantErr)
			}

			if !reflect.DeepEqual(out, tt.want) {
				t.Errorf("got %v, want %v", out, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parsnip.Parse(`(?P<first>\S+) (?P<middle>\S+) (?P<last>\S+)`, "Jamie Kieran Martin")
	}
}
