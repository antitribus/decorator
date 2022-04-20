package types

import (
	"errors"
	"fmt"
)

type Sugar struct {
	Ingredient IngredientAdder
}

func (o *Sugar) AddIngredient() (string, float64, error) {
	if o.Ingredient == nil {
		return "", 0, errors.New("An IngredientAdder is needed on the Ingredient field of the Sugar")
	}

	name, price, err := o.Ingredient.AddIngredient()

	if err != nil {
		return "", 0, err
	}

	return fmt.Sprintf("%s %s,", name, "sugar"), price, nil
}
