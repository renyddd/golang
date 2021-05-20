package main

import (
	"fmt"
)

// https://refactoringguru.cn/design-patterns/abstract-factory/go/example#example-0
type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand passed")
	}
}

type adidas struct{}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
		},
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
		},
	}
}

type nike struct{}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
		},
	}
}

type iShoe interface {
	setLogo(string)
}

type shoe struct {
	logo string
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

type adidasShoe struct {
	shoe
}

type nikeShoe struct {
	shoe
}

type iShirt interface {
	setLogo(string)
}

type shirt struct {
	logo string
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

type adidasShirt struct {
	shirt
}

type nikeShirt struct {
	shirt
}
