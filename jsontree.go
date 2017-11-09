package jsontree

import (
	"encoding/json"
	"fmt"
	"sort"
)

// Sort options.
const (
	SortNone = iota
	SortAsc
	SortDesc
)

// JSONTree interface.
type JSONTree interface {
	UnmarshalPrint([]byte) error
	Print(interface{})
}

// Options structure.
type Options struct {
	// Truncate characters (Optional).
	Truncate bool

	// Number of characters before they are truncated (Optional defaults to 40).
	NumChars int

	// Don't print values in tree (Optional).
	NoValues bool

	// Sorting either None, Asc or Desc (Optional).
	Sort int

	// Print summary
	Summary bool
}

type jsonTree struct {
	truncate bool
	numChars int
	noValues bool
	sort     int
	summary  bool

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
		noValues: o.NoValues,
		sort:     o.Sort,
		summary:  o.Summary,
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

	// For any other value assume SortNone
	switch t.sort {
	case SortAsc:
		sort.Strings(a)
	case SortDesc:
		sort.Strings(a)

		var b []string
		for i := len(a) - 1; i >= 0; i-- {
			b = append(b, a[i])
		}
		a = b
	}

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
		if !t.noValues {
			s := fmt.Sprintf("%v", v)

			if t.truncate && len(s) > t.numChars {
				s = s[0:t.numChars] + "..."
			}

			fmt.Printf(": %s", s)
		}
		t.nKeys++
	}
}

// Unmarshal JSON document and print tree.
func (t *jsonTree) UnmarshalPrint(b []byte) error {
	var d interface{}
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}

	t.Print(d)
	return nil
}

// Print tree.
func (t *jsonTree) Print(v interface{}) {
	fmt.Printf("/")
	t.traverse(v, "")
	if t.summary {
		fmt.Printf("\n\n%d objects, %d arrays, %d keys\n", t.nMaps, t.nArr, t.nKeys)
	} else {
		fmt.Printf("\n")
	}
}
