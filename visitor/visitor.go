package main

type visitor interface {
	visitForStocks(*stocks)
	visitForBonds(*bonds)
	visitForFutures(*futures)
}
