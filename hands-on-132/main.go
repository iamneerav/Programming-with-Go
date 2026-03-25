package main

import "fmt"

func main() {

p1 := struct{
	first string
	friends  map[string]int
	favDrinks  []string
}{
	first: "Neerav",
	friends:  map[string]int{"Rahul": 10, "John": 5, "Smith": 3},
	favDrinks:  []string{"Coffee", "Tea"},
}

fmt.Println(p1)

fmt.Printf("%T \t\n", p1)


for x,y := range p1.friends {
	fmt.Printf("My friend is %v and his age is %v\n",x,y)
}
	
}