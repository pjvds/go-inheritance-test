package main

import (
	"fmt"
)

func main() {
	f := Foo{
		A: "foo",
	}
	b := Bar{
		Foo: f,
		B:   "bar",
	}

	f.NonPointerMethod()
	b.NonPointerMethod()

	f.PointerMethod()
	b.PointerMethod()

	AcceptingFooAsArgument(f)
	AcceptingFooAsArgument(b)
}

type Foo struct {
	A string
}

type Bar struct {
	Foo
	B string
}

func (f Foo) NonPointerMethod() {
	fmt.Println("NonPointerMethod: " + f.A)
}

func (f *Foo) PointerMethod() {
	fmt.Println("PointerMethod: " + f.A)
}

func AcceptingFooAsArgument(f Foo) {
	fmt.Println("AcceptingFooAsArgument: " + f.A)
}
