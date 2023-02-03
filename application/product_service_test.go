package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/petruspierre/ports-and-adapters/application"
	mock_application "github.com/petruspierre/ports-and-adapters/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	productPersistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: productPersistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	productPersistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	productPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: productPersistence,
	}

	result, err := service.Create("Product 1", 10.0)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
