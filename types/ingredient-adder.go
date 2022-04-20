package types

type IngredientAdder interface {
	AddIngredient() (string, float32, error)
}
