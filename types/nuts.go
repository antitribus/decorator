package types

import (
	"errors"
	"fmt"
)

type Nuts struct {
	Ingredient IngredientAdder
}

func (n *Nuts) AddIngredient() (string, error) {
	if n.Ingredient == nil {
		return "", errors.New("An IngredientAdder is needed on the Ingredient field of the Nuts")
	}

	name, err := n.Ingredient.AddIngredient()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", name, "nuts"), nil
}

func (n *Nuts) GetPrice() (float64, error) {
	err := n.Validate()

	if err != nil {
		return 0, err
	}

	price, err := n.Ingredient.GetPrice()
	return price + 0.2, err
}

func (n *Nuts) Validate() error {
	if n.Ingredient == nil {
		return errors.New("An IngredientAdder is needed on the Ingredient field of the Nuts")
	}

	return nil
}
