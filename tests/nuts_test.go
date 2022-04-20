package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestNuts_AddIngredient(t *testing.T) {
	nuts := &types.Nuts{}
	nutsResult, _, err := nuts.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the nuts decorator without "+
			"an IngredientAdder on its Ingredient field it must return an error, "+
			"not a string with '%s'", nutsResult)
	}

	nuts = &types.Nuts{&types.Cupcake{}}
	nutsResult, _, err = nuts.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(nutsResult, "nuts") {
		t.Errorf("When calling the add ingredient of the nuts decorator it "+
			"must return a text with the word 'Nuts', not '%s'", nutsResult)
	}
}
