# bestchange-api wrapper on golang

## Usage

Simple example of usage:
```bash
$ touch test.go
```

```go
package main

import (
    bca "github.com/bulatok/bestchange-api"
    "log"
    "fmt"
)

func main(){
	bc, err := bca.NewBestchange()
	if err != nil{
		log.Fatal(err)
        }
	
	
	rates, err := bc.GetRatesFromTo("BTC", "QIWI") // BTC -> QIWI
	if err != nil{
		log.Fatal(err)
	}
	
	for _, v := range rates{
		fmt.Println(v.String())
	}
}
```


```bash
   $ go run test.go
```
Result - 
```bash
BTC -> RUB QIWI
Market - BitOkk
Price - 3688119.19000004 (can buy from 0.000724 - to 0.014481)
Market link - https://www.bestchange.ru/click.php?id=749&from=93&to=63&city=0
Link to full list - https://www.bestchange.ru/bitcoin-to-qiwi.html

BTC -> RUB QIWI
Market - GrandChange
Price - 2830009.02530003 (can buy from 0.0025 - to 100)
Market link - https://www.bestchange.ru/click.php?id=980&from=93&to=63&city=0
Link to full list - https://www.bestchange.ru/bitcoin-to-qiwi.html

BTC -> RUB QIWI
Market - RoyalCash
Price - 3021598.27540002 (can buy from 0.00165475 - to 29.78869)
Market link - https://www.bestchange.ru/click.php?id=929&from=93&to=63&city=0
Link to full list - https://www.bestchange.ru/bitcoin-to-qiwi.html

.........
......... # etc
```

### or

```go
package main

import (
    bca "github.com/bulatok/bestchange-api"
    "log"
    "fmt"
    "time"
)

func main() {
	bc, err := bca.NewBestchange()
	if err != nil {
		log.Fatal(err)
	}

	bc.SortRatesByPrice() // sorting the rates by their prices
	
	rates, err := bc.GetRatesFromTo("BTC", "QIWI") // BTC -> QIWI
	if err != nil{
		log.Fatal(err)
	}
	

	for _, v := range rates{
		json, err := v.JSON()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(json)
	}

	time.Sleep(time.Minute) // just for example to see the differences

	if err := bc.UpdateRates(); err != nil{ // we can easily update the rates without creating new objects.
		log.Fatal(err)
	}

	for _, v := range rates{
		json, err := v.JSON()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(json)
	}
}
```


```bash
   $ go run test.go
```
Result -
```bash
{
  "Rate": {
    "coin_from": {
      "coin_id": "93",
      "coin_full_name": "Bitcoin (BTC)",
      "coin_short_name": "BTC"
    },
    "coin_to": {
      "coin_id": "63",
      "coin_full_name": "QIWI RUB",
      "coin_short_name": "RUB QIWI"
    },
    "price": "2294314.64999997",
    "price_from": "0.0130758",
    "price_till": "0.21793",
    "market": {
      "market_id": "577",
      "market_name": "Bit-Обменка"
    },
    "rating": "0.330"
  },
  "MarketLink": "https://www.bestchange.ru/bitcoin-to-qiwi.html",
  "FullListLink": "https://www.bestchange.ru/bitcoin-to-qiwi.html"
}
.........
......... # etc

# after 1 minute

{
  "Rate": {
    "coin_from": {
      "coin_id": "93",
      "coin_full_name": "Bitcoin (BTC)",
      "coin_short_name": "BTC"
    },
    "coin_to": {
      "coin_id": "63",
      "coin_full_name": "QIWI RUB",
      "coin_short_name": "RUB QIWI"
    },
    "price": "2268985.29149998",
    "price_from": "0.01322177",
    "price_till": "0.22036282",
    "market": {
      "market_id": "577",
      "market_name": "Bit-Обменка"
    },
    "rating": "0.330"
  },
  "MarketLink": "https://www.bestchange.ru/bitcoin-to-qiwi.html",
  "FullListLink": "https://www.bestchange.ru/bitcoin-to-qiwi.html"
}
.......
....... # etc
```
