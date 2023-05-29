package product

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type CreateProductPayload struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type UpdateProductRequest struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}
