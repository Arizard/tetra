package tetra

// Transform describes a transformation on a csv file.
type Transform struct {
	Operation string
	KWArgs    map[string]interface{}
}

// Config defines meta information and a list of transformations for the
// csv file.
type Config struct {
	// See https://golang.org/pkg/encoding/csv/#Reader.Read
	Delimiter        rune
	Comment          rune
	FieldsPerRecord  int
	LazyQuotes       bool
	TrimLeadingSpace bool
	ReuseRecord      bool
	Transforms       []Transform
}
