package main

import "fmt"

func (r *Recipe) printRecipe() {
	fmt.Println(r.ItemName, "\t|\tCount:", r.Count, "\tTime:", r.Time)
	for _, ingredient := range r.Ingredients {
		fmt.Println("\t"+ingredient.item.Name, ingredient.Count)
	}
}
