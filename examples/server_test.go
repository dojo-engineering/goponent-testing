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
					ContextSetters: goponent.ContextSetterJson[Car]{
						Properties: map[string]func(response Car) any{
							"carId": func(response Car) any {
								return response.Id
							},
						},
					},
					Assertions: goponent.JsonResponseAsserter[Car]{
						ExpectedBodyFunc: func(ctx *goponent.Context) Car {
							return Car{Make: "Subaru", Model: "Outback", Id: ctx.GetString("carId")}
						},
						ExpectedStatusCode: 200,
					},
				},
				{
					Name: "Get car returns 200 and car",
					Arrange: []goponent.Arranger{
						goponent.ArrangeHttpDependencyAction{
							Host:   "https://www.example.com",
							Path:   "/example-payload",
							Method: "GET",
							Body:   "1234",
						},
					},
					Act: goponent.HttpRequestAction{
						Method: "GET",
						PathFunc: func(ctx *goponent.Context) string {
							return "/car/" + ctx.GetString("carId")
						},
					},
					Assertions: goponent.JsonResponseAsserter[Car]{
						ExpectedBodyFunc: func(ctx *goponent.Context) Car {
							return Car{
								Make:           "Subaru",
								Model:          "Outback",
								Id:             ctx.GetString("carId"),
								RegistrationId: "1234",
							}
						},
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
