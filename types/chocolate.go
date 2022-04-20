package types

import (
	"errors"
	"fmt"
)

type Chocolate struct {
	Ingredient IngredientAdder
}

func (c *Chocolate) AddIngredient() (string, float64, error) {
	if c.Ingredient == nil {
		return "", 0, errors.New("An IngredientAdder is needed on the Ingredient field of the Chocolate")
	}

	name, price, err := c.Ingredient.AddIngredient()

	if err != nil {
		return "", 0, err
	}

	chocoPrice := price + 0.1

	return fmt.Sprintf("%s %s,", name, "chocolate"), chocoPrice, nil
}
