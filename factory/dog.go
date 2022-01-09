package main

type dog struct {
	being
}

func newDog() animal {
	return &dog{
		being: being{
			name:   "dog",
			sound:  "bark",
			weight: 10,
			height: 5,
		},
	}
}
