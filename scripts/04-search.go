package main

import (
	"flag"
	"fmt"
	"math"
	"recipe-api-golang/pkg/client"
	"recipe-api-golang/pkg/types"
	"recipe-api-golang/pkg/utils"
	"strconv"
	"strings"
)

func main() {
	query := flag.String("q", "", "Search query")
	page := flag.Int("page", 1, "Page number")
	perPage := flag.Int("per_page", 10, "Items per page")
	flag.Parse()

	if *query == "" {
		fmt.Println("Please provide a search query with --q")
		return
	}

	utils.Header(fmt.Sprintf("Search: \"%s\"", *query))

	params := map[string]string{
		"q":        *query,
		"page":     strconv.Itoa(*page),
		"per_page": strconv.Itoa(*perPage),
	}

	resp, err := client.ApiRequest[types.RecipeListResponse]("/api/v1/recipes", params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	recipes := resp.Data
	meta := resp.Meta

	if len(recipes) == 0 {
		fmt.Println("No recipes found matching your search.\n")
		fmt.Println("Try:")
		fmt.Println("  * Different keywords")
		fmt.Println("  * Broader terms")
		fmt.Println("  * Check spelling\n")
		return
	}

	totalPages := int(math.Ceil(float64(meta.Total) / float64(meta.PerPage)))
	fmt.Printf("Page %d of %d (%d matching recipes)\n\n", meta.Page, totalPages, meta.Total)

	for _, recipe := range recipes {
		fmt.Println(utils.Highlight(recipe.Name))
		utils.Label("ID", recipe.ID)
		utils.Label("Category", fmt.Sprintf("%s | %s", recipe.Category, recipe.Cuisine))
		utils.Label("Difficulty", recipe.Difficulty)
		utils.Label("Time", utils.FormatDuration(recipe.Meta.TotalTime))

		if len(recipe.Dietary.Flags) > 0 {
			limit := 4
			if len(recipe.Dietary.Flags) < limit {
				limit = len(recipe.Dietary.Flags)
			}
			utils.Label("Dietary", strings.Join(recipe.Dietary.Flags[:limit], ", "))
		}
		utils.Label("Calories", fmt.Sprintf("%.0f kcal", recipe.NutritionSummary.Calories))
		fmt.Printf("  %s\n", utils.Truncate(recipe.Description, 80))
		utils.Divider()
	}
}
