package goponent

import (
	"net/http"
	"testing"
)

type ResponseAsserter interface {
	Assert(t *testing.T, res *http.Response) error
}

type ContextSetter interface {
}

type Test struct {
	Name  string
	Steps []Step
}

type Step struct {
	Name           string
	Request        RequestGenerator
	Assertions     ResponseAsserter
	ContextSetters []ContextSetter
}

func RunTests(t *testing.T, tests []Test, baseUrl string) {
	for _, test := range tests {
		RunTest(t, test, baseUrl)
	}
}

func RunTest(t *testing.T, test Test, baseUrl string) {
	t.Run(test.Name, func(t *testing.T) {
		for num, step := range test.Steps {
			t.Run(step.Name, func(t *testing.T) {
				body := step.Request.GetBody()

				req, err := http.NewRequest(http.MethodPost, baseUrl+step.Request.GetPath(), body)
				if err != nil {
					t.Fatalf("error making request for step: %d - %s: %+v", num, step.Name, err)
				}
				req.Header.Set("Content-Type", step.Request.GetContentType())
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatalf("error making request for step: %d - %s: %+v", num, step.Name, err)
				}

				if step.Assertions != nil {
					err := step.Assertions.Assert(t, res)
					if err != nil {
						t.Fatalf("error making request for step: %d - %s: %+v", num, step.Name, err)
					}
				}

			})
		}
	})
}
