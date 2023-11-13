package goponent

import (
	"net/http"
	"testing"

	"github.com/h2non/gock"
)

type ResponseAsserter interface {
	Assert(t *testing.T, res *http.Response) error
}

type Test struct {
	Name  string
	Steps []Step
}

type Arranger interface {
	Arrange(t *testing.T, context *Context, stepContext *Context) error
}

type Actor interface {
	Act(t *testing.T, context *Context, stepContext *Context) error
}

type ContextSetter interface {
	SetContext(t *testing.T, context *Context, stepContext *Context) error
}

type Asserter interface {
	Assert(t *testing.T, context *Context, stepContext *Context) error
}

type Step struct {
	Name           string
	Arrange        []Arranger
	Act            Actor
	ContextSetters ContextSetter
	Assertions     Asserter
}

func RunTests(t *testing.T, tests []Test, baseUrl string) {
	for _, test := range tests {
		RunTest(t, test, baseUrl)
	}
}

func RunTest(t *testing.T, test Test, baseUrl string) {
	gock.EnableNetworking()
	t.Run(test.Name, func(t *testing.T) {
		testContext := newContext()
		ContextSet(testContext, "baseUrl", baseUrl)
		for num, step := range test.Steps {

			t.Run(step.Name, func(t *testing.T) {
				stepContext := newContext()

				if step.Arrange != nil {
					for _, a := range step.Arrange {
						err := a.Arrange(t, testContext, stepContext)
						if err != nil {
							t.Fatalf("error arraging for step: %d - %s: %+v", num, step.Name, err)
						}
					}
				}

				if step.Act == nil {
					t.Fatalf("step has no act: %d - %s", num, step.Name)
				}
				err := step.Act.Act(t, testContext, stepContext)
				if err != nil {
					t.Fatalf("error acting for step: %d - %s: %+v", num, step.Name, err)
				}

				if step.ContextSetters != nil {
					err = step.ContextSetters.SetContext(t, testContext, stepContext)
					if err != nil {
						t.Fatalf("error setting context for step: %d - %s: %+v", num, step.Name, err)
					}
				}

				if step.Assertions != nil {
					err = step.Assertions.Assert(t, testContext, stepContext)
					if err != nil {
						t.Fatalf("error asserting for step: %d - %s: %+v", num, step.Name, err)
					}
				}
			})

			gock.Clean()
			gock.Flush()
		}
	})
}
