package types

type IngredientAdder interface {
	AddIngredient() (string, float64, error)
}
