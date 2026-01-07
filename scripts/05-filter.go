package main

import (
	"flag"
	"fmt"
	"recipe-api-golang/pkg/client"
	"recipe-api-golang/pkg/types"
	"recipe-api-golang/pkg/utils"
	"strconv"
	"strings"
)

func main() {
	category := flag.String("category", "", "Recipe category")
	cuisine := flag.String("cuisine", "", "Cuisine type")
	difficulty := flag.String("difficulty", "", "Difficulty level")
	dietary := flag.String("dietary", "", "Dietary preference")
	maxCalories := flag.Int("max_calories", 0, "Maximum calories")
	minProtein := flag.Int("min_protein", 0, "Minimum protein")
	page := flag.Int("page", 1, "Page number")
	perPage := flag.Int("per_page", 10, "Items per page")
	flag.Parse()

hasFilters := *category != "" || *cuisine != "" || *difficulty != "" || *dietary != "" || *maxCalories > 0 || *minProtein > 0

	if !hasFilters {
		fmt.Println("\nFilter recipes by multiple criteria\n")
		fmt.Println("Available filters:")
		fmt.Println("  --category     Recipe category (Breakfast, Main, Dessert, etc.)")
		fmt.Println("  --cuisine      Cuisine type (run `go run scripts/02-cuisines.go` for list)")
		fmt.Println("  --difficulty   Beginner, Intermediate, or Advanced")
		fmt.Println("  --dietary      Vegetarian, Vegan, Gluten-Free, etc. (run `go run scripts/01-categories.go`)")
		fmt.Println("  --max_calories Maximum calories per serving")
		fmt.Println("  --min_protein  Minimum protein in grams\n")
		fmt.Println("Examples:")
		fmt.Println("  go run scripts/05-filter.go --cuisine=\"Italian\" --difficulty=\"Beginner\"")
		fmt.Println("  go run scripts/05-filter.go --dietary=\"Vegan\" --max_calories=400")
		fmt.Println("  go run scripts/05-filter.go --category=\"Dessert\" --cuisine=\"French\"\n")
		return
	}

	var filters []string
	params := map[string]string{
		"page":     strconv.Itoa(*page),
		"per_page": strconv.Itoa(*perPage),
	}

	if *category != "" {
		params["category"] = *category
		filters = append(filters, "category="+*category)
	}
	if *cuisine != "" {
		params["cuisine"] = *cuisine
		filters = append(filters, "cuisine="+*cuisine)
	}
	if *difficulty != "" {
		params["difficulty"] = *difficulty
		filters = append(filters, "difficulty="+*difficulty)
	}
	if *dietary != "" {
		params["dietary"] = *dietary
		filters = append(filters, "dietary="+*dietary)
	}
	if *maxCalories > 0 {
		params["max_calories"] = strconv.Itoa(*maxCalories)
		filters = append(filters, fmt.Sprintf("max_calories=%d", *maxCalories))
	}
	if *minProtein > 0 {
		params["min_protein"] = strconv.Itoa(*minProtein)
		filters = append(filters, fmt.Sprintf("min_protein=%d", *minProtein))
	}

	utils.Header("Filtered Recipes")
	fmt.Printf("Filters: %s\n\n", strings.Join(filters, ", "))

	resp, err := client.ApiRequest[types.RecipeListResponse]("/api/v1/recipes", params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	recipes := resp.Data
	meta := resp.Meta

	if len(recipes) == 0 {
		fmt.Println("No recipes match your filters.\n")
		fmt.Println("Try relaxing some criteria.\n")
		return
	}

	fmt.Printf("Found %d matching recipes\n\n", meta.Total)

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

	fmt.Println("\n>> Get full recipe: go run scripts/06-recipe.go --id=<recipe_id>\n")
}
