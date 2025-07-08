package file

import (
	"encoding/json"
	"os"
	"strings"
)

func CreateShapeless(result string, count int, slots map[string]int) {
	type Ingredient struct {
		Item  string `json:"item"`
		Count int    `json:"count"`
	}
	type Result struct {
		ID    string `json:"id"`
		Count int    `json:"count"`
	}
	type ShapelessRecipe struct {
		Type        string       `json:"type"`
		Ingredients []Ingredient `json:"ingredients"`
		Result      Result       `json:"result"`
	}

	ingredients := []Ingredient{}
	for item, c := range slots {
		ingredients = append(ingredients, Ingredient{Item: item, Count: c})
	}

	config := ShapelessRecipe{
		Type:        "minecraft:crafting_shapeless",
		Ingredients: ingredients,
		Result:      Result{ID: result, Count: count},
	}

	file, err := os.Create(strings.Split(result, ":")[1] + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(config); err != nil {
		panic(err)
	}
}

func CreateShaped(result string, count int, slots map[string][]int) {
	letters := []string{"#", "A", "B", "C", "D", "E", "F", "G", "H"}
	itemKeys := make(map[string]string)
	i := 0
	for item := range slots {
		if i < len(letters) {
			itemKeys[item] = letters[i]
			i++
		}
	}
	grid := make([][]string, 3)
	for i := range grid {
		grid[i] = make([]string, 3)
	}
	for item, positions := range slots {
		letter := itemKeys[item]
		for _, pos := range positions {
			if pos >= 0 && pos < 9 {
				row := pos / 3
				col := pos % 3
				grid[row][col] = letter
			}
		}
	}
	pattern := make([]string, 3)
	for row := range 3 {
		line := ""
		for col := range 3 {
			if grid[row][col] != "" {
				line += grid[row][col]
			} else {
				line += " "
			}
		}
		pattern[row] = line
	}
	key := make(map[string]map[string]string)
	for item, letter := range itemKeys {
		key[letter] = map[string]string{"item": item}
	}
	type Result struct {
		ID    string `json:"id"`
		Count int    `json:"count"`
	}
	type ShapedRecipe struct {
		Type    string                       `json:"type"`
		Pattern []string                     `json:"pattern"`
		Key     map[string]map[string]string `json:"key"`
		Result  Result                       `json:"result"`
	}
	config := ShapedRecipe{
		Type:    "minecraft:crafting_shaped",
		Pattern: pattern,
		Key:     key,
		Result:  Result{ID: result, Count: count},
	}
	file, err := os.Create(strings.Split(result, ":")[1] + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(config); err != nil {
		panic(err)
	}
}
