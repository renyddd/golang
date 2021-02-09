package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe{
			logo: "nike",
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt{
			logo: "nike",
		},
	}
}