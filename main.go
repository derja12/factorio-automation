package main

import "fmt"

type Item struct {
	name    string
	recipe  *Recipe
	machine *Machine
}

type Recipe struct {
	ingredients []*Ingredient
	count       int
	time        float32
}

type Ingredient struct {
	item  *Item
	count int
}

type Machine struct {
	name string
}

func main() {
	InitializeRedPotionItems()
}

func InitializeRedPotionItems() error {

	copperPlateItem := InitializeCopperPlate()
	IronGearItem := InitializeIronGear()

	fmt.Println(copperPlateItem.name, copperPlateItem.recipe, copperPlateItem.machine, IronGearItem.name, IronGearItem.recipe.ingredients[0].item.name)
	return nil
}

func InitializeCopperPlate() *Item {
	// Create Copper Ore Item / Ingredient
	copperOreItem := &Item{
		name: "Copper Ore",
	}
	copperOreIngredient := &Ingredient{
		item:  copperOreItem,
		count: 1,
	}

	// Create Copper Plate Recipe
	var ingredientsList []*Ingredient
	ingredientsList = append(ingredientsList, copperOreIngredient)
	copperPlateRecipe := &Recipe{
		ingredients: ingredientsList,
		count:       1,
		time:        3.2,
	}

	// Create Copper Plate Item
	copperPlateItem := &Item{
		name:    "Copper Plate",
		recipe:  copperPlateRecipe,
		machine: nil,
	}

	return copperPlateItem
}

func InitializeIronGear() *Item {
	// Create Iron Ore Item / Ingredient
	ironOreItem := &Item{
		name: "Iron Ore",
	}
	ironOreIngredient := &Ingredient{
		item:  ironOreItem,
		count: 1,
	}

	// Create Iron Plate Recipe
	var ingredientsList []*Ingredient
	ingredientsList = append(ingredientsList, ironOreIngredient)
	ironPlateRecipe := &Recipe{
		ingredients: ingredientsList,
		count:       1,
		time:        3.2,
	}

	// Create Iron Plate Item
	ironPlateItem := &Item{
		name:    "Iron Plate",
		recipe:  ironPlateRecipe,
		machine: nil,
	}
	ironPlateIngredient := &Ingredient{
		item:  ironPlateItem,
		count: 2,
	}

	// Create Iron Gear Recipe
	ingredientsList = []*Ingredient{}
	ingredientsList = append(ingredientsList, ironPlateIngredient)
	ironGearRecipe := &Recipe{
		ingredients: ingredientsList,
		count:       1,
		time:        .5,
	}

	// Create Iron Gear Item
	ironGearItem := &Item{
		name:    "Iron Gear",
		recipe:  ironGearRecipe,
		machine: nil,
	}

	return ironGearItem
}
