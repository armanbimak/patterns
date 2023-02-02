package component

type MeatPizza struct {
	price       int
	description string
}

func (p MeatPizza) Cost() int {
	return p.price
}

func (p MeatPizza) Description() string {
	return p.description
}
