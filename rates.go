package bcapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Rate contains information for coinFrom -> coinTo transaction
type Rate struct {
	// CoinFrom is coin from where to where we want
	// to convert the coin
	CoinFrom Coin `json:"coin_from"`

	// CoinTo this is the coin we want
	// to convert to
	CoinTo Coin `json:"coin_to"`

	// Price is a cost per 1 coinTo in coinTo currency
	//
	// for example we want to convert QIWI_RUB -> BTC
	// we will got price 4 448 698 QIWI_RUB per 1 BTC
	Price string `json:"price"`

	// PriceFrom is a price starting from how many we can purchase
	PriceFrom string `json:"price_from"`

	// PriceTill is a price till to we can purchase
	PriceTill string `json:"price_till"`

	// Market is a market
	Market Market `json:"market"`

	// Rating is a rating of market in format negative/positive
	//
	// for example "WW-Pay" has 0/8467 rating,
	// 0 - negative, 8467 postive
	Rating string `json:"rating"`

}

// JSON return the human-readable JSON format string
func (r *Rate) JSON() (string, error){
	type tmp struct{
		Rate         *Rate
		MarketLink   string
		FullListLink string
	}
	q := &tmp{Rate: r,MarketLink: r.GenerateLink(), FullListLink: r.GenerateLink()}
	data, err := json.Marshal(q)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// String return the human-readable string
func (r *Rate) String() string {
	return fmt.Sprintf("%s -> %s\nMarket - %s\nPrice - %s (can buy from %s - to %s)\nMarket link - %s\nLink to full list - %s\n", r.CoinFrom.ShortName, r.CoinTo.ShortName, r.Market.Name, r.Price, r.PriceFrom, r.PriceTill, r.GenerateMarketLink(), r.GenerateLink())
}

// GenerateLink generates link to bestchange site
//
// Unfortunately, works incorrectly sometimes
func (r *Rate) GenerateLink() string{
	alias1 := coinNames[r.CoinFrom.FullName]
	alias2 := coinNames[r.CoinTo.FullName]
	return fmt.Sprintf("%s%s-to-%s.html", exchLink, alias1, alias2)
}


// GenerateMarketLink generates the link to the market of this offer
func (r *Rate) GenerateMarketLink() string{
	return fmt.Sprintf("%sid=%s&from=%s&to=%s&city=0", marketLink, r.Market.ID, r.CoinFrom.ID, r.CoinTo.ID)
}


// getRates return the rates by a given string data
func getRates(data string, coins Coins, markets Markets) ([]Rate, error) {
	var res []Rate
	s := ""
	for _, v := range data {
		if v == '\n' {
			splt := strings.Split(s, ";")
			if len(splt) < 11 {
				return nil, fmt.Errorf("invalid data")
			}
			price := splt[4]
			if strings.Contains(splt[3], ".") {
				price = splt[3]
			}
			res = append(res, Rate{
				CoinFrom:  coins[splt[0]],
				CoinTo:    coins[splt[1]],
				Market:    markets[splt[2]],
				Price:     price,
				Rating:    splt[6],
				PriceFrom: splt[8],
				PriceTill: splt[9],
			})
			s = ""
			continue
		}
		s += string(v)
	}

	splt := strings.Split(s, ";")
	if len(splt) < 11 {
		return nil, fmt.Errorf("invalid data")
	}

	price := splt[4]
	if strings.Contains(splt[3], ".") {
		price = splt[3]
	}

	res = append(res, Rate{
		CoinFrom:  coins[splt[0]],
		CoinTo:    coins[splt[1]],
		Market:    markets[splt[2]],
		Price:     price,
		Rating:    splt[6],
		PriceFrom: splt[8],
		PriceTill: splt[9],
	})

	return res, nil
}


func newRates(coins Coins, markets Markets) ([]Rate, error) {
	data, err := openFile(ratesFileName)
	if err != nil {
		return nil, err
	}

	rates, err := getRates(data, coins, markets)
	if err != nil {
		return nil, err
	}

	return rates, nil
}
