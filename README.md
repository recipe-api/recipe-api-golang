# Recipe API - Go Starter

This is a Go starter kit for the [Recipe API](https://recipe-api.com). It includes a client library and example scripts to help you get started.

## Setup

1.  **Clone the repository:**
    ```bash
    git clone <your-repo-url>
    cd recipe-api-golang
    ```

2.  **Install dependencies:**
    ```bash
    go mod download
    ```

3.  **Configure API Key:**
    *   Copy `.env.example` to `.env`
    *   Get your API key from [recipe-api.com](https://recipe-api.com)
    *   Add it to `.env`: `RECIPE_API_KEY=rapi_...`

## Usage

Run the example scripts to explore the API:

*   **List dietary flags:**
    ```bash
    go run scripts/01-categories.go
    ```

*   **List cuisines:**
    ```bash
    go run scripts/02-cuisines.go
    ```

*   **Browse recipes:**
    ```bash
    go run scripts/03-browse.go
    go run scripts/03-browse.go --page=2
    ```

*   **Search recipes:**
    ```bash
    go run scripts/04-search.go --q="pasta"
    ```

*   **Filter recipes:**
    ```bash
    go run scripts/05-filter.go --cuisine="Italian" --difficulty="Beginner"
    ```

*   **Get full recipe details:**
    ```bash
    go run scripts/06-recipe.go --id=<recipe_id>
    ```

## Project Structure

*   `pkg/client/client.go`: API client configuration and request handling
*   `pkg/types/types.go`: Struct definitions for API responses
*   `pkg/utils/utils.go`: Helper functions for formatting output
*   `scripts/`: Example scripts demonstrating various endpoints
