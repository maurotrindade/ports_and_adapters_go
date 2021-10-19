package application

import "errors"

type ProductProtocol interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

const (
	ENABLE_ERROR = "The price must be greater than 0 to enabled the product"
)

type Product struct {
	ID     string
	Name   string
	Status string
	Price  float64
}

// implementation
func (p *Product) IsValid() (bool, error) {
	return false, errors.New("The price must be greater than 0 to enabled the product")
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New(ENABLE_ERROR)
}

func (p *Product) Disable() error {
	return errors.New("The price must be greater than 0 to enabled the product")
}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
