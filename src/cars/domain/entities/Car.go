package entities

type Car struct {
	ID       int32   `json:"id"`
	Make     string  `json:"make"`
	Model    string  `json:"model"`
	Year     int32   `json:"year"`
	Mileage  int32   `json:"mileage"`
	FuelType string  `json:"fuel_type"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"`
	Color    string  `json:"color"`
}
