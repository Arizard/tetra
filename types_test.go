package tetra

import "testing"

var sampleJSON = `{
    "comma_string": ",",
    "transforms": [
        {
            "operation": "slice_rows",
            "kw_args": {
                "start": 1,
                "end": -1
            }
        }
    ]
}`

func TestConfig_LoadFromJSON(t *testing.T) {
	type fields struct {
		Comma            rune
		Comment          rune
		FieldsPerRecord  int
		LazyQuotes       bool
		TrimLeadingSpace bool
		ReuseRecord      bool
		Transforms       []Transform
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test unmarshal json",
			fields{},
			args{
				[]byte(sampleJSON),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Comma:            tt.fields.Comma,
				Comment:          tt.fields.Comment,
				FieldsPerRecord:  tt.fields.FieldsPerRecord,
				LazyQuotes:       tt.fields.LazyQuotes,
				TrimLeadingSpace: tt.fields.TrimLeadingSpace,
				ReuseRecord:      tt.fields.ReuseRecord,
				Transforms:       tt.fields.Transforms,
			}
			if err := c.LoadFromJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Config.LoadFromJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
