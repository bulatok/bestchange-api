package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer func(t time.Time) {
		fmt.Println(time.Since(t))
	}(time.Now())
	bc, err := NewBestchange()
	if err != nil {
		log.Fatalln(err)
	}



	bc.SortRatesByPrice()

	res, err := bc.GetRatesFromTo("BTC", "QIWI RUB")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res{
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

	for _, v := range res{
		json, err := v.JSON()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(json)
	}
}
