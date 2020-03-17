package products

type Product struct {
	ID       int  `json:"id"`
	Name 	 string `json:"name"`
	Materials [] *ProdMat
	Finish   int  `json:"finish"`
	Hours    float32 `json:"hours"`
	Fees     string `json:"fees"`
	Total    string `json:"total"`
}

type ProdMat struct {
ID       int `json:"id"`
Board 	 int `json:"board"`
Wood     int `json:"wood"`
Quantity int `json:"quantity"`
Price    float32 `json:"price"`
}
