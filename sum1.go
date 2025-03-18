package main

import(
	"fmt"
)

func main(){
	var name string = "Alice"
	age:= 25
	fmt.Println("Hello", name, age)
	num:=10
	if num >5{
		fmt.Println("Greater than 5")
	}else if num == 10{
		fmt.Println("This is me")
	}

	for i:=1; i<num; i++{
		fmt.Println(i)
	}
	//No while loop
	for num>5{
		fmt.Print(num)
		num--
	}
	add(12,5)
}

func add(num int, b int) int{
	return num*b;
}

