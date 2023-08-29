package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestContext(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name     string
		expected *string
	}

	strategyAExpectedValue := "ABC"
	strategyBExpectedValue := "CCC"

	testCases := []TestCase{
		{name: "A", expected: &strategyAExpectedValue},
		{name: "B", expected: &strategyBExpectedValue},
		{name: "C", expected: nil},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var (
				context Context
			)

			options := []fx.Option{
				StrategyModule,
				fx.Populate(&context),
			}

			if !testing.Verbose() {
				options = append(options, fx.NopLogger)
			}
			app := fxtest.New(t, options...)
			app.RequireStart()
			defer func() {
				app.RequireStop()
			}()

			assert.Equal(t, tc.expected, context.resolve(tc.name), "Invalid result of strategy")
		})
	}
}
