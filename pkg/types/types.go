package types

type Cuisine struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type DietaryFlag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type RecipeMeta struct {
	ActiveTime        string `json:"active_time"`
	PassiveTime       string `json:"passive_time"`
	TotalTime         string `json:"total_time"`
	OvernightRequired bool   `json:"overnight_required"`
	Yields            string `json:"yields"`
	YieldCount        int    `json:"yield_count"`
	ServingSizeG      int    `json:"serving_size_g"`
}

type DietaryInfo struct {
	Flags          []string `json:"flags"`
	NotSuitableFor []string `json:"not_suitable_for"`
}

type NutritionSummary struct {
	Calories      float64 `json:"calories"`
	ProteinG      float64 `json:"protein_g"`
	CarbohydratesG float64 `json:"carbohydrates_g"`
	FatG          float64 `json:"fat_g"`
}

type RecipeSummary struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Category         string           `json:"category"`
	Cuisine          string           `json:"cuisine"`
	Difficulty       string           `json:"difficulty"`
	Tags             []string         `json:"tags"`
	Meta             RecipeMeta       `json:"meta"`
	Dietary          DietaryInfo      `json:"dietary"`
	NutritionSummary NutritionSummary `json:"nutrition_summary"`
}

type ListMeta struct {
	Total       int  `json:"total"`
	Page        int  `json:"page"`
	PerPage     int  `json:"per_page"`
	TotalCapped bool `json:"total_capped,omitempty"`
}

type CuisinesResponse struct {
	Data []Cuisine `json:"data"`
}

type DietaryFlagsResponse struct {
	Data []DietaryFlag `json:"data"`
}

type RecipeListResponse struct {
	Data []RecipeSummary `json:"data"`
	Meta ListMeta        `json:"meta"`
}

type Equipment struct {
	Name        string  `json:"name"`
	Required    bool    `json:"required"`
	Alternative *string `json:"alternative"`
}

type IngredientItem struct {
	Name            string   `json:"name"`
	Quantity        any      `json:"quantity"` // Can be number or string
	Unit            *string  `json:"unit"`
	Preparation     *string  `json:"preparation"`
	Notes           *string  `json:"notes"`
	Substitutions   []string `json:"substitutions"`
	IngredientID    string   `json:"ingredient_id"`
	NutritionSource string   `json:"nutrition_source"`
}

type IngredientGroup struct {
	GroupName string           `json:"group_name"`
	Items     []IngredientItem `json:"items"`
}

type InstructionStructured struct {
	Action       string   `json:"action"`
	Temperature  *string  `json:"temperature"`
	Duration     *string  `json:"duration"`
	DonenessCues []string `json:"doneness_cues"`
}

type Instruction struct {
	StepNumber int                   `json:"step_number"`
	Phase      string                `json:"phase"`
	Text       string                `json:"text"`
	Structured InstructionStructured `json:"structured"`
	Tips       []string              `json:"tips"`
}

type StorageInfo struct {
	Refrigerator *struct {
		Notes    string `json:"notes"`
		Duration string `json:"duration"`
	} `json:"refrigerator,omitempty"`
	Freezer *struct {
		Duration *string `json:"duration"`
	} `json:"freezer,omitempty"`
	Reheating   *string `json:"reheating,omitempty"`
	DoesNotKeep bool    `json:"does_not_keep,omitempty"`
}

type FullNutrition struct {
	PerServing struct {
		Calories       float64  `json:"calories"`
		ProteinG       float64  `json:"protein_g"`
		CarbohydratesG float64  `json:"carbohydrates_g"`
		FatG           float64  `json:"fat_g"`
		FiberG         *float64 `json:"fiber_g,omitempty"`
		SugarG         *float64 `json:"sugar_g,omitempty"`
		SodiumMg       *float64 `json:"sodium_mg,omitempty"`
		SaturatedFatG  *float64 `json:"saturated_fat_g,omitempty"`
		CholesterolMg  *float64 `json:"cholesterol_mg,omitempty"`
	} `json:"per_serving"`
	Sources []string `json:"sources"`
}

type RecipeDetail struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Category        string            `json:"category"`
	Cuisine         string            `json:"cuisine"`
	Difficulty      string            `json:"difficulty"`
	Tags            []string          `json:"tags"`
	Meta            RecipeMeta        `json:"meta"`
	Dietary         DietaryInfo       `json:"dietary"`
	Nutrition       FullNutrition     `json:"nutrition"`
	Storage         *StorageInfo      `json:"storage,omitempty"`
	Equipment       []Equipment       `json:"equipment"`
	Ingredients     []IngredientGroup `json:"ingredients"`
	Instructions    []Instruction     `json:"instructions"`
	Troubleshooting map[string]string `json:"troubleshooting,omitempty"`
	ChefNotes       []string          `json:"chef_notes,omitempty"`
	CulturalContext *string           `json:"cultural_context,omitempty"`
}

type Usage struct {
	MonthlyRemaining int `json:"monthly_remaining"`
	MonthlyLimit     int `json:"monthly_limit"`
	DailyRemaining   int `json:"daily_remaining"`
	DailyLimit       int `json:"daily_limit"`
}

type RecipeResponse struct {
	Data  RecipeDetail `json:"data"`
	Usage *Usage       `json:"usage,omitempty"`
}

type ApiError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Ingredient category (from /api/v1/ingredient-categories)
type IngredientCategory struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type IngredientCategoriesResponse struct {
	Data []IngredientCategory `json:"data"`
}

// Ingredient (from /api/v1/ingredients)
type Ingredient struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Source   string `json:"source"`
}

type IngredientsResponse struct {
	Data []Ingredient `json:"data"`
	Meta ListMeta     `json:"meta"`
}
