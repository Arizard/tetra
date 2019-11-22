package tetra

import (
	"encoding/json"

	"github.com/golang/glog"
)

// Transform describes a transformation on a csv file.
type Transform struct {
	Operation string                 `json:"operation,omitempty"`
	KWArgs    map[string]interface{} `json:"kw_args,omitempty"`
	Config    *Config
}

// Config defines meta information and a list of transformations for the
// csv file.
type Config struct {
	// See https://golang.org/pkg/encoding/csv/#Reader.Read
	commaString      string `json:"comma_string,string,omitempty"`
	Comma            rune
	commentString    string `json:"comment_string,string,omitempty"`
	Comment          rune
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

// LoadFromJSON converts json config into a struct.
func (c *Config) LoadFromJSON(b []byte) error {
	err := json.Unmarshal(b, c)
	for _, tran := range c.Transforms {
		tran.Config = c
	}
	glog.Infof("%+v", c)
	deRef := *c
	c.Comma = []rune(deRef.commaString)[0]
	c.Comment = []rune(deRef.commentString)[0]
	return err
}
