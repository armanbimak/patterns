package topping_decorator

import (
	"fmt"
	"github.com/armanbimak/patterns/decorator/component"
)

type TabascoTopping struct {
	component   component.IPizza
	price       int
	description string
}

func NewTabascoTopping(component component.IPizza, price int, description string) component.IPizza {
	return ArtichokesTopping{
		component:   component,
		price:       price,
		description: description,
	}
}

func (t TabascoTopping) Cost() int {
	return t.component.Cost() + t.price
}

func (t TabascoTopping) Description() string {
	return fmt.Sprintf("%s, %s", t.component.Description(), t.description)
}
