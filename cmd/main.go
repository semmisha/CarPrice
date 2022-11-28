package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const (
	RUB                       = 73
	EngineVolumeTax           = 3.6
	TransportTax              = 13
	EuVat                     = 19
	GermanBuyerRevenuePercent = 8
)

type CarResearch struct {
	RawPrice        string `human:"raw_price"`
	EngineVolume    string `human:"engine_volume"`
	ExtraPay        string `human:"extra_pay"`
	CurrentCarPrice string `human:"current_car_price"`
}

func NewCarResearch() *CarResearch {
	return &CarResearch{}
}

func main() {

	for {
		r := NewCarResearch()

		bufio.NewReader(os.Stdin)

		fmt.Print("\n Price: ")
		_, err := fmt.Scanln(&r.RawPrice)
		if err != nil {

			fmt.Println(err)
		}
		fmt.Print("Volume cm3: ")
		_, err = fmt.Scanln(&r.EngineVolume)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Extra RUB: ")
		_, err = fmt.Scanln(&r.ExtraPay)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Car price RUB: ")
		_, err = fmt.Scanln(&r.CurrentCarPrice)
		if err != nil {
			fmt.Println(err)
		}
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Printf("%+v\n", r)

		price, _ := strconv.ParseFloat(r.RawPrice, 64)
		VAT := ProcessPrice(r.RawPrice, EuVat)
		buyerRevenue := ProcessPrice(r.RawPrice, GermanBuyerRevenuePercent)
		custom := ProcessVolume(r.EngineVolume, EngineVolumeTax)
		extraFloat, _ := strconv.ParseFloat(r.ExtraPay, 64)
		currentCarPrice, _ := strconv.ParseFloat(r.CurrentCarPrice, 64)
		currentCarTax := ProcessTax(r.CurrentCarPrice, TransportTax)
		fmt.Printf("\n Raw price EUR: %.2f\n Price without VAT EUR: %.2f\n Buyer Revenue EUR: %.2f\n Custom Revenue EUR: %.2f\n Current car price - tax RUB: %.2f\n Final Price EUR: %.2f\n Final Price RUB %.2f\n Balance RUB: %.2f\n", price, price-VAT, buyerRevenue, custom, currentCarPrice-currentCarTax, price-VAT+buyerRevenue+custom, (price-VAT+buyerRevenue+custom)*RUB+extraFloat, (currentCarPrice-currentCarTax)-((price-VAT+buyerRevenue+custom)*RUB+extraFloat))

	}
}

func ProcessPrice(rawPrice string, percent float64) float64 {
	percent = percent / (100 + percent)
	floatPrice, err := strconv.ParseFloat(rawPrice, 64)
	if err != nil {
		fmt.Println(err)
	}
	floatPrice = floatPrice * percent
	return floatPrice

}

func ProcessTax(rawPrice string, percent float64) float64 {
	percent = percent / 100
	floatPrice, err := strconv.ParseFloat(rawPrice, 64)
	if err != nil {
		fmt.Println(err)
	}
	floatPrice = (floatPrice - 6000000) * percent
	return floatPrice

}

func ProcessVolume(hp string, evt float64) float64 {

	hpINT, err := strconv.Atoi(hp)
	if err != nil {
		fmt.Println(err)
	}

	return float64(hpINT) * evt

}
