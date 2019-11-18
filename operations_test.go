package tetra

import "testing"

func Test_operate(t *testing.T) {
	var NoneTransform = Transform{
		Operation: "none",
		KWArgs:    make(map[string]string),
	}
	var UndefTransform = Transform{
		Operation: "undef",
		KWArgs:    make(map[string]string),
	}
	type args struct {
		transform Transform
		csvData   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Initial Test",
			args{
				NoneTransform,
				"one,two,three,four,\na,b,c,d,",
			},
			"one,two,three,four,\na,b,c,d,",
		},
		{
			"Test undefined operation",
			args{
				UndefTransform,
				"one,two,three,four,\na,b,c,d,",
			},
			"one,two,three,four,\na,b,c,d,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operate(tt.args.transform, tt.args.csvData); got != tt.want {
				t.Errorf("operate() = %v, want %v", got, tt.want)
			}
		})
	}
}
