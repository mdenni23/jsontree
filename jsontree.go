package jsontree

import (
	"fmt"
	"sort"
)

// JSONTree interface.
type JSONTree interface {
	Print(interface{})
}

// Options structure.
type Options struct {
	// Truncate characters.
	Truncate bool

	// Number of characters before they are truncated.
	NumChars int
}

type jsonTree struct {
	truncate bool
	numChars int

	nMaps int
	nArr  int
	nKeys int
}

// New constructor.
func New(o *Options) JSONTree {
	if o.NumChars == 0 {
		o.NumChars = 40
	}

	return &jsonTree{
		truncate: o.Truncate,
		numChars: o.NumChars,
	}
}

func (t *jsonTree) traverseArray(in []interface{}, indent string) {
	for i, v := range in {
		if i == len(in)-1 {
			fmt.Printf("\n%s└── %d", indent, i)
			t.traverse(v, indent+"    ")
		} else {
			fmt.Printf("\n%s├── %d", indent, i)
			t.traverse(v, indent+"|   ")
		}
	}
}

func (t *jsonTree) traverseMapStr(in map[string]interface{}, indent string) {
	var a []string

	for k := range in {
		a = append(a, k)
	}
	sort.Strings(a)

	for i, k := range a {
		if i == len(in)-1 {
			fmt.Printf("\n%s└── %s", indent, k)
			t.traverse(in[k], indent+"    ")
		} else {
			fmt.Printf("\n%s├── %s", indent, k)
			t.traverse(in[k], indent+"│   ")
		}
	}
}

func (t *jsonTree) traverse(v interface{}, indent string) {
	switch v := v.(type) {
	case []interface{}:
		t.traverseArray(v, indent)
		t.nMaps++
	case map[string]interface{}:
		t.traverseMapStr(v, indent)
		t.nArr++
	default:
		s := fmt.Sprintf("%v", v)

		if t.truncate && len(s) > t.numChars {
			s = s[0:t.numChars] + "..."
		}

		fmt.Printf(": %s", s)
		t.nKeys++
	}
}

// Include Unmarshal

func (t *jsonTree) Print(v interface{}) {
	fmt.Printf("/")
	t.traverse(v, "")
	fmt.Printf("\n\n%d objects, %d arrays, %d keys\n", t.nMaps, t.nArr, t.nKeys)
}
