package examples

import (
	"net/http/httptest"
	"testing"

	"github.com/dojo-engineering/goponent-testing/pkg/goponent"
)

func Test_ExampleServer(t *testing.T) {
	tests := []goponent.Test{
		{
			Name: "Test hello",
			Steps: []goponent.Step{
				{
					Name: "Get request to /hello",
					Request: goponent.StringRequestGenerator{
						Method: "GET",
						Path:   "/hello",
					},
					Assertions: goponent.StringResponseAsserter{
						ExpectedBody:       "hello",
						ExpectedStatusCode: 200,
					},

					ContextSetters: nil,
				},
			},
		},
		{
			Name: "Unknown path returns 404",
			Steps: []goponent.Step{
				{
					Name: "Get request to /invalid",
					Request: goponent.StringRequestGenerator{
						Method: "GET",
						Path:   "/invalid",
					},
					Assertions: goponent.StringResponseAsserter{
						ExpectedBody:       "404 page not found\n",
						ExpectedStatusCode: 404,
					},

					ContextSetters: nil,
				},
			},
		},
	}

	handler := buildServer()
	server := httptest.NewServer(handler)
	defer server.Close()

	goponent.RunTests(t, tests, server.URL)

}
