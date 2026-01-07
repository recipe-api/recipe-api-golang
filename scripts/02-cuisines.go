package main

import (
	"fmt"
	"recipe-api-golang/pkg/client"
	"recipe-api-golang/pkg/types"
	"recipe-api-golang/pkg/utils"
	"sort"
	"strings"
)

func main() {
	utils.Header("Cuisines")

	resp, err := client.ApiRequest[types.CuisinesResponse]("/api/v1/cuisines", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data := resp.Data
	fmt.Printf("Found %d cuisines:\n\n", len(data))

	sort.Slice(data, func(i, j int) bool {
		return data[i].Count > data[j].Count
	})

	for _, cuisine := range data {
		count := cuisine.Count
		barLen := count / 20
		if barLen > 20 {
			barLen = 20
		}
		bar := ""
		for k := 0; k < barLen; k++ {
			bar += "█"
		}
		
		// Pad the name to 20 characters
		padding := 20 - len(cuisine.Name)
		if padding < 0 {
			padding = 0
		}
		paddedName := cuisine.Name + strings.Repeat(" ", padding)
		
		fmt.Printf("  %s %s %d\n", paddedName, bar, count)
	}

	utils.Divider()
	fmt.Println("\n>> Next step: Run `go run scripts/03-browse.go` to see recipes\n")
}
