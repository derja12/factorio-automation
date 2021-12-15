package main

import "fmt"

type Item struct {
	Name    string
	recipe  *Recipe
	machine *Machine
}

type Recipe struct {
	Ingredients []Ingredient
	itemRef     *Item
	ItemName    string
	Count       int
	Time        float32
}

type Ingredient struct {
	item     *Item
	ItemName string
	Count    int
}

type Machine struct {
	Name string
}

func InitializeRedPotionItems() error {

	copperPlateItem := InitializeCopperPlate()
	IronGearItem := InitializeIronGear()

	fmt.Println(copperPlateItem.Name, copperPlateItem.recipe, copperPlateItem.machine, IronGearItem.Name, IronGearItem.recipe.ingredients[0].item.Name)
	return nil
}

func InitializeCopperPlate() *Item {
	// Create Copper Ore Item / Ingredient
	copperOreItem := &Item{
		Name: "Copper Ore",
	}
	copperOreIngredient := &Ingredient{
		item:  copperOreItem,
		Count: 1,
	}

	// Create Copper Plate Recipe
	var ingredientsList []*Ingredient
	ingredientsList = append(ingredientsList, copperOreIngredient)
	copperPlateRecipe := &Recipe{
		ingredients: ingredientsList,
		Count:       1,
		Time:        3.2,
	}

	// Create Copper Plate Item
	copperPlateItem := &Item{
		Name:    "Copper Plate",
		recipe:  copperPlateRecipe,
		machine: nil,
	}

	return copperPlateItem
}

func InitializeIronGear() *Item {
	// Create Iron Ore Item / Ingredient
	ironOreItem := &Item{
		Name: "Iron Ore",
	}
	ironOreIngredient := &Ingredient{
		item:  ironOreItem,
		Count: 1,
	}

	// Create Iron Plate Recipe
	var ingredientsList []*Ingredient
	ingredientsList = append(ingredientsList, ironOreIngredient)
	ironPlateRecipe := &Recipe{
		ingredients: ingredientsList,
		Count:       1,
		Time:        3.2,
	}

	// Create Iron Plate Item
	ironPlateItem := &Item{
		Name:    "Iron Plate",
		recipe:  ironPlateRecipe,
		machine: nil,
	}
	ironPlateIngredient := &Ingredient{
		item:  ironPlateItem,
		Count: 2,
	}

	// Create Iron Gear Recipe
	ingredientsList = []*Ingredient{}
	ingredientsList = append(ingredientsList, ironPlateIngredient)
	ironGearRecipe := &Recipe{
		ingredients: ingredientsList,
		Count:       1,
		Time:        .5,
	}

	// Create Iron Gear Item
	ironGearItem := &Item{
		Name:    "Iron Gear",
		recipe:  ironGearRecipe,
		machine: nil,
	}

	return ironGearItem
}
