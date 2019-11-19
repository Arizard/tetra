package tetra

import (
	"testing"
)

func TestTransformCSV(t *testing.T) {
	config := Config{
		Comma:           ',',
		FieldsPerRecord: -1,
	}
	config.AddTransform(
		"none",
		map[string]interface{}{},
	)
	config.AddTransform(
		"slice_rows",
		map[string]interface{}{
			"start": 1,
			"end":   -1,
		},
	)
	type args struct {
		cfg Config
		csv string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{
				config,
				"a,b,c,d,\n1,2,3,4,\ne,f,g,h,\n5,6,7,8,\n",
			},
			"1,2,3,4,\ne,f,g,h,\n5,6,7,8,\n",
		},
		{
			"Test variable fields config",
			args{
				config,
				"a,b,c,d,\n1,2,3,4,\ne,f,g,h,\n5,6,7,8,\n",
			},
			"1,2,3,4,\ne,f,g,h,\n5,6,7,8,\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformCSV(tt.args.cfg, tt.args.csv); got != tt.want {
				t.Errorf("TransformCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
