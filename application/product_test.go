package application_test

import (
	"errors"
	"testing"

	"github.com/maurotrindade/ports_and_adapters_go/application"
	"github.com/stretchr/testify/require"
)

func makeSut() application.Product {
	return application.Product{
		ID:     "A2S3D4",
		Name:   "Teste",
		Status: application.DISABLED,
		Price:  10,
	}
}

func TestProduct_Enabled(t *testing.T) {
	sut := makeSut()
	err := sut.Enable()
	require.Nil(t, err)

	sut.Price = 0
	err = sut.Enable()
	require.Error(t, errors.New(application.ENABLE_ERROR))
}

func TestProduct_Disable(t *testing.T) {
	sut := makeSut()
	err := sut.Disable()
	require.Error(t, errors.New(application.DISABLE_ERROR))

	sut.Price = 0
	err = sut.Disable()
	require.Nil(t, err)
}
