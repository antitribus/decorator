package tests

import (
	"decorator/types"
	"log"
	"strings"
	"testing"
)

func TestCupcake_AddIngredient(t *testing.T) {
	cupcake := &types.Cupcake{}
	cupcakeResult, _, err := cupcake.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	expectedText := "Cupcake with"
	if !strings.Contains(cupcakeResult, expectedText) {
		t.Errorf("When calling the add ingredient of the cupcake decorator it "+
			"must return the text %s, not '%s'", expectedText, cupcakeResult)
	}
}

func TestCupcake_FullStack(t *testing.T) {
	cupcake := &types.Sugar{&types.Nuts{&types.Chocolate{&types.Cupcake{}}}}
	cupcakeResult, price, err := cupcake.AddIngredient()

	log.Println(">>>>>>>>>>>>> ", price)

	if err != nil {
		t.Error(err)
	}

	expectedText := "Cupcake with chocolate nuts sugar"

	if !strings.Contains(cupcakeResult, expectedText) {
		t.Errorf("When asking for a cupcake with onion and meat the returned "+
			"string must contain the text '%s' but '%s' didn't have it", expectedText,
			cupcakeResult)
	}

	t.Log(cupcakeResult)
}
