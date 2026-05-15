package main

import "fmt"

type Cat struct {
	name  string
	age   int
	color string
}

func newCat(name string, age int, color string) Cat {
	c := Cat{name, age, color}
	return c
}

func (c Cat) getAge() int {
	return c.age
}

func (c *Cat) setAge(a int) {
	if a > 0 && a < 30 {
		c.age = a
	} else {
		fmt.Println("Неправильный возраст")
	}
}

func (c Cat) show() {
	fmt.Println("Кличка:", c.name)
	fmt.Println("Возраст:", c.age)
	fmt.Println("Цвет:", c.color)
}

func main() {
	var name string
	var age int
	var color string

	fmt.Print("Кличка: ")
	fmt.Scan(&name)

	fmt.Print("Возраст: ")
	fmt.Scan(&age)

	fmt.Print("Цвет: ")
	fmt.Scan(&color)

	// создаем кошку
	cat := newCat(name, age, color)

	// меняем возраст
	var newAge int
	fmt.Print("Введи новый возраст: ")
	fmt.Scan(&newAge)
	cat.setAge(newAge)

	fmt.Println("Новый возраст:", cat.getAge())
}
