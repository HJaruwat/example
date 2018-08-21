package rabbitmq

import "testing"

func TestEmitMassageMigrate(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				body: []byte(`{"exe_id":"kerkoleng","est_id":"ker54933945","server_id":"18"}`),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				body: []byte(``),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EmitMassageMigrate(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("EmitMassageMigrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
