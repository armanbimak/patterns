package main

import (
	"fmt"
	"github.com/armanbimak/patterns/decorator/component"
	"github.com/armanbimak/patterns/decorator/topping_decorator"
)

func main() {

	pizza := component.NewCheesePizza(10, "cheese pizza")

	pizzaWithArtichokes := topping_decorator.NewArtichokesTopping(pizza, 2, "artichokes")
	pizzaWithArtichokesTabasco := topping_decorator.NewTabascoTopping(pizzaWithArtichokes, 1, "tabasco")

	fmt.Println(pizzaWithArtichokesTabasco.Description())
	fmt.Println(pizzaWithArtichokesTabasco.Cost())
}
