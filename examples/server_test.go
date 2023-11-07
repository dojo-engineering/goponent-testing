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

					Act: goponent.HttpRequestAction{
						Method:      "GET",
						Path:        "/hello",
						ContentType: "text/plain",
					},
					Assertions: goponent.HttpResponseAsserter{
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
					Act: goponent.HttpRequestAction{
						Method: "GET",
						Path:   "/invalid",
					},
					Assertions: goponent.HttpResponseAsserter{
						ExpectedBody:       "404 page not found\n",
						ExpectedStatusCode: 404,
					},

					ContextSetters: nil,
				},
			},
		},
		{
			Name: "Post to create car endpoint creates a new car",
			Steps: []goponent.Step{
				{
					Name: "Post to car endpoint returns 200",
					Act: goponent.JsonRequestAction[Car]{
						Method: "POST",
						Body: Car{
							Make:  "Subaru",
							Model: "Outback",
						},
						Path: "/car",
					},
					Assertions: goponent.JsonResponseAsserter[Car]{
						ExpectedBody:       Car{Make: "Subaru", Model: "Outback", Id: "1"},
						ExpectedStatusCode: 200,
					},

					ContextSetters: nil,
				},
				{
					Name: "Get car returns 200 and car",
					Act: goponent.HttpRequestAction{
						Method: "GET",
						Path:   "/car/1",
					},
					Assertions: goponent.JsonResponseAsserter[Car]{
						ExpectedBody:       Car{Make: "Subaru", Model: "Outback", Id: "1"},
						ExpectedStatusCode: 200,
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
