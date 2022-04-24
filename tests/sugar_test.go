package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestSugar_AddIngredient(t *testing.T) {
	sugar := &types.Sugar{}
	name, err := sugar.AddIngredient()

	if err == nil {
		t.Errorf("When calling AddIngredient on the sugar decorator without "+
			"an IngredientAdder on its Ingredient field it must return an error, "+
			"not a string with '%s'", name)
	}

	if name != "" {
		t.Errorf("Sugar result \"%s\" should be \"%s\"", name, "")
	}

	sugar = &types.Sugar{&types.Cupcake{}}
	name, err = sugar.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(name, "sugar") {
		t.Errorf(`When calling the add ingredient of the sugar decorator it 
		must return a text with the word "sugar", not "%s"`, name)
	}
}

func TestSugar_GetPrice(t *testing.T) {
	sugar := &types.Sugar{}
	price, err := sugar.GetPrice()

	if err == nil {
		t.Error(`When calling GetPrice on the sugar decorator without an IngredientAdder on its Ingredient field it must return an error`)
	}

	if price != 0 {
		t.Errorf("Sugar price %f should be %d", price, 0)
	}

	sugar = &types.Sugar{&types.Cupcake{}}
	price, err = sugar.GetPrice()

	if err != nil {
		t.Error(err)
	}

	priceExpected := 1.00

	if price != priceExpected {
		t.Errorf(`Sugar price "%f" should be "%f"`, price, priceExpected)
	}
}

func TestSugar_Validate(t *testing.T) {
	sugar := &types.Sugar{}
	err := sugar.Validate()

	if err == nil {
		t.Error("Should be return a validation error, Ingredient is nil")
	}

	if err != nil && !strings.Contains(err.Error(), "IngredientAdder is needed") {
		t.Error(err)
	}

	sugar = &types.Sugar{&types.Cupcake{}}
	err = sugar.Validate()

	if err != nil {
		t.Error(err)
	}
}