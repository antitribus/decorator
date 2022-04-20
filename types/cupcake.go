package types

type Cupcake struct {
	Ingredient IngredientAdder
}

func (c *Cupcake) AddIngredient() (string, float32, error) {
	return "Cupcake with", 1, nil
}
