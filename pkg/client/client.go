package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const BaseURL = "https://recipe-api.com"

func init() {
	// Load .env file if it exists
	_ = godotenv.Load()
}

type RecipeApiError struct {
	Message string
	Status  int
	Code    string
}

func (e *RecipeApiError) Error() string {
	return fmt.Sprintf("%s (Status: %d, Code: %s)", e.Message, e.Status, e.Code)
}

func getApiKey() string {
	key := os.Getenv("RECIPE_API_KEY")

	if key == "" {
		fmt.Println("\n[X] Missing API key!")
		fmt.Println("To fix this:")
		fmt.Println("  1. Copy .env.example to .env")
		fmt.Println("  2. Add your API key from https://recipe-api.com\n")
		os.Exit(1)
	}

	if !strings.HasPrefix(key, "rapi_") {
		fmt.Println("\n[X] Invalid API key format!")
		fmt.Println("API keys should start with \"rapi_\"")
		fmt.Println("Get your key from https://recipe-api.com\n")
		os.Exit(1)
	}

	return key
}

func ApiRequest[T any](endpoint string, params map[string]string) (*T, error) {
	apiKey := getApiKey()

	u, err := url.Parse(BaseURL + endpoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("\n[X] Network error!")
		fmt.Println("Could not connect to the API. Please check:")
		fmt.Println("  - Your internet connection")
		fmt.Println("  - The API status at https://recipe-api.com\n")
		os.Exit(1)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		if resp.StatusCode == 404 {
			return nil, &RecipeApiError{Message: "Resource not found", Status: 404, Code: "NOT_FOUND"}
		}
		handleErrorResponse(resp)
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func handleErrorResponse(resp *http.Response) {
	status := resp.StatusCode

	switch status {
	case 401:
		fmt.Println("\n[X] Authentication failed!")
		fmt.Println("Your API key was rejected. Please check:")
		fmt.Println("  - The key is copied correctly (no extra spaces)")
		fmt.Println("  - The key is active in your dashboard\n")
		os.Exit(1)
	case 403:
		fmt.Println("\n[X] Access denied!")
		fmt.Println("Your account may not have access to this endpoint.")
		fmt.Println("Check your plan limits at https://recipe-api.com\n")
		os.Exit(1)
	case 404:
		// Let the caller handle 404 if they want, but usually we exit or panic in this simple client.
		// However, to match TS, we should probably throw an error.
		// But in this helper, we are exiting for most errors.
		// For 404, the TS code throws.
		// Since we can't throw, we'll need to handle it in ApiRequest or return error.
		// I'll handle it here by printing and exiting unless it's handled upstream?
		// Actually, let's just print error and exit for now to keep it simple, 
		// BUT 06-recipe.ts catches 404 specifically.
		// So we should verify 404 handling.
		// I'll change ApiRequest to check 404 specifically before calling handleErrorResponse?
		// Or handleErrorResponse can panic/exit for everything EXCEPT 404?
		return // Return so ApiRequest can return the error
	case 429:
		fmt.Println("\n[X] Rate limit exceeded!")
		fmt.Println("You have exceeded your API limits.")
		fmt.Println("Check your remaining quota in the dashboard.\n")
		os.Exit(1)
	default:
		fmt.Printf("\n[X] API error (%d)!\n", status)
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Response: %s\n", string(body)[:min(len(body), 500)])
		os.Exit(1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
