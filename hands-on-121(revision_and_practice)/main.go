package main

import "fmt"

func main() {

	// x := map[string][]string{"bond_james" : {"a", "b"} ,  "moneypenny_miss" : {"c", "d"}, "no_dr" : {"e", "f"}}
	// fmt.Println(x)

	fmt.Println("Printing Int")
	x := []int{1,2,3,4,5}
	fmt.Println(x)

	fmt.Println("---------")
	
	fmt.Println("Printing Int Version 2 using Make function")

	b := make ([]int, 5, 10)  // 5 is the length and 10 is the capacity. This will create 5 zero values and the capacity is 10. We can append to this slice until we reach the capacity of 10.
	k := append(b, 32, 5353,5454)
	fmt.Println(k) 

	// If we need to assign our own values to first 5 positions than we should do like this in a simple way:

	b[0] = 10
	b[1] = 20
	b[2] = 30
	b[3] = 40
	b[4] = 50

	// This will assign values to first 5 positions and we can append to this slice until we reach the capacity of 10. This will give 8 positions

	fmt.Println(k)

	fmt.Println("---------")

	fmt.Println("Printing String")
	y := []string{"a", "b", "c", "d", "e"}
	fmt.Println(y)

	fmt.Println("--------")	

	fmt.Println("Printing Map Version 1")
	z := map[string]string{"a":"1", "b":"2", "c":"3", "d":"4", "e":"5"}
	fmt.Println(z)

	fmt.Println("--------")		

	fmt.Println("Printing Map Version 2")
	a:= map[string][]string{"g": {"1", "10"}, "h":{"2", "9"}, "i": {"3", "8"}, "j": {"4", "7"}, "k": {"5","6"}}
	fmt.Println(a)

	fmt.Println("--------")		

	fmt.Println("Printing Map Version 3: Using Make Function")
	
	j:= make(map[string][]string)
	j["g"] = []string{"x","y","z"}
	j["h"] = []string{"p","q","r"}
	fmt.Println(j)

	fmt.Println("--------")

	fmt.Println("Now solving the problem of 121")

	l := make(map[string][]string)
	l["bond_james"] = []string{"Shaken, not stirred", "Martini", "fast cars"}
	l["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	l["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}
	fmt.Println(l)


	for i,k := range l {
		fmt.Printf("Index: %v and Value : %v \n", i,k)
		}
	
		
	fmt.Println("--------")
	
	fmt.Println("Solving for loop")
		
	for i := range l {
		fmt.Println(i)
	}

	fmt.Println("--------")
	
	fmt.Println("Solving for loop 2")
		
	for i := range l {
		fmt.Println(i)
		for j := range l[i] {
			fmt.Printf("Index is %v and Value is %v\n",i,l[i][j])		}
	}	

	fmt.Println("--------")
	
	fmt.Println("Solving for loop 3")
		
	for a,b := range l {
		fmt.Printf("Index: %v and Value : %v \n", a,b)
		for c,d := range l[a] {
			fmt.Printf("Index is %v and Value is %v\n",c,d)
		}

	}	

	fmt.Println("--------")
	
	fmt.Println("Solving for loop 4")
		
	for a := range l {
		fmt.Printf("Key is: %v \n", a)
		for c := range l[a] {
			fmt.Printf("Value is %v\n", l[a][c])
		}
	}

	fmt.Println("--------")
	
	fmt.Println("Solving for Another Loop Method:5")	// Easiest Method as far

	for x,y:= range l{
		fmt.Printf("Key is %v and Value is %v\n",x,y)
		for a,b:=range y { // Y are the different values of x and we can loop through y to get the index and value of y
			fmt.Printf("Index: %v \t Value:%v\n",a,b)
			// We can also do like
			// for _,b:=range y { // Y are the different values of x and we can loop through y to get the index and value of y
			// 	fmt.Printf("Value:%v\n"b)

		}
	}	

}