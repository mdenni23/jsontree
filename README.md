[![GoDoc](https://godoc.org/github.com/mickep76/jsontree?status.svg)](https://godoc.org/github.com/mickep76/jsontree)

# jsontree - JSON Tree

Go package for printing JSON as a Tree in the terminal.

**Example:**

```bash
/
└── 0
    ├── Architecture: amd64
    ├── Author:
    ├── Comment:
    ├── Config
    │   ├── AttachStderr: false
    │   ├── AttachStdin: false
...
```

# jsontree
    import "github.com/mickep76/jsontree"







## type JSONTree
``` go
type JSONTree interface {
    UnmarshalPrint([]byte) error
    Print(interface{})
}
```
JSONTree interface.









### func New
``` go
func New(o *Options) JSONTree
```
New constructor.




## type Options
``` go
type Options struct {
    // Truncate characters (Optional).
    Truncate bool

    // Number of characters before they are truncated (Optional defaults to 40).
    NumChars int

    // Don't print values in tree (Optional).
    NoValues bool
}
```
Options structure.

















- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
