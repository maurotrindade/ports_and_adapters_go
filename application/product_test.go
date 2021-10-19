package application_test

import (
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
	require.Equal(t, application.ENABLE_ERROR, err.Error())
}

func TestProduct_Disable(t *testing.T) {
	sut := makeSut()
	err := sut.Disable()
	require.Equal(t, application.DISABLE_ERROR, err.Error())

	sut.Price = 0
	err = sut.Disable()
	require.Nil(t, err)
}
