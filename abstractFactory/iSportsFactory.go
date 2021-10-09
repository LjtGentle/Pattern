package main

import (
	"errors"
)

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

const (
	Adidas  string = "adidas"
	Nike string = "nick"
)

func getSportsFactory(brand string)(iSportsFactory, error) {
	if brand == Adidas {
		return &adidas{}, nil
	}
	if brand == Nike {
		return &nike{}, nil
	}
	return nil, errors.New("Wrong brand type passed")
}

