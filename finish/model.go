package finish

type Finish struct {
	ID      int `json:"id"`
	Finish 	string `json:"finish"`
	Stain   bool   `json:"stain"`
	Price float32 `json:"price"`

}
