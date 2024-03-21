package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select an option")
	fmt.Println("1 print Menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)

	//output
	type menuItem struct {
		name  string
		price map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", price: map[string]float64{"small": 1.65, "medium": 2.10, "large": 2.90}},
		{name: "Latte", price: map[string]float64{"small": 1.90, "medium": 2.80, "large": 3.65}},
	}

	//fmt.Println(menu)

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range menu {
			fmt.Println("\t%10s%10.2f\n", size, price)

		}
	}
}
