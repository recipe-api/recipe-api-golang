package main

import (
	"flag"
	"fmt"
	"recipe-api-golang/pkg/client"
	"recipe-api-golang/pkg/types"
	"recipe-api-golang/pkg/utils"
	"strings"
)

func main() {
	id := flag.String("id", "", "Recipe ID")
	flag.Parse()

	if *id == "" {
		fmt.Println("\nGet full recipe details\n")
		utils.Warning("!! Note: This endpoint costs 1 credit per request !!\n")
		fmt.Println("Usage: go run scripts/06-recipe.go --id=<recipe_id>\n")
		fmt.Println("To find recipe IDs:")
		fmt.Println("  1. Run `go run scripts/03-browse.go` or `go run scripts/04-search.go`")
		fmt.Println("  2. Copy the ID from a recipe you want\n")
		return
	}

	utils.Warning("\n!! Fetching full recipe (costs 1 credit) ...\n")

	resp, err := client.ApiRequest[types.RecipeResponse](fmt.Sprintf("/api/v1/recipes/%s", *id), nil)
	if err != nil {
		if apiErr, ok := err.(*client.RecipeApiError); ok && apiErr.Status == 404 {
			fmt.Printf("\n[X] Recipe not found: %s\n\n", *id)
			fmt.Println("Make sure the ID is correct. Find IDs with:")
			fmt.Println("  go run scripts/03-browse.go")
			fmt.Println("  go run scripts/04-search.go --q=\"...\"\n")
			return
		}
		fmt.Println("Error:", err)
		return
	}

	recipe := resp.Data
	usage := resp.Usage

	utils.Header(recipe.Name)

	fmt.Println(recipe.Description)
	fmt.Println()

	// Overview
	utils.Label("Category", fmt.Sprintf("%s | %s", recipe.Category, recipe.Cuisine))
	utils.Label("Difficulty", recipe.Difficulty)
	utils.Label("Active time", utils.FormatDuration(recipe.Meta.ActiveTime))
	utils.Label("Passive time", utils.FormatDuration(recipe.Meta.PassiveTime))
	utils.Label("Total time", utils.FormatDuration(recipe.Meta.TotalTime))
	utils.Label("Yields", recipe.Meta.Yields)

	if len(recipe.Dietary.Flags) > 0 {
		utils.Label("Dietary", strings.Join(recipe.Dietary.Flags, ", "))
	}
	if recipe.Meta.OvernightRequired {
		utils.Warning("  ** Requires overnight preparation **")
	}

	// Nutrition
	nutrition := recipe.Nutrition.PerServing
	fmt.Println()
	utils.Subheader("Nutrition (per serving)")
	utils.Label("Calories", fmt.Sprintf("%.0f kcal", nutrition.Calories))
	utils.Label("Protein", fmt.Sprintf("%.0f g", nutrition.ProteinG))
	utils.Label("Carbs", fmt.Sprintf("%.0f g", nutrition.CarbohydratesG))
	utils.Label("Fat", fmt.Sprintf("%.0f g", nutrition.FatG))
	if nutrition.FiberG != nil {
		utils.Label("Fiber", fmt.Sprintf("%.0f g", *nutrition.FiberG))
	}

	// Equipment
	if len(recipe.Equipment) > 0 {
		fmt.Println()
		utils.Subheader("Equipment")
		for _, item := range recipe.Equipment {
			alt := ""
			if item.Alternative != nil {
				alt = fmt.Sprintf(" (or: %s)", *item.Alternative)
			}
			req := ""
			if !item.Required {
				req = " [optional]"
			}
			fmt.Printf("  * %s%s%s\n", item.Name, alt, req)
		}
	}

	// Ingredients
	fmt.Println()
	utils.Subheader("Ingredients")
	for _, group := range recipe.Ingredients {
		if group.GroupName != "" {
			fmt.Printf("\n  [%s]\n", group.GroupName)
		}
		for _, ing := range group.Items {
			amount := ""
			if ing.Unit != nil {
				amount = fmt.Sprintf("%v %s", ing.Quantity, *ing.Unit)
			} else {
				amount = fmt.Sprintf("%v", ing.Quantity)
			}
			prep := ""
			if ing.Preparation != nil {
				prep = fmt.Sprintf(", %s", *ing.Preparation)
			}
			notes := ""
			if ing.Notes != nil {
				notes = fmt.Sprintf(" (%s)", *ing.Notes)
			}
			fmt.Printf("  * %s %s%s%s\n", amount, ing.Name, prep, notes)
		}
	}

	// Instructions
	fmt.Println()
	utils.Subheader("Instructions")
	for _, step := range recipe.Instructions {
		duration := ""
		if step.Structured.Duration != nil {
			duration = fmt.Sprintf(" %s", utils.Highlight(fmt.Sprintf("[%s]", utils.FormatDuration(*step.Structured.Duration))))
		}

		fmt.Printf("\n  %d. [%s] %s%s\n", step.StepNumber, step.Phase, step.Text, duration)

		if len(step.Tips) > 0 {
			for _, tip := range step.Tips {
				fmt.Printf("     >> %s\n", tip)
			}
		}
	}

	// Chef notes
	if len(recipe.ChefNotes) > 0 {
		fmt.Println()
		utils.Subheader("Chef Notes")
		for _, note := range recipe.ChefNotes {
			fmt.Printf("  * %s\n", note)
		}
	}

	// Cultural context
	if recipe.CulturalContext != nil {
		fmt.Println()
		utils.Subheader("About This Dish")
		fmt.Printf("  %s\n", *recipe.CulturalContext)
	}

	// Storage
	if recipe.Storage != nil {
		fmt.Println()
		utils.Subheader("Storage")
		if recipe.Storage.DoesNotKeep {
			fmt.Println("  Best eaten immediately.")
		}
		if recipe.Storage.Refrigerator != nil {
			val := recipe.Storage.Refrigerator.Notes
			if val == "" {
				val = recipe.Storage.Refrigerator.Duration
			}
			fmt.Printf("  Refrigerator: %s\n", val)
		}
		if recipe.Storage.Reheating != nil {
			fmt.Printf("  Reheating: %s\n", *recipe.Storage.Reheating)
		}
	}

	// Usage info
	if usage != nil {
		fmt.Println("\n--- API Usage ---")
		// Comma formatting for thousands is annoying in Go standard lib, using simple output
		// Or using a helper/printer. I'll just print %d for simplicity to avoid huge dependencies or complex logic
		// Or implement a simple comma helper.
		fmt.Printf("Monthly remaining: %d\n", usage.MonthlyRemaining)
		fmt.Printf("Daily remaining:   %d\n", usage.DailyRemaining)
	}

	fmt.Println()
}
