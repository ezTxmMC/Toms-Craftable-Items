package file

import (
	"fmt"
	"os"
)

func DeleteRecipe(filename string) {
	recipeFile := filename + ".json"
	if err := os.Remove(recipeFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s does not exist.\n", recipeFile)
		} else {
			fmt.Printf("Failed to delete file %s: %v\n", recipeFile, err)
		}
		os.Exit(1)
	}
	fmt.Printf("Deleted file: %s\n", recipeFile)
	os.Exit(0)
}
