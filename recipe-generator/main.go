				fmt.Printf("Usage: --delete <filename>\n")
package main

import (
	"fmt"
	"os"
	"recipe-generator/file"
	"strings"
)

func main() {
	workingDir := "."
	for i, arg := range os.Args {
		if arg == "--dir" || arg == "--directory" {
			if i+1 < len(os.Args) {
				workingDir = os.Args[i+1]
				os.Args = append(os.Args[:i], os.Args[i+2:]...)
				if _, err := os.Stat(workingDir); os.IsNotExist(err) {
					if err := os.MkdirAll(workingDir, 0755); err != nil {
						fmt.Printf("Failed to create directory %s: %v\n", workingDir, err)
						os.Exit(1)
					}
				}
				break
			} else {
				fmt.Println("Missing directory after", arg)
				os.Exit(1)
			}
		}
	}
	if err := os.Chdir(workingDir); err != nil {
		fmt.Printf("Failed to change directory to %s: %v\n", workingDir, err)
		os.Exit(1)
	}
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "--") {
		switch os.Args[1] {
		case "--create":
			if len(os.Args) < 5 {
				fmt.Printf("Usage: --create <craftingType> <result> <count> <itemSlots>\n")
				os.Exit(1)
			}
			craftingType := strings.ToUpper(os.Args[2])
			result := os.Args[3]
			count := 0
			_, err := fmt.Sscanf(os.Args[4], "%d", &count)
			if err != nil {
				fmt.Println("Invalid count:", os.Args[4])
				os.Exit(1)
			}
			itemSlots := make(map[string][]int)
			itemCounts := make(map[string]int)
			itemPairs := os.Args[5:]
			switch craftingType {
			case "SHAPED_CRAFTING":
				for _, pair := range itemPairs {
					openIdx := strings.Index(pair, "[")
					closeIdx := strings.LastIndex(pair, "]")
					if openIdx == -1 || closeIdx == -1 || closeIdx < openIdx {
						fmt.Printf("Invalid itemSlot format: %s\n", pair)
						os.Exit(1)
					}
					id := strings.TrimSpace(pair[:openIdx])
					slotsStr := pair[openIdx+1 : closeIdx]
					slotParts := strings.Split(slotsStr, ",")
					var slots []int
					for _, s := range slotParts {
						s = strings.TrimSpace(s)
						if s == "" {
							continue
						}
						var num int
						_, err := fmt.Sscanf(s, "%d", &num)
						if err != nil {
							fmt.Printf("Invalid slot number: %s\n", s)
							os.Exit(1)
						}
						slots = append(slots, num)
					}
					itemSlots[id] = slots
				}
				fmt.Println(itemSlots)
				file.CreateShaped(result, count, itemSlots)
			case "SHAPELESS_CRAFTING":
				for _, pair := range itemPairs {
					parts := strings.SplitN(pair, "=", 2)
					if len(parts) != 2 {
						fmt.Printf("Invalid itemCount format: %s\n", pair)
						os.Exit(1)
					}
					id := strings.TrimSpace(parts[0])
					var cnt int
					_, err := fmt.Sscanf(parts[1], "%d", &cnt)
					if err != nil {
						fmt.Printf("Invalid count: %s\n", parts[1])
						os.Exit(1)
					}
					itemCounts[id] = cnt
				}
				file.CreateShapeless(result, count, itemCounts)
				fmt.Println("Create recipe for " + result)
			default:
				fmt.Printf("Unknown crafting type: %s\n", craftingType)
				os.Exit(1)
			}
			os.Exit(0)
		case "--delete":
			if len(os.Args) < 3 {
				fmt.Printf("Usage: --delete <filename>\n")
				os.Exit(1)
			}
			filename := os.Args[2]
			file.DeleteRecipe(filename)
			fmt.Println("Deleted recipe for " + filename)
		default:
			fmt.Printf("Unknown flag: %s\n", os.Args[1])
			os.Exit(1)
		}
	}
	fmt.Println("No valid flag provided. Use '--create' or '--delete' for recipe operations.")
	fmt.Println("You can also use '--dir <directory>' or '--directory <directory>' to set the working directory.")
	os.Exit(0)
}
