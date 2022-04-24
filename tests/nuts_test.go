package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestNuts_AddIngredient(t *testing.T) {
	nuts := &types.Nuts{}
	name, err := nuts.AddIngredient()

	if err == nil {
		t.Errorf(`When calling AddIngredient on the nuts decorator without an IngredientAdder on its 
		Ingredient field it must return an error, not a string with "%s"`, name)
	}

	if name != "" {
		t.Errorf(`Nuts result "%s" should be "%s"`, name, "")
	}

	nuts = &types.Nuts{&types.Cupcake{}}
	name, err = nuts.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(name, "nuts") {
		t.Errorf(`When calling the add ingredient of the nuts decorator it must 
		return a text with the word "Nuts", not "%s"`, name)
	}
}

func TestNuts_GetPrice(t *testing.T) {
	nuts := &types.Nuts{}
	price, err := nuts.GetPrice()

	if err == nil {
		t.Error(err)
	}

	if price != 0 {
		t.Errorf(`Nuts price "%f" should be "%d"`, price, 0)
	}

	nuts = &types.Nuts{&types.Cupcake{}}

	price, err = nuts.GetPrice()

	if err != nil {
		t.Error(err)
	}

	priceExpected := 1.20

	if price != priceExpected {
		t.Errorf(`Nuts price "%f" should be "%f"`, price, priceExpected)
	}
}

func TestNuts_Validate(t *testing.T) {
	nuts := &types.Nuts{}
	err := nuts.Validate()

	if err == nil {
		t.Error("Should be return a validation error, Ingredient is nil")
	}

	if err != nil && !strings.Contains(err.Error(), "IngredientAdder is needed") {
		t.Error(err)
	}

	nuts = &types.Nuts{&types.Cupcake{}}
	err = nuts.Validate()

	if err != nil {
		t.Error(err)
	}
}