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
	page := flag.Int("page", 1, "Page number")
	perPage := flag.Int("per_page", 10, "Items per page")
	flag.Parse()

	utils.Header("Browse Recipes")

	params := map[string]string{
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

	totalPages := int(math.Ceil(float64(meta.Total) / float64(meta.PerPage)))
	fmt.Printf("Page %d of %d (%d total recipes)\n\n", meta.Page, totalPages, meta.Total)

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

	fmt.Println("\n>> Tips:")
	fmt.Println("   * Browse more: go run scripts/03-browse.go --page=2")
	fmt.Println("   * Search: go run scripts/04-search.go --q=\"pasta\"")
	fmt.Println("   * Get full recipe: go run scripts/06-recipe.go --id=<recipe_id>\n")
}
