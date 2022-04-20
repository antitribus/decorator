package main

import (
	"decorator/types"
	"fmt"
)

func main() {
	fmt.Println("Preparando alguns cupcakes 🧁...")

	cupcakes := []types.IngredientAdder{
		&types.Chocolate{&types.Cupcake{}},
		&types.Nuts{&types.Cupcake{}},
		&types.Sugar{&types.Cupcake{}},
		&types.Nuts{&types.Chocolate{&types.Cupcake{}}},
		&types.Sugar{&types.Chocolate{&types.Cupcake{}}},
		&types.Sugar{&types.Nuts{&types.Chocolate{&types.Cupcake{}}}},
	}

	for _, cupcake := range cupcakes {
		description, price, _ := cupcake.AddIngredient()
		fmt.Printf("%s no valor de R$%.2f\n", description, price)
	}
}
