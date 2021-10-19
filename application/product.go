package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true) // iniciei o validador
}

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
	ENABLE_ERROR            = "The price must be greater than 0 to enabled the product"
	DISABLE_ERROR           = "The price must be 0 to be disabled"
	VALIDATION_STATUS_ERROR = "The status must be enable or disabled"
	VALIDATION_PRICE_VALUE  = "The value must be greater or igual to 0"
)

type Product struct {
	ID     string  `valid:"uuidv4"` // tipo uma anotation
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
}

// implementation
func (p *Product) IsValid() (bool, error) {
	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New(VALIDATION_STATUS_ERROR)
	}
	if p.Price < 0 {
		return false, errors.New(VALIDATION_PRICE_VALUE)
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New(ENABLE_ERROR)
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New(DISABLE_ERROR)
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
