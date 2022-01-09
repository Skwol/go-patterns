package main

type cat struct {
	being
}

func newCat() animal {
	return &cat{
		being: being{
			name:   "cat",
			sound:  "mew",
			weight: 5,
			height: 3,
		},
	}
}
