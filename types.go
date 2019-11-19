package tetra

// Transform describes a transformation on a csv file.
type Transform struct {
	Operation string
	KWArgs    map[string]interface{}
	Config    *Config
}

// Config defines meta information and a list of transformations for the
// csv file.
type Config struct {
	// See https://golang.org/pkg/encoding/csv/#Reader.Read
	Comma            rune
	Comment          rune
	FieldsPerRecord  int
	LazyQuotes       bool
	TrimLeadingSpace bool
	ReuseRecord      bool
	Transforms       []Transform
}

// AddTransform is a method of *Config which adds a new transform, making sure
// to include the reference back to the original config.
func (c *Config) AddTransform(op string, kwargs map[string]interface{}) {
	newTransform := Transform{
		Operation: op,
		KWArgs:    kwargs,
		Config:    c,
	}

	c.Transforms = append(c.Transforms, newTransform)
}
