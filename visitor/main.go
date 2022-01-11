package main

import "fmt"

func main() {
	stocksObj := &stocks{}
	bondsObj := &bonds{}
	futuresObj := &futures{}

	riskCalc := &riskCalculator{}
	stocksObj.accept(riskCalc)
	bondsObj.accept(riskCalc)
	futuresObj.accept(riskCalc)
	fmt.Println("------------------------------------------------------------------------")
	profitCalc := &profitCalculator{}
	stocksObj.accept(profitCalc)
	bondsObj.accept(profitCalc)
	futuresObj.accept(profitCalc)
}
