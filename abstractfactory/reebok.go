package main

type reebok struct{}

func (a *reebok) makeShoe() sneakers {
	return &reebokShoe{
		shoe: shoe{
			logo: "reebok",
			size: 14,
		},
	}
}

func (a *reebok) makeShirt() tshirt {
	return &reebokShirt{
		shirt: shirt{
			logo: "reebok",
			size: 14,
		},
	}
}
