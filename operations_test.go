package tetra

import "testing"

func Test_operate(t *testing.T) {
	config := Config{
		Comma:           ',',
		FieldsPerRecord: -1,
	}
	type args struct {
		transform Transform
		csvData   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"None Transform",
			args{
				Transform{
					"none",
					map[string]interface{}{},
					&config,
				},
				"a,b,c,d,\r\n1,2,3,4,\r\nw,x,y,z,",
			},
			"a,b,c,d,\r\n1,2,3,4,\r\nw,x,y,z,",
			false,
		},
		{
			"Undefined Transform",
			args{
				Transform{
					"undef",
					map[string]interface{}{},
					&config,
				},
				"a,b,c,d,\r\n1,2,3,4,\r\nw,x,y,z,",
			},
			"",
			true,
		},
		{
			"Slice Rows Transform",
			args{
				Transform{
					"slice_rows",
					map[string]interface{}{
						"start": 1.0,
						"end":   -1.0,
					},
					&config,
				},
				"a,b,c,d,\n1,2,3,4,\nw,x,y,z,\na,b,c,d,\n1,2,3,4,\nw,x,y,z,\n",
			},
			"1,2,3,4,\nw,x,y,z,\na,b,c,d,\n1,2,3,4,\nw,x,y,z,\n",
			false,
		},
		{
			"TitleCase Transform",
			args{
				Transform{
					"titlecase_column",
					map[string]interface{}{
						"index": 1.0,
					},
					&config,
				},
				"hello world,michAEL o'connor,\n",
			},
			"hello world,Michael O'Connor,\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := operate(tt.args.transform, tt.args.csvData)
			if (err != nil) != tt.wantErr {
				t.Errorf("operate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("operate() = %v, want %v", got, tt.want)
			}
		})
	}
}
