package material
type Material struct {
	ID    int `json:"id"`
	Wood  int `json:"wood"`
	Board int `json:"board"`
	Price float32 `json:"price"`
}

type PriceUpdateMaterial struct {
	ID    int `json:"id"`
	Price float32 `json:"price"`
}
