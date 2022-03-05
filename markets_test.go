package bcapi

import (
	"reflect"
	"testing"
)

func Test_getMarkets(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want Markets
	}{
		{
			"valid data",
			args{
				`56;M-Obmen;963130820766;3162;0
443;ОбменОМ;;0;287126
80;Vmex;060994506852;1673;363909
90;Ex-Money;;0;193271062`,
			},
			Markets{
				"56":  Market{"56", "M-Obmen"},
				"443": Market{"443", "ОбменОМ"},
				"80":  Market{"80", "Vmex"},
				"90":  Market{"90", "Ex-Money"},
			},
		},
		{
			"invalid data",
			args{
				`qweqqweeeq
12;qwe12eqwe;qwe`,
			},
			nil,
		},
		{
			"empty data",
			args{
				``,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := getMarkets(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMarkets() = %v, want %v", got, tt.want)
			}
		})
	}
}
