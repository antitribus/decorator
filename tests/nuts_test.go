package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestNuts_AddIngredient(t *testing.T) {
	nuts := &types.Nuts{}
	name, price, err := nuts.AddIngredient()

	if err == nil {
		t.Errorf(`When calling AddIngredient on the nuts decorator without an IngredientAdder on its 
		Ingredient field it must return an error, not a string with "%s"`, name)
	}

	if name != "" {
		t.Errorf(`Nuts result "%s" should be "%s"`, name, "")
	}

	if price != 0 {
		t.Errorf(`Nuts price "%f" should be "%d"`, price, 0)
	}

	nuts = &types.Nuts{&types.Cupcake{}}
	name, price, err = nuts.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	priceExpected := 1.20

	if price != float64(priceExpected) {
		t.Errorf(`Nuts price "%f" should be "%f"`, price, priceExpected)
	}

	if !strings.Contains(name, "nuts") {
		t.Errorf(`When calling the add ingredient of the nuts decorator it must 
		return a text with the word "Nuts", not "%s"`, name)
	}
}
