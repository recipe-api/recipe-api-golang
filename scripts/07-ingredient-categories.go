package main

import (
	"fmt"
	"recipe-api-golang/pkg/client"
	"recipe-api-golang/pkg/types"
	"recipe-api-golang/pkg/utils"
	"sort"
)

func main() {
	utils.Header("Ingredient Categories")

	resp, err := client.ApiRequest[types.IngredientCategoriesResponse]("/api/v1/ingredient-categories", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data := resp.Data
	fmt.Printf("Found %d ingredient categories:\n\n", len(data))

	sort.Slice(data, func(i, j int) bool {
		return data[i].Count > data[j].Count
	})

	limit := 20
	if len(data) < limit {
		limit = len(data)
	}

	for _, category := range data[:limit] {
		fmt.Printf("  * %s (%d ingredients)\n", category.Name, category.Count)
	}

	if len(data) > 20 {
		fmt.Printf("  ... and %d more\n", len(data)-20)
	}

	utils.Divider()
	fmt.Println("\n>> Next step: Run `go run scripts/08-ingredients.go` to browse ingredients\n")
}
