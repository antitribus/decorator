package types

type IngredientAdder interface {
	AddIngredient() (string, error)
	GetPrice() (float64, error)
	Validate() error
}
