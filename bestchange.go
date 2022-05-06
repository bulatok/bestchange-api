package bcapi

import (
	"fmt"
	"net/http"
)

// Bestchange is a main object of api that provides
// the way to get the best rates for coins
type Bestchange struct {
	Coins   Coins
	Markets Markets
	Rates   []Rate
}

// NewBestchange returns the Bestchange object with which we can
// find the best price for coins
//
// As argument you should pass your specified http.Client or just pass http.DefaultClient
func NewBestchange(client *http.Client) (*Bestchange, error) {
	// downloading the ZIP file
	if err := getZipFile(client); err != nil {
		return nil, err
	}

	// getting coins
	coins, err := newCoins()
	if err != nil {
		return nil, err
	}

	// getting markets
	markets, err := newMarkets()
	if err != nil {
		return nil, err
	}

	// getting rates
	rates, err := newRates(coins, markets)
	if err != nil {
		return nil, err
	}

	return &Bestchange{
		Coins:   coins,
		Markets: markets,
		Rates:   rates,
	}, nil
}

// GetRatesFromTo return rates
//
// You can pass to this function either the ID, full-name, short-name.
//
// But ID is preferable, because it will 100% work
func (b *Bestchange) GetRatesFromTo(from, to string) ([]Rate, error) {
	var res []Rate
	for _, v := range b.Rates {

		exp1 := v.CoinFrom.ShortName == from || v.CoinFrom.FullName == from || v.CoinFrom.ID == from
		exp2 := v.CoinTo.ShortName == to || v.CoinTo.FullName == to || v.CoinTo.ID == to

		if exp1 && exp2 {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("bcapi : %s-%s rate does not exits", from, to)
	}
	return res, nil
}

// UpdateCoins again download the zip file and updates exactly Coins
func (b *Bestchange) UpdateCoins(client *http.Client) error {
	if err := getZipFile(client); err != nil {
		return err
	}

	coins, err := newCoins()
	if err != nil {
		return err
	}

	b.Coins = coins
	return nil
}


// UpdateMarktes again download the zip file and updates exactly Markets
func (b *Bestchange) UpdateMarktes(client *http.Client) error {
	if err := getZipFile(client); err != nil {
		return err
	}

	markets, err := newMarkets()
	if err != nil {
		return err
	}

	b.Markets = markets
	return nil
}


// UpdateRates again download the zip file and updates exactly Rates
func (b *Bestchange) UpdateRates(client *http.Client) error {
	if err := getZipFile(client); err != nil {
		return err
	}

	rates, err := newRates(b.Coins, b.Markets)
	if err != nil {
		return err
	}

	b.Rates = rates
	return nil
}


