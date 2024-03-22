package topping_decorator

import (
	"fmt"
	"github.com/armanbimak/patterns/structural/decorator/component"
)

type ArtichokesTopping struct {
	component   component.IPizza
	price       int
	description string
}

func NewArtichokesTopping(component component.IPizza, price int, description string) component.IPizza {
	return ArtichokesTopping{
		component:   component,
		price:       price,
		description: description,
	}
}

func (t ArtichokesTopping) Cost() int {
	return t.component.Cost() + t.price
}

func (t ArtichokesTopping) Description() string {
	return fmt.Sprintf("%s, %s", t.component.Description(), t.description)
}
