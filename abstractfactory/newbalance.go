package main

type newbalance struct{}

func (n *newbalance) makeShoe() sneakers {
	return &newbalanceShoe{
		shoe: shoe{
			logo: "newbalance",
			size: 14,
		},
	}
}

func (n *newbalance) makeShirt() tshirt {
	return &newbalanceShirt{
		shirt: shirt{
			logo: "newbalance",
			size: 14,
		},
	}
}
