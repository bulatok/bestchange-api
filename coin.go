package main

import (
	"fmt"
	"strings"
)
var(
	coinNames = map[string]string{
		"WMZ":                  "wmz",
		"WMR":                  "wmr",
		"WME":                  "wme",
		"Беларусбанк":          "belarusbank",
		"Приват 24 USD":        "privat24-usd",
		"ЮMoney":               "yoomoney",
		"WU RUB":               "wu-rub",
		"Сбербанк":             "sberbank",
		"OMG Network (OMG)":    "omg",
		"WU EUR":               "wu",
		"WMB":                  "wmb",
		"Binance Coin (BNB)":   "binance-coin",
		"Advanced Cash TRY":    "advanced-cash-try",
		"Bitcoin BEP20 (BTC)":  "bitcoin",
		"ВТБ":                  "vtb",
		"WMK":                  "wmk",
		"Perfect Money USD":    "perfectmoney-usd",
		"Perfect Money EUR":    "perfectmoney-eur",
		"PayPal USD":           "paypal-usd",
		"Exmo USDT":            "exmo",
		"Альфа-Банк":           "alfabank-uah",
		"Промсвязьбанк":        "rosbank",
		"Счет телефона RUB":    "settlement-rub",
		"Visa/MasterCard BYN":  "visa-mastercard-byr",
		"Приват 24 UAH":        "privat24-uah",
		"Skrill USD":           "skrill",
		"Visa/MasterCard USD":  "visa-mastercard-usd",
		"Visa/MasterCard RUB":  "visa-mastercard-rub",
		"Visa/MasterCard UAH":  "visa-mastercard-uah",
		"Тинькофф cash-in":     "tinkoff-cash-in",
		"BAT (BAT)":            "bat",
		"Альфа cash-in RUB":    "alfabank-cash-in",
		"QIWI RUB":             "qiwi",
		"Русский Стандарт":     "russtandart",
		"Visa/MasterCard EUR":  "visa-mastercard-euro",
		"Кукуруза":             "kykyryza",
		"Kaspi Bank":           "kaspi-bank",
		"WU USD":               "wire-usd",
		"Ощадбанк":             "oschadbank",
		"Wrapped BTC (WBTC)":   "wrapped-bitcoin",
		"Любой банк RUB":       "wire-rub",
		"Neteller USD":         "neteller",
		"Paxum":                "paxum",
		"Idram":                "idram",
		"ForteBank":            "fortebank",
		"WMP":                  "wmp",
		"MoneyGram EUR":        "moneygram",
		"MoneyGram USD":        "moneygram-euro",
		"Авангард":             "avangard",
		"PayPal EUR":           "paypal-euro",
		"Любой банк GBP":       "wire-gbp",
		"Solana (SOL)":         "solana",
		"Visa/MasterCard TRY":  "visa-mastercard-try",
		"Монобанк":             "monobank",
		"Capitalist RUB":       "capitalist-rub",
		"UNI USD":              "uni",
		"Paymer USD":           "paymer",
		"Advanced Cash USD":    "advanced-cash-uah",
		"HalykBank":            "halykbank",
		"Наличные RUB":         "binance-rub",
		"Наличные UAH":         "binance-uah",
		"Bitcoin (BTC)":        "bitcoin",
		"Газпромбанк":          "gazprombank",
		"WMX":                  "wmx",
		"Epay EUR":             "epay",
		"PayPal RUB":           "paypal-rub",
		"Litecoin (LTC)":       "litecoin",
		"Карта UnionPay":       "unionpay",
		"Contact USD":          "contact-usd",
		"Любой банк UAH":       "wire-uah",
		"Payoneer":             "payoneer",
		"ICON (ICX)":           "icon",
		"Тинькофф":             "tinkoff",
		"Contact RUB":          "contact",
		"ЗК RUB":               "uni-rub",
		"Payeer USD":           "payeer",
		"NixMoney USD":         "nixmoney",
		"Binance UAH":          "sberbank-uah",
		"Visa/MasterCard KZT":  "visa-mastercard-kzt",
		"Global24":             "global24",
		"Любой банк KZT":       "wire-kzt",
		"Сбербанк KZT":         "sberbank-kzt",
		"Dogecoin (DOGE)":      "dogecoin",
		"ЗК USD":               "ria-usd",
		"Payeer RUB":           "payeer-rub",
		"ПУМБ":                 "pumb",
		"Любой банк INR":       "wire-inr",
		"Advanced Cash EUR":    "advanced-cash-euro",
		"Advanced Cash RUB":    "advanced-cash-rub",
		"Payeer EUR":           "payeer-euro",
		"Skrill EUR":           "skrill-euro",
		"Verge (XVG)":          "verge",
		"NixMoney EUR":         "nixmoney-euro",
		"QIWI KZT":             "qiwi-kzt",
		"Perfect Money BTC":    "perfectmoney-btc",
		"Exmo USD":             "exmo-uah",
		"Exmo RUB":             "exmo-rub",
		"Bitcoin LN (BTC)":     "bitcoin-ln",
		"РНКБ":                 "rnkb",
		"Waves (WAVES)":        "waves",
		"Komodo (KMD)":         "komodo",
		"Ontology (ONT)":       "ontology",
		"Neteller EUR":         "neteller-euro",
		"Bitcoin SV (BSV)":     "bitcoin-sv",
		"Polygon (MATIC)":      "polygon",
		"Ethereum (ETH)":       "ethereum",
		"Dash (DASH)":          "dash",
		"Advanced Cash UAH":    "advanced-cash",
		"Альфа cash-in USD":    "alfabank-cashin-usd",
		"Paymer RUB":           "paymer-rub",
		"Capitalist USD":       "capitalist",
		"Visa/MasterCard GBP":  "visa-mastercard-gbp",
		"Тинькофф QR":          "tinkoff-qr-codes",
		"Monero (XMR)":         "monero",
		"Ria USD":              "trade-usd",
		"Ria EUR":              "ria-euro",
		"PaySera USD":          "paysera",
		"Epay USD":             "epay-euro",
		"Visa/MasterCard SEK":  "visa-mastercard-sek",
		"PM e-Voucher USD":     "pm-voucher",
		"Райффайзен UAH":       "raiffeisen-bank-uah",
		"Ether Classic (ETC)":  "ethereum-classic",
		"Ripple (XRP)":         "ripple",
		"Zcash (ZEC)":          "zcash",
		"Tether Omni (USDT)":   "tether",
		"PayPal GBP":           "paypal-gbp",
		"Alipay":               "alipay",
		"Любой банк CNY":       "wire-cny",
		"Любой банк THB":       "wire-thb",
		"0x (ZRX)":             "zrx",
		"Exmo UAH":             "exmo-tether",
		"Почта Банк":           "post-bank",
		"Sepa EUR":             "sepa",
		"Bitcoin Cash (BCH)":   "bitcoin-cash",
		"NEM (XEM)":            "nem",
		"Augur (REP)":          "augur",
		"Tezos (XTZ)":          "tezos",
		"NEO (NEO)":            "neo",
		"EOS (EOS)":            "eos",
		"IOTA (MIOTA)":         "iota",
		"Lisk (LSK)":           "lisk",
		"Cardano (ADA)":        "cardano",
		"Stellar (XLM)":        "stellar",
		"Любой банк IDR":       "wire-idr",
		"Bitcoin Gold (BTG)":   "bitcoin-gold",
		"TRON (TRX)":           "tron",
		"Exmo BTC":             "exmo-btc",
		"ЕРИП Расчет":          "erip",
		"Visa/MasterCard AMD":  "visa-mastercard-amd",
		"VeChain (VET)":        "vechain",
		"Tether TRC20 (USDT)":  "tether-trc20",
		"Любой банк NGN":       "wire-ngn",
		"Garantex":             "garantex",
		"СБП":                  "sbp",
		"УкрСиббанк":           "ukrsibbank",
		"USD Coin (USDC)":      "usd-coin",
		"TrueUSD (TUSD)":       "trueusd",
		"Visa/MasterCard KGS":  "visa-mastercard-kgs",
		"Qtum (QTUM)":          "qtum",
		"BitTorrent (BTT)":     "bittorrent",
		"Kuna":                 "kuna",
		"Любой банк PLN":       "wire-pln",
		"Visa/MasterCard CNY":  "visa-mastercard-cny",
		"Любой банк BYN":       "wire-byn",
		"VelesPay":             "velespay",
		"Pax Dollar (USDP)":    "pax-dollar",
		"Cryptex":              "cryptex",
		"Advanced Cash KZT":    "advanced-cash-kzt",
		"Россельхозбанк":       "rosselhozbank",
		"PaySera EUR":          "paysera-euro",
		"Tether ERC20 (USDT)":  "tether-erc20",
		"Visa/MasterCard PLN":  "visa-mastercard-pln",
		"Любой банк TRY":       "wire-try",
		"МТС Банк":             "mts-bank",
		"Revolut USD":          "revolut-usd",
		"Revolut EUR":          "revolut-euro",
		"Visa/MasterCard MDL":  "visa-mastercard-mdl",
		"Росбанк":              "psbank",
		"Сбербанк UAH":         "sberbank-code",
		"Chainlink (LINK)":     "chainlink",
		"Cosmos (ATOM)":        "cosmos",
		"ecoPayz":              "ecopayz",
		"Polkadot (DOT)":       "polkadot",
		"Uniswap (UNI)":        "uniswap",
		"Dai (DAI)":            "dai",
		"Ravencoin (RVN)":      "ravencoin",
		"Jysan Bank":           "jysan-bank",
		"Tether BEP20 (USDT)":  "tether-bep20",
		"Shiba Inu (SHIB)":     "shiba-inu",
		"Ethereum BEP20 (ETH)": "ethereum-bep20",
		"Maker (MKR)":          "maker",
		"WhiteBIT":             "whitebit",
		"Хоум Кредит":          "homecredit",
		"Algorand (ALGO)":      "algorand",
		"Avalanche (AVAX)":     "avalanche",
		"Wrapped ETH (WETH)":   "wrapped-eth",
		"Любой банк GEL":       "wire-gel",
		"Yearn.finance (YFI)":  "yearn-finance",
		"Visa/MasterCard CAD":  "visa-mastercard-cad",
		"Visa/MasterCard BGN":  "visa-mastercard-bgn",
		"Visa/MasterCard HUF":  "visa-mastercard-huf",
		"Visa/MasterCard CZK":  "visa-mastercard-czk",
		"Visa/MasterCard NOK":  "visa-mastercard-nok",
		"Capitalist EUR":       "capitalist-euro",
		"Decentraland (MANA)":  "decentraland",
		"Любой банк USD":       "wire-usd",
		"Любой банк EUR":       "wire-eur",
		"Наличные USD":         "dollar-cash",
		"Binance RUB":          "binance-rub",
		"Наличные EUR":         "euro-cash",
		"Криптобиржи USD":      "trade-usd",
		"Криптобиржи EUR":      "trade-eur",
		"Binance USD (BUSD)":   "busd",
	}
)
// Coin has id, FullName and ShortName
type Coin struct {
	ID        string `json:"coin_id"`
	FullName  string `json:"coin_full_name"`
	ShortName string `json:"coin_short_name"`
}

// Coins is map[string]Coin, which key is the ID
// and value is a Coin
type Coins map[string]Coin

func getCoins(data string) (Coins, error) {
	res := Coins{}

	// working with first file
	s := ""
	for _, v := range data {
		if v == '\n' {
			splt := strings.Split(s, ";")
			if len(splt) < 4 {
				return nil, fmt.Errorf("invalid data")
			}

			ID := splt[0]
			fullName := splt[2]
			shortName := splt[3]

			res[splt[0]] = Coin{
				ID,
				fullName,
				shortName,
			}
			s = ""
			continue
		}
		s += string(v)
	}
	splt := strings.Split(s, ";")
	if len(splt) < 4 {
		return nil, fmt.Errorf("invalid data")
	}
	res[splt[0]] = Coin{
		splt[0],
		splt[2],
		splt[3],
	}
	return res, nil
}

func newCoins() (Coins, error) {
	data, err := openFile(coinsFileName)
	if err != nil {
		return nil, err
	}

	coins, err := getCoins(data)
	if err != nil {
		return nil, err
	}
	return coins, nil
}
