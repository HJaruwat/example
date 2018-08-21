package rabbitmq

import (
	"cabal-api/common"
	"testing"
	"time"
)

func TestEmitMassageItem(t *testing.T) {
	code := common.RandomRefCode(10)

	var rewards []Rewards

	reward1 := Rewards{
		ItemID: "3588",
		Amount: 10,
		Name:   "Superioi Odd Clircle",
	}
	rewards = append(rewards, reward1)
	reward2 := Rewards{
		ItemID: "33555596",
		Amount: 2,
		Name:   "GM's Blessing (LV.1) Holy Water",
	}
	rewards = append(rewards, reward2)

	type args struct {
		body *InputItemMultiple
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
				body: &InputItemMultiple{
					CreatedAt:     time.Now(),
					ReferenceCode: code,
					ServerID:      "1",
					Username:      "kerkoleng",
					Rewards:       rewards,
				},
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				body: &InputItemMultiple{
					CreatedAt:     time.Now(),
					ReferenceCode: code,
					ServerID:      "1",
					Username:      "kerkoleng",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EmitMassageItem(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("EmitMassageItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
