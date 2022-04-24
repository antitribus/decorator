package types

import (
	"errors"
	"fmt"
)

type Chocolate struct {
	Ingredient IngredientAdder
}

func (c *Chocolate) AddIngredient() (string, error) {
	if c.Ingredient == nil {
		return "", errors.New("An IngredientAdder is needed on the Ingredient field of the Chocolate")
	}

	name, err := c.Ingredient.AddIngredient()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", name, "chocolate"), nil
}

func (c *Chocolate) GetPrice() (float64, error) {
	err := c.Validate()

	if err != nil {
		return 0, err
	}

	price, err := c.Ingredient.GetPrice()
	return price + 0.1, err
}

func (c *Chocolate) Validate() error {
	if c.Ingredient == nil {
		return errors.New("An IngredientAdder is needed on the Ingredient field of the Chocolate")
	}

	return nil
}