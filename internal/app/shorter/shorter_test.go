package shorter

import "testing"

func TestShorting(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			want: "http://localhost:8080/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shorting(tt.args.id); got != tt.want {
				t.Errorf("Shorting() = %v, want %v", got, tt.want)
			}
		})
	}
}
