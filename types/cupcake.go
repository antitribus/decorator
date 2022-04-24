package types

type Cupcake struct {
	Ingredient IngredientAdder
}

func (c *Cupcake) AddIngredient() (string, error) {
	return "Cupcake with:", nil
}

func (c *Cupcake) GetPrice() (float64, error) {
	return 1.0, nil
}

func (c *Cupcake) Validate() error {
	return nil
}