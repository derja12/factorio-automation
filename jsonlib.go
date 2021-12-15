package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func SaveJsonTest() {

	recipeList := CreateRecipes()
	itemList := CreateItems()

	recipes, err1 := json.Marshal(recipeList)
	items, err2 := json.Marshal(itemList)

	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(string(recipes))
	WriteToFile("Recipes.json", string(recipes))
	WriteToFile("Items.json", string(items))
}

func CreateItems() []Item {
	list := []Item{}

	ironPlate := Item{
		Name: "IronPlate",
	}
	list = append(list, ironPlate)

	copperPlate := Item{
		Name: "CopperPlate",
	}
	list = append(list, copperPlate)

	ironGear := Item{
		Name: "IronGear",
	}
	list = append(list, ironGear)

	belt := Item{
		Name: "TransportBelt",
	}
	list = append(list, belt)

	potion := Item{
		Name: "AutoSciencePack",
	}
	list = append(list, potion)

	return list
}

func CreateRecipes() []Recipe {
	list := []Recipe{}

	gear := Recipe{
		ItemName:    "IronGear",
		Count:       1,
		Time:        .5,
		Ingredients: []Ingredient{{ItemName: "IronPlate", Count: 2}},
	}
	list = append(list, gear)

	belt := Recipe{
		ItemName:    "TransportBelt",
		Count:       2,
		Time:        .5,
		Ingredients: []Ingredient{{ItemName: "IronPlate", Count: 1}, {ItemName: "IronGear", Count: 1}},
	}
	list = append(list, belt)

	potion := Recipe{
		ItemName:    "AutoSciencePack",
		Count:       1,
		Time:        5,
		Ingredients: []Ingredient{{ItemName: "IronGear", Count: 1}, {ItemName: "CopperPlate", Count: 1}},
	}
	list = append(list, potion)

	return list
}

func LoadJsonTest() {
	unstructuredJson := ReadFromFile("RecipeTest.txt")

	var result map[string]interface{}

	json.Unmarshal([]byte(unstructuredJson), &result)

	fmt.Println(result["ItemName"])
	fmt.Println(result)
}

func WriteToFile(fileName string, content string) {
	//Create file object, and check for error
	file, err := os.Create("./" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	//defer keyword triggers code when function ends
	defer file.Close()

	//Write content to file
	file.WriteString(content)
}

func ReadFromFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	return string(content)
}
