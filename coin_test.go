package bcapi

import (
	"reflect"
	"testing"
)

func Test_getCoins(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    Coins
		wantErr bool
	}{
		{
			"valid",
			args{`201;51;Polkadot (DOT);DOT;1008;0;10111111111111111111111111111111111111111110111111111111111111000000001111111111111111111110001010000000000010101000000100001101111111111111111111111111111111111101011100010001000001111111111000000000111100100000000000011100001
202;52;Uniswap (UNI);UNI;1009;0;10111111111111111111111111111111111111111110111111111111111111000000001111111111111111111110001010000000000010101000000100000101111111111111111111101111111111111101011100010001000001111111111000000000111100100000000000011100001`,
			},
			Coins{
				"201": Coin{"201", "Polkadot (DOT)", "DOT"},
				"202": Coin{"202", "Uniswap (UNI)", "UNI"},
			},
			false,
		},
		{
			"invalid",
			args{`58CARDUSD
59;CARDRUB
60;CARDUAH
46;TCSBCRUB
61;BAT`,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCoins(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCoins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCoins() got = %v, want %v", got, tt.want)
			}
		})
	}
}
