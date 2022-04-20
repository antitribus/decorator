package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestSugar_AddIngredient(t *testing.T) {
	sugar := &types.Sugar{}
	name, price, err := sugar.AddIngredient()

	if err == nil {
		t.Errorf("When calling AddIngredient on the sugar decorator without "+
			"an IngredientAdder on its Ingredient field it must return an error, "+
			"not a string with '%s'", name)
	}

	if name != "" {
		t.Errorf("Sugar result \"%s\" should be \"%s\"", name, "")
	}

	if price != 0 {
		t.Errorf("Sugar price %f should be %d", price, 0)
	}

	sugar = &types.Sugar{&types.Cupcake{}}
	name, price, err = sugar.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	priceExpected := 1.00

	if price != float64(priceExpected) {
		t.Errorf(`Sugar price "%f" should be "%f"`, price, priceExpected)
	}

	if !strings.Contains(name, "sugar") {
		t.Errorf(`When calling the add ingredient of the sugar decorator it 
		must return a text with the word "sugar", not "%s"`, name)
	}
}
