package main

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe{
			logo: "adidas",
		},
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt{
			logo: "adidas",
		},
	}
}