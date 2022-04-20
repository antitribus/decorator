package types

type Cupcake struct {
	Ingredient IngredientAdder
}

func (c *Cupcake) AddIngredient() (string, float64, error) {
	return "Cupcake with", 1.0, nil
}
