package main

type stocks struct { // some fields to calculate profitability and risk
}

func (s *stocks) accept(v visitor) {
	v.visitForStocks(s)
}

type bonds struct { // some fields to calculate profitability and risk
}

func (b *bonds) accept(v visitor) {
	v.visitForBonds(b)
}

type futures struct { // some fields to calculate profitability and risk
}

func (f *futures) accept(v visitor) {
	v.visitForFutures(f)
}
