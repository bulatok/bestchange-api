package bcapi

import (
	"fmt"
	"strings"
)

type Market struct {
	ID   string `json:"market_id"`
	Name string `json:"market_name"`
}

// Markets is alias for map[string]string
type Markets map[string]Market

// getMarkets returns the Markets
// by a given string in format
// ID;Name;....
func getMarkets(data string) (Markets, error) {
	res := Markets{}
	s := ""
	for _, v := range data {
		if v == '\n' {
			splt := strings.Split(s, ";")
			if len(splt) < 2 {
				return nil, fmt.Errorf("invalid data")
			}
			res[splt[0]] = Market{splt[0], splt[1]}
			s = ""
			continue
		}
		s += string(v)
	}
	splt := strings.Split(s, ";")
	if len(splt) < 2 {
		return nil, fmt.Errorf("invalid data")
	}
	res[splt[0]] = Market{splt[0], splt[1]}
	return res, nil
}
func newMarkets() (Markets, error) {
	data, err := openFile(marketsFileName)
	if err != nil {
		return nil, err
	}
	markets, err := getMarkets(data)
	if err != nil {
		return nil, err
	}
	return markets, nil
}
