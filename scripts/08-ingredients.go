package main

import (
	"flag"
	"fmt"
	"recipe-api-golang/pkg/client"
	"recipe-api-golang/pkg/types"
	"recipe-api-golang/pkg/utils"
)

func main() {
	q := flag.String("q", "", "Search by ingredient name")
	category := flag.String("category", "", "Filter by category")
	page := flag.Int("page", 1, "Page number")
	perPage := flag.Int("per_page", 20, "Results per page")
	flag.Parse()

	utils.Header("Browse Ingredients")

	if *q != "" {
		fmt.Printf("Search: \"%s\"\n", *q)
	}
	if *category != "" {
		fmt.Printf("Category: %s\n", *category)
	}
	fmt.Println()

	params := map[string]string{
		"page":     fmt.Sprintf("%d", *page),
		"per_page": fmt.Sprintf("%d", *perPage),
	}
	if *q != "" {
		params["q"] = *q
	}
	if *category != "" {
		params["category"] = *category
	}

	resp, err := client.ApiRequest[types.IngredientsResponse]("/api/v1/ingredients", params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Found %d ingredients (page %d):\n\n", resp.Meta.Total, resp.Meta.Page)

	for _, ingredient := range resp.Data {
		fmt.Printf("  %s\n", ingredient.Name)
		fmt.Printf("    ID: %s\n", ingredient.ID)
		fmt.Printf("    Category: %s\n", ingredient.Category)
		fmt.Printf("    Source: %s\n", ingredient.Source)
		fmt.Println()
	}

	utils.Divider()

	fmt.Println("\nUsage examples:")
	fmt.Println("  go run scripts/08-ingredients.go --q=\"chicken\"")
	fmt.Println("  go run scripts/08-ingredients.go --category=\"Vegetables\"")
	fmt.Println("  go run scripts/08-ingredients.go --page=2")
	fmt.Println("\nUse ingredient IDs to filter recipes:")
	fmt.Println("  go run scripts/05-filter.go --ingredients=\"<id1>,<id2>\"\n")
}
