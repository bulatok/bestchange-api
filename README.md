# bestchange-api wrapper on golang

## Installation
1) Install the package
```bash
    go get -u github.com/bulatok/bestchange-api
```
2) Import to your code
```go
    import bcapi "github.com/bulatok/bestchange-api"
```
## Usage

Simple example of usage:
```bash
$ touch test.go
```

```go
package main

import (
	bcapi "github.com/bulatok/bestchange-api"
	"log"
	"fmt"
	"net/http"
)

func main() {
	myClient := &http.Client{
		Timeout: 30 * time.Second,
        }
	bc, err := bcapi.NewBestchange(myClient)
	if err != nil {
		log.Fatal(err)
	}

	rates, err := bc.GetRatesFromTo("Bitcoin (BTC)", "QIWI RUB") // BTC -> QIWI
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range rates {
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
Market - ExLine
Price (BTC) - 1 (can buy from 0.006 - to 0.024594)
Recive (RUB QIWI) - 4066000.00000005
Rating - 0.738
Market link - https://www.bestchange.ru/click.php?id=952&from=93&to=63&city=0
Link to full list - https://www.bestchange.ru/bitcoin-to-qiwi.html

BTC -> RUB QIWI
Market - 1Обмен
Price (BTC) - 1 (can buy from 0.008 - to 0.188452)
Recive (RUB QIWI) - 3986705.50779997
Rating - 0.11174 # negative - 0, positive - 11174
Market link - https://www.bestchange.ru/click.php?id=666&from=93&to=63&city=0
Link to full list - https://www.bestchange.ru/bitcoin-to-qiwi.html

.........
......... # etc
```

### or

```go
package main

import (
    bcapi "github.com/bulatok/bestchange-api"
    "log"
    "fmt"
    "time"
	"net/http"
)

func main() {
	myClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	bc, err := bcapi.NewBestchange(myClient)
	if err != nil {
		log.Fatal(err)
	}
	
	rates, err := bc.GetRatesFromTo("Bitcoin (BTC)", "QIWI RUB") // BTC -> QIWI
	if err != nil{
		log.Fatal(err)
	}
	
	rates = bcapi.SortRatesByReceive(rates) // sorting the rates by their receive prices

	for _, v := range rates{
		json, err := v.JSON()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(json)
	}

	time.Sleep(time.Minute) // just for example to see the differences

	if err := bc.UpdateRates(myClient); err != nil{ // we can easily update the rates without creating new objects.
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
    "price": "1",
    "price_from": "0.01334854",
    "price_till": "0.22247561",
    "market": {
      "market_id": "577",
      "market_name": "Bit-Обменка"
    },
    "rating": "0.330",
    "receive": "2247437.4"
  },
  "MarketLink": "https://www.bestchange.ru/bitcoin-to-qiwi.html",
  "FullListLink": "https://www.bestchange.ru/bitcoin-to-qiwi.html"
}
.........
......... # etc

# after 1 minute
# it is another offer
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
    "price": "1",
    "price_from": "0.000185",
    "price_till": "15",
    "market": {
      "market_id": "992",
      "market_name": "OneMoment"
    },
    "rating": "0.4097",
    "receive": "2700143.55999998"
  },
  "MarketLink": "https://www.bestchange.ru/bitcoin-to-qiwi.html",
  "FullListLink": "https://www.bestchange.ru/bitcoin-to-qiwi.html"
}
.......
....... # etc
```
