package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name:"tom", age:30})
	fmt.Println(person{name:"fred"})
	fmt.Println(&person{age:40})

	s := person{"my", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 52
	fmt.Println(sp.age)
}

