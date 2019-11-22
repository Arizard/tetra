package tetra

import (
	"encoding/json"
)

// Transform describes a transformation on a csv file.
type Transform struct {
	Operation string                 `json:"operation,omitempty"`
	KWArgs    map[string]interface{} `json:"kw_args,omitempty"`
	Config    *Config                `json:"config,omitempty"`
}

// Config defines meta information and a list of transformations for the
// csv file.
type Config struct {
	// See https://golang.org/pkg/encoding/csv/#Reader.Read
	Comma            rune        `json:"comma,omitempty"`
	Comment          rune        `json:"comment,omitempty"`
	FieldsPerRecord  int         `json:"fields_per_record,omitempty"`
	LazyQuotes       bool        `json:"lazy_quotes,omitempty"`
	TrimLeadingSpace bool        `json:"trim_leading_space,omitempty"`
	ReuseRecord      bool        `json:"reuse_record,omitempty"`
	Transforms       []Transform `json:"transforms,omitempty"`
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

// UnmarshalJSON converts json config into a struct.
func (c *Config) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, c)
	return err
}
