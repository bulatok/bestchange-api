package bcapi

import (
	"archive/zip"
	"errors"
	"strings"
)

var (
	ErrMarketsParsing = errors.New("markets parsing")
)

// Market holds the ID and its Name
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
				return nil, wrapErrors(ErrMarketsParsing.Error())
			}
			res[splt[0]] = Market{splt[0], splt[1]}
			s = ""
			continue
		}
		s += string(v)
	}
	splt := strings.Split(s, ";")
	if len(splt) < 2 {
		return nil, wrapErrors(ErrMarketsParsing.Error())
	}
	res[splt[0]] = Market{splt[0], splt[1]}
	return res, nil
}
func newMarkets(zipArchive *zip.ReadCloser) (Markets, error) {
	data, err := openFile(zipArchive, marketsFileName)
	if err != nil {
		return nil, wrapErrors(ErrMarketsParsing.Error(), err.Error())
	}

	markets, err := getMarkets(data)
	if err != nil {
		return nil, wrapErrors(ErrMarketsParsing.Error(), err.Error())
	}
	return markets, nil
}
