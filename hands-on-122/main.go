package main

import "fmt"

func main() {

	l := make(map[string][]string)
	l["bond_james"] = []string{"Shaken, not stirred", "Martini", "fast cars"}
	l["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	l["no_dr"] = []string{"Being  evil", "Ice cream", "Sunsets"}
	fmt.Println(l)

	fmt.Println("Adding record to the current map")

	l["fleming_jan"] = []string{"steaks", "cigars", "espionage"}

	for x := range l {
		fmt.Printf("Key is %v\n",x)
		for a := range l[x] {
			fmt.Printf("Value is %v\n", l[x][a])
		}
	}

	fmt.Println("Another method to print the map")

	for x,y:= range l{
		fmt.Printf("Key is %v and Value is %v\n",x,y)
		for a,b:=range y {
			fmt.Printf("Index:%v \t Value:%v\n",a,b)

		}
	}
}