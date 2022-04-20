package types

import (
	"errors"
	"fmt"
)

type Nuts struct {
	Ingredient IngredientAdder
}

func (o *Nuts) AddIngredient() (string, float32, error) {
	if o.Ingredient == nil {
		return "", 0, errors.New("An IngredientAdder is needed on the Ingredient field of the Nuts")
	}

	name, price, err := o.Ingredient.AddIngredient()

	if err != nil {
		return "", 0, err
	}

	nutsPrice := price + 0.2

	return fmt.Sprintf("%s %s", name, "nuts"), nutsPrice, nil
}
