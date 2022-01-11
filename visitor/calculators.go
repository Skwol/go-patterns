package main

import (
	"fmt"
)

type riskCalculator struct{}

func (r riskCalculator) visitForStocks(s *stocks) {
	fmt.Println("calculating risk for stocks")
}

func (r riskCalculator) visitForBonds(b *bonds) {
	fmt.Println("calculating risk for bonds")
}

func (r riskCalculator) visitForFutures(f *futures) {
	fmt.Println("calculating risk for futures")
}

type profitCalculator struct{}

func (p profitCalculator) visitForStocks(s *stocks) {
	fmt.Println("calculating profit for stocks")
}

func (p profitCalculator) visitForBonds(b *bonds) {
	fmt.Println("calculating profit for bonds")
}

func (p profitCalculator) visitForFutures(f *futures) {
	fmt.Println("calculating profit for futures")
}
