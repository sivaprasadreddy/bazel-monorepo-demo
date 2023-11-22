package products_test

import (
	"github.com/sivaprasadreddy/products"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAll(t *testing.T) {
	expected := 3
	actual := len(products.GetAll())
	require.Equal(t, expected, actual)

}
