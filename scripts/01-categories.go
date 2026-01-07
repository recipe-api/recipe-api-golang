package main

import (
	"fmt"
	"recipe-api-golang/pkg/client"
	"recipe-api-golang/pkg/types"
	"recipe-api-golang/pkg/utils"
	"sort"
)

func main() {
	utils.Header("Dietary Flags")

	resp, err := client.ApiRequest[types.DietaryFlagsResponse]("/api/v1/dietary-flags", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data := resp.Data
	fmt.Printf("Found %d dietary options:\n\n", len(data))

	sort.Slice(data, func(i, j int) bool {
		return data[i].Count > data[j].Count
	})

	limit := 20
	if len(data) < limit {
		limit = len(data)
	}

	for _, flag := range data[:limit] {
		fmt.Printf("  * %s (%d recipes)\n", flag.Name, flag.Count)
	}

	if len(data) > 20 {
		fmt.Printf("  ... and %d more\n", len(data)-20)
	}

	utils.Divider()
	fmt.Println("\n>> Next step: Run `go run scripts/02-cuisines.go` to see available cuisines\n")
}
