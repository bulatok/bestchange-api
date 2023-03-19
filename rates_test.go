package bcapi

import (
	"archive/zip"
	"net/http"
	"reflect"
	"testing"
)

// Test_getRates it is not required to pass this test
func Test_getRates(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    []Rate
		wantErr bool
	}{
		{
			name: "OK. valid",
			args: args{
				data: `176;10;840;1;118.44316327;5150000;0.2898;1;10000;49000;0`,
			},
			want: []Rate{
				{Coin{"176", "Открытие RUB", "RUB Открытие"},
					Coin{"10", "Tether TRC20 (USDT)", "USDT TRC20"},
					"1", "10000", "49000", Market{"840", "YoChange"},
					"0.2898", "118.44316327"},
			},
			wantErr: false,
		},
		{
			name: "Invalid data",
			args: args{
				data: `12:212:123
123:123;132`},
			want:    nil,
			wantErr: true,
		},
	}

	if *needDownloadZip && !downloaded.Load() {
		downloaded.Store(true)
		if err := downloadZipArchive(http.DefaultClient); err != nil {
			t.Fatalf(err.Error())
		}
	}

	zipOpened, err := zip.OpenReader(zipFileName)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer zipOpened.Close()

	coins, _ := newCoins(zipOpened)
	markets, _ := newMarkets(zipOpened)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRates(tt.args.data, coins, markets)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}
