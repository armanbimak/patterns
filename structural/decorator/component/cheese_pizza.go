package component

type CheesePizza struct {
	price       int
	description string
}

func NewCheesePizza(price int, description string) IPizza {
	return CheesePizza{
		price:       price,
		description: description,
	}
}

func (p CheesePizza) Cost() int {
	return p.price
}

func (p CheesePizza) Description() string {
	return p.description
}
