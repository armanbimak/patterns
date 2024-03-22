package main

import (
	"fmt"
	"github.com/armanbimak/patterns/structural/decorator/component"
	topping_decorator2 "github.com/armanbimak/patterns/structural/decorator/topping_decorator"
)

func main() {

	pizza := component.NewCheesePizza(10, "cheese pizza")

	pizzaWithArtichokes := topping_decorator2.NewArtichokesTopping(pizza, 2, "artichokes")
	pizzaWithArtichokesTabasco := topping_decorator2.NewTabascoTopping(pizzaWithArtichokes, 1, "tabasco")

	fmt.Println(pizzaWithArtichokesTabasco.Description())
	fmt.Println(pizzaWithArtichokesTabasco.Cost())
}
