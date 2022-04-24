package types

import (
	"errors"
	"fmt"
)

type Sugar struct {
	Ingredient IngredientAdder
}

func (s *Sugar) AddIngredient() (string, error) {
	err := s.Validate()

	if err != nil {
		return "", err
	}

	name, err := s.Ingredient.AddIngredient()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", name, "sugar"), nil
}

func (s *Sugar) GetPrice() (float64, error) {
	err := s.Validate()

	if err != nil {
		return 0, err
	}

	return s.Ingredient.GetPrice()
}

func (s *Sugar) Validate() error {
	if s.Ingredient == nil {
		return errors.New("An IngredientAdder is needed on the Ingredient field of the Sugar")
	}

	return nil
}