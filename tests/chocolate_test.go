package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestChocolate_AddIngredient(t *testing.T) {
	chocolate := &types.Chocolate{}
	chocolateResult, _, err := chocolate.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the chocolate decorator without "+
			"an IngredientAdder on its Ingredient field it must return an error, "+
			"not a string with '%s'", chocolateResult)
	}

	chocolate = &types.Chocolate{&types.Cupcake{}}
	chocolateResult, _, err = chocolate.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(chocolateResult, "chocolate") {
		t.Errorf("When calling the add ingredient of the chocolate decorator it "+
			"must return a text with the word 'chocolate', not '%s'", chocolateResult)
	}
}
