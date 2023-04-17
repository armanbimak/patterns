package component

import "fmt"

type V8EngineBMW struct{}

func (v V8EngineBMW) Power() {
	fmt.Println("V8 BMW Engine Power")
}

type V6EngineMercedes struct{}

func (v V6EngineMercedes) Power() {
	fmt.Println("V6 Mercedes Engine Power")
}

type V8EngineMercedes struct{}

func (v V8EngineMercedes) Power() {
	fmt.Println("V8 Mercedes Engine Power")
}

type V6EngineBMW struct{}

func (v V6EngineBMW) Power() {
	fmt.Println("V6 BMW Engine Power")
}
