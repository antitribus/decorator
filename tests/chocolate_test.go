package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestChocolate_AddIngredient(t *testing.T) {
	chocolate := &types.Chocolate{}
	name, price, err := chocolate.AddIngredient()

	if err == nil {
		t.Errorf("When calling AddIngredient on the chocolate decorator without "+
			"an IngredientAdder on its Ingredient field it must return an error, "+
			"not a string with '%s'", name)
	}

	if name != "" {
		t.Errorf("Cholocate result \"%s\" should be \"%s\"", name, "")
	}

	if price != 0 {
		t.Errorf("Cholocate price %f should be %d", price, 0)
	}

	chocolate = &types.Chocolate{&types.Cupcake{}}
	name, price, err = chocolate.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	priceExpected := 1.10

	if price != float32(priceExpected) {
		t.Errorf("Cholocate price %f should be %f", price, priceExpected)
	}

	if !strings.Contains(name, "chocolate") {
		t.Errorf("When calling the add ingredient of the chocolate decorator it "+
			"must return a text with the word 'chocolate', not '%s'", name)
	}
}
