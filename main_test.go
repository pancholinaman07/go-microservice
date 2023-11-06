package main

import (
	"fmt"
	"go-microservice/product-api/sdk/client/products"
	"product-api/sdk/client"
	"testing"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("%#v", prod.GetPayload()[0])
}
