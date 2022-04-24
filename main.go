package main

import (
	"decorator/types"
	"fmt"
)

func main() {
	fmt.Println("Preparing some cupcakes 🧁...")

	cupcakes := []types.IngredientAdder{
		&types.Chocolate{&types.Cupcake{}},
		&types.Nuts{&types.Cupcake{}},
		&types.Sugar{&types.Cupcake{}},
		&types.Nuts{&types.Chocolate{&types.Cupcake{}}},
		&types.Sugar{&types.Chocolate{&types.Cupcake{}}},
		&types.Sugar{&types.Nuts{&types.Chocolate{&types.Cupcake{}}}},
	}

	for _, cupcake := range cupcakes {
		description, _ := cupcake.AddIngredient()
		price, _ := cupcake.GetPrice()

		fmt.Printf("%s costing $%.2f\n", description, price)
	}
}
