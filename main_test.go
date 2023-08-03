package main_test

import (
	"context"
	"fmt"
	app "go-bdd-calculator"
	"testing"

	"github.com/cucumber/godog"
)

type appContext struct{}

func given(ctx context.Context, num int) (context.Context, error) {
	return context.WithValue(ctx, appContext{}, app.NewCalculator(num)), nil
}

func add(ctx context.Context, num int) error {
	calculator := ctx.Value(appContext{}).(*app.Calculator)

	calculator.Add(num)

	return nil
}

func subtract(ctx context.Context, num int) error {
	calculator := ctx.Value(appContext{}).(*app.Calculator)

	calculator.Subtract(num)

	return nil
}

func multiply(ctx context.Context, num int) error {
	calculator := ctx.Value(appContext{}).(*app.Calculator)

	calculator.Multiply(num)

	return nil
}

func divide(ctx context.Context, num int) error {
	calculator := ctx.Value(appContext{}).(*app.Calculator)

	calculator.Divide(num)

	return nil
}

func expected(ctx context.Context, num int) error {
	calculator := ctx.Value(appContext{}).(*app.Calculator)
	if calculator.Result() != num {
		return fmt.Errorf("expected %d, got %d", num, calculator.Result())
	}

	return nil
}

func initializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^(\d+)$`, given)
	ctx.Step(`^\+ (\d+)$`, add)
	ctx.Step(`^- (\d+)$`, subtract)
	ctx.Step(`^\* (\d+)$`, multiply)
	ctx.Step(`^\/ (\d+)$`, divide)
	ctx.Step(`^It should be (\d+)$`, expected)
}

func TestFeatures(t *testing.T) {
	godog.TestSuite{
		ScenarioInitializer: initializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}.Run()
}
