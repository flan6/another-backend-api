package product

import "api/pkg/entity"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

func NewProduct(p entity.Product) Product {
	return Product{
		ID:    p.ID,
		Name:  p.Name,
		Value: float32(p.Value) / 100,
	}
}

func GetProduct(p Product) entity.Product {
	return entity.Product{
		ID:    p.ID,
		Name:  p.Name,
		Value: int(p.Value),
	}
}

func (p Product) Validate() bool {
	if p.ID < 0 {
		return false
	}

	if p.Name == "" {
		return false
	}

	if p.Value < 0 {
		return false
	}

	return true
}

func (p *Product) Empty() bool {
	return p == nil || *p == Product{}
}
