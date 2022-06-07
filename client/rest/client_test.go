package rest_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/client/rest"
)

func TestClient(t *testing.T) {
	c := rest.NewRESTClient()

	err := c.Get("").
		Do(context.Background()).
		Into(nil)
	if err != nil {
		t.Fatal(err)
	}
}
