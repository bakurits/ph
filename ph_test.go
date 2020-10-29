package ph

import "testing"

func TestHashAndSalt(t *testing.T) {
	type args struct {
		realP  string
		inputP string
	}
	tests := []struct {
		name      string
		args      args
		wantMatch bool
		wantErr   bool
	}{
		{
			name: "match test",
			args: args{
				realP:  "giorgi121",
				inputP: "giorgi121",
			},
			wantMatch: true,
			wantErr:   false,
		},
		{
			name: "not match test",
			args: args{
				realP:  "giorgi121",
				inputP: "giorgi12",
			},
			wantMatch: false,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashAndSalt(tt.args.realP)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashAndSalt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if Compare(got, tt.args.inputP) != tt.wantMatch {
				t.Errorf("Unexpected result on plain text %v", tt.args.inputP)
				return
			}
		})
	}
}
