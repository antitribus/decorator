package tests

import (
	"decorator/types"
	"strings"
	"testing"
)

func TestCupcake_AddIngredient(t *testing.T) {
	cupcake := &types.Cupcake{}
	name, price, err := cupcake.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	expectedPrice := 1.00

	if price != float64(expectedPrice) {
		t.Errorf(`Cupcake price "%f" should be "%f"`, price, expectedPrice)
	}

	expectedText := "Cupcake with"

	if name != expectedText {
		t.Errorf(`When calling the add ingredient of the cupcake decorator it 
		must return the text "%s", not "%s`, expectedText, name)
	}
}

func TestCupcake_FullStack(t *testing.T) {
	cupcake := &types.Sugar{&types.Nuts{&types.Chocolate{&types.Cupcake{}}}}
	name, price, err := cupcake.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	expectedPrice := 1.30

	if price != expectedPrice {
		t.Errorf(`Cupcake price "%f" should be "%f"`, price, expectedPrice)
	}

	expectedText := "Cupcake with chocolate nuts sugar"

	if !strings.Contains(name, expectedText) {
		t.Errorf(`When asking for a cupcake with chocolate, nuts and sugar the returned 
			string must contain the text "%s" but "%s" didn't have it`, expectedText, name)
	}
}
