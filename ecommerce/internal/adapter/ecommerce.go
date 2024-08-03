package adapter

type IEcommerceAdapter interface {
}

type EcommerceAdapter struct {
}

func NewEcommerceAdapter() IEcommerceAdapter {
	return &EcommerceAdapter{}
}

func (a *EcommerceAdapter) GetProducts(param interface{}) (interface{}, error) {
	return nil, nil
}
