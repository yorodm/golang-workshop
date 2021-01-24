package main

import "fmt"

type Human struct{}

type Animal interface {
	Name() string
}

type Dog struct{}

func (d *Dog) Name() string {
	return "Dog"
}

func  Bark(d *Dog) {
	fmt.Println("Woof!")
}

type GuideDog struct {
	*Dog
}

func (gd *GuideDog) Help(h *Human) {
	fmt.Printf("Hey human, grab %s’s leash!\n", gd.Name())
}

type Cat struct{}

func (c *Cat) Name() string {
	return "Cat"
}

type CatDog struct {
	*Cat
	*Dog
}

func (cd *CatDog) Name() string {
	return fmt.Sprintf("%s%s", cd.Cat.Name(), cd.Dog.Name())
}

func main() {
	//1 - Interfaces
	var animal Animal
	animal = &Dog{}            // returns a pointer to a new Dog
	fmt.Println(animal.Name()) // Dog

	//2 - Struct embedding
	gd := &GuideDog{}
	gd.Help(nil) // prints “Hey human, grab Dog’s leash!”

	//3 - Multi Struct embedding
	cd := &CatDog{}
	fmt.Printf("My favorite animal is the %s!\n", cd.Name())
}
