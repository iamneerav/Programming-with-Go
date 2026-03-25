package main

import "fmt"

type engine struct{
	electric bool
}

type vehicle struct {
	engine
	make string
	model string
	doors string
	color string
}

func main() {

v1 := vehicle {
	engine: engine { 
	electric: true,
	},
	make: "Tesla",
	model: "Model 3",
	doors: "4",
	color: "Red",
}

v2 := vehicle {
	engine: engine {
		electric: false,
	},
	make: "Ford",
	model: "Mustang",
	doors: "2",
	color: "Blue",
}

fmt.Println(v1)

fmt.Println(v2)

fmt.Println(v1.make)
fmt.Println(v1.electric)

fmt.Println(v2.make)
fmt.Println(v2.electric)
}