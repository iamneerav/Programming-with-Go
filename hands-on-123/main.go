package main

import "fmt"

func main() {

	l := make(map[string][]string)
	l["bond_james"] = []string{"Shaken, not stirred", "Martini", "fast cars"}
	l["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	l["no_dr"] = []string{"Being  evil", "Ice cream", "Sunsets"}
	fmt.Println(l)

	fmt.Println("--------")
	fmt.Println("Deleting a record from the map")
	
	delete(l, "no_dr")  // This will delete the record with the key "no_dr" from the map l
	fmt.Println(l)

	for x,y:= range l{
		fmt.Printf("Key is %v and Value is %v\n",x,y)
		for a,b:=range y {
			fmt.Printf("Index:%v \t Value:%v\n",a,b)

		}
	}	

}