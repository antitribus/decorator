package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestChocolate_AddIngredient(t *testing.T) {
	chocolate := &types.Chocolate{}
	name, err := chocolate.AddIngredient()

	if err == nil {
		t.Errorf(`When calling AddIngredient on the chocolate decorator without an IngredientAdder on its 
		Ingredient field it must return an error, not a string with "%s"`, name)
	}

	if name != "" {
		t.Errorf(`Cholocate result "%s" should be "%s"`, name, "")
	}

	chocolate = &types.Chocolate{&types.Cupcake{}}
	name, err = chocolate.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(name, "chocolate") {
		t.Errorf(`When calling the add ingredient of the chocolate decorator it must return a text with 
		the word "chocolate", not "%s"`, name)
	}
}

func TestChocolate_GetPrice(t *testing.T) {
	chocolate := &types.Chocolate{}
	price, err := chocolate.GetPrice()

	if err == nil {
		t.Error(`When calling GetPrice on the chocolate decorator without an IngredientAdder on its Ingredient field it must return an error`)
	}

	if price != 0 {
		t.Errorf("CHocolate price %f should be %d", price, 0)
	}

	chocolate = &types.Chocolate{&types.Cupcake{}}
	price, err = chocolate.GetPrice()

	if err != nil {
		t.Error(err)
	}

	priceExpected := 1.1

	if price != priceExpected {
		t.Errorf(`Chocolate price "%f" should be "%f"`, price, priceExpected)
	}
}

func TestChocolate_Validate(t *testing.T) {
	chocolate := &types.Chocolate{}
	err := chocolate.Validate()

	if err == nil {
		t.Error("Should be return a validation error, Ingredient is nil")
	}

	if err != nil && !strings.Contains(err.Error(), "IngredientAdder is needed") {
		t.Error(err)
	}

	chocolate = &types.Chocolate{&types.Cupcake{}}
	err = chocolate.Validate()

	if err != nil {
		t.Error(err)
	}
}